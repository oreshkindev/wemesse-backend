// Package router provides
package router

import (
	"context"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// определим структуру
type Router struct {
	*chi.Mux
}

// создаем новый роутер
func NewRouter(ctx context.Context) (*Router, error) {
	r := chi.NewRouter()

	// определяем что будем работать с json
	r.Use(render.SetContentType(render.ContentTypeJSON))

	return &Router{r}, nil
}
