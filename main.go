package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/mattn/go-sqlite3"

	pkgHttp "todo-golang.com/api/http"
	pkgSql "todo-golang.com/repo/sql"
	"todo-golang.com/views"
)

const PORT = 8080

func main() {
	views.Init()

	db, err := sql.Open("sqlite3", "./sqlite3.db")
	if err != nil {
		log.Fatalln("Error during connecting to database")
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", index)

	taskRepo := pkgSql.NewTaskRamStorage(db)
	taskHandler := pkgHttp.NewTask(&taskRepo)

	r.Route("/task", func(r chi.Router) {
		taskHandler.UseHandlers(r)
	})

	log.Println("Starting server on ", PORT)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), r); err != nil {
		log.Fatalln("Error starting server:", err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	if err := views.Render(w, "index", nil); err != nil {
		http.Error(w, "template `index` not found", http.StatusInternalServerError)
	}
}
