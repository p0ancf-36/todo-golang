package api

import "github.com/go-chi/chi/v5"

type Handlers interface {
	UseHandlers(r chi.Router)
}
