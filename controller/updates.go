// Package controller provides
package controller

import (
	"context"
	"net/http"
	"os"
	"path/filepath"
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
			updates[i].URI = ctr.conf.SourceURI + updates[i].AppVersion + "/" + updates[i].AppName
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

		// err := r.ParseMultipartForm(32 << 20)
		// if err != nil {
		// 	http.Error(w, "большой файл", http.StatusNoContent)
		// }

		// получаем приложение из постмана
		f, h, err := r.FormFile("tmpFile")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer f.Close() // закроем наш пул

		appChecksum := util.GetChecksum(f)

		appName := h.Filename

		appNotes := r.FormValue("appNotes")

		appSize := util.ByteCountSI(h.Size)

		appVersion := r.FormValue("appVersion")

		skipped := util.GetSuffix(appVersion)

		app := &model.App{
			AppChecksum: appChecksum,
			AppName:     appName,
			AppNotes:    appNotes,
			AppSize:     appSize,
			AppVersion:  appVersion,
			Skipped:     skipped,
			URI:         ctr.conf.DestURI + appName,
		}

		// создаем директорию на сервере в /var/www/messenger.tbcc.com/source/
		source, err := util.CreateDir(ctr.conf.Deploy, app.AppVersion)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else {
			// переместим приложение в директорию на сервере
			move, err := os.Create(filepath.Join(source, filepath.Base(app.AppName)))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			defer move.Close() // закроем наш пул
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
			URI:         ctr.conf.SourceURI + updates.AppVersion + "/" + appName,
		}

		// возвращаем json
		render.JSON(w, r, response)
	}
}
