// Package controller provides
package controller

import (
	"context"
	"net/http"
	"wemesse/conf"
	"wemesse/model"
	"wemesse/service"
	"wemesse/util"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// определяем структуру контроллера
type UpdatesController struct {
	ctx      context.Context
	services *service.Manager
	conf     *conf.Conf
}

// создаем контроллер для проверки обновлений
func NewUpdates(ctx context.Context, services *service.Manager, conf *conf.Conf) *UpdatesController {
	return &UpdatesController{
		ctx:      ctx,
		services: services,
		conf:     conf,
	}
}

// структура тела ответа. позже разберемся с ней
type Response struct {
	Message string
	Error   string
}

// структура тела ответа.
type URI struct {
	URI string
}

// получаем версию приложения с клиента
func (ctr *UpdatesController) GetUpdate(w http.ResponseWriter, r *http.Request) {
	// ВАЖНО:
	// не проверяем сигнатуру т.к у старых пользователей еще нет ключа

	// достаем версию из запроса
	appVersion := chi.URLParam(r, "version")
	// передаем версию в интерфейс менеджера
	updates, err := ctr.services.Updates.GetUpdate(ctr.ctx, appVersion)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	updates.URI = util.GenerateURI(ctr.conf.DestURI, updates.AppName)

	// возвращаем json
	render.JSON(w, r, updates)
}

// служебный метод
// получаем списки обновлений и статистику
func (ctr *UpdatesController) GetUpdates(w http.ResponseWriter, r *http.Request) {

	// достаем подпись из заголовка запроса
	signature := r.Header.Get("signature")

	// проверяем, есть ли совпадение в бд и если есть то ...
	if err := ctr.services.Updates.VerifySignature(ctr.ctx, signature); err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	} else {
		// стучимся в интерфейс менеджера
		updates, err := ctr.services.Updates.GetUpdates(ctr.ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		// заменим значение в массиве для каждого эллемента
		for i := range updates {
			updates[i].URI = util.GenerateURI(ctr.conf.SourceURI, updates[i].TargetABI, updates[i].AppVersion, updates[i].AppName)
		}

		// возвращаем json
		render.JSON(w, r, updates)
	}
}

// служебный метод
// получаем приложение и описание
func (ctr *UpdatesController) PostUpdate(w http.ResponseWriter, r *http.Request) {
	// достаем подпись из заголовка запроса
	signature := r.Header.Get("signature")

	// проверяем, есть ли совпадение в бд и если есть то ...
	if err := ctr.services.Updates.VerifySignature(ctr.ctx, signature); err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	} else {

		// забираем файл из формы
		f, header, err := r.FormFile("tmpFile")
		if err != nil {

			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer f.Close()

		// создаем директорию на сервере /var/www/домен/source/версия/
		dst, err := util.CreateDir(ctr.conf, r.FormValue("targetABI"), r.FormValue("appVersion"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// создаем локальный файл
		err = util.CreateFile(dst, header.Filename, f)
		if err != nil {

			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		appChecksum := util.GetChecksum(f)

		appName := header.Filename

		appNotes := r.FormValue("appNotes")

		appSize := util.ByteCountSI(header.Size)

		appVersion := r.FormValue("appVersion")

		skipped := util.GetSuffix(appVersion)

		targetABI := r.FormValue("targetABI")

		app := &model.App{
			AppChecksum: appChecksum,
			AppName:     appName,
			AppNotes:    appNotes,
			AppSize:     appSize,
			AppVersion:  appVersion,
			Skipped:     skipped,
			TargetABI:   targetABI,
			URI:         util.GenerateURI(ctr.conf.SourceURI, targetABI, appVersion, appName),
		}

		// передаем модель в интерфейс менеджера
		updates, err := ctr.services.Updates.PostUpdate(ctr.ctx, app)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response := &model.App{
			ID:          updates.ID,
			AppChecksum: updates.AppChecksum,
			AppName:     updates.AppName,
			AppNotes:    updates.AppNotes,
			AppSize:     updates.AppSize,
			AppVersion:  updates.AppVersion,
			Skipped:     updates.Skipped,
			TargetABI:   updates.TargetABI,
			URI:         util.GenerateURI(ctr.conf.SourceURI, updates.TargetABI, updates.AppVersion, updates.AppName),
		}

		// возвращаем json
		render.JSON(w, r, response)
	}
}
