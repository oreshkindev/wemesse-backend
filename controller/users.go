// Package controller provides
package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"wemesse/model"
	"wemesse/service"

	"github.com/go-chi/render"
)

// определяем структуру контроллера
type UsersController struct {
	ctx      context.Context
	services *service.Manager
}

// создаем контроллер для пользователей
func NewUser(ctx context.Context, services *service.Manager) *UsersController {
	return &UsersController{
		ctx:      ctx,
		services: services,
	}
}

// служебный метод
// получаем списки пользователей
func (ctr *UsersController) GetUsers(w http.ResponseWriter, r *http.Request) {
	// достаем подпись из заголовка запроса
	signature := r.Header.Get("signature")

	// проверяем, есть ли совпадение в бд и если есть то ...
	if err := ctr.services.Updates.VerifySignature(ctr.ctx, signature); err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	} else {
		// стучимся в интерфейс менеджера
		users, err := ctr.services.Users.GetUsers(ctr.ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		// возвращаем json
		render.JSON(w, r, users)
	}
}

// получаем данные с клиента
func (ctr *UsersController) PostUser(w http.ResponseWriter, r *http.Request) {
	// достаем подпись из заголовка запроса
	signature := r.Header.Get("signature")

	// проверяем, есть ли совпадение в бд и если есть то ...
	if err := ctr.services.Updates.VerifySignature(ctr.ctx, signature); err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	} else {

		// т.к мы получаем модель пользователя с клиента, определим ее
		user := model.Users{}

		// декодируем полученный json и заполняем модель Users
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "CreateConfigHandler read invalid params", http.StatusBadRequest)
			return
		}

		// передаем заполненую модель в интерфейс менеджера
		postUser, err := ctr.services.Users.PostUser(ctr.ctx, &user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		// возвращаем json
		render.JSON(w, r, postUser)
	}
}
