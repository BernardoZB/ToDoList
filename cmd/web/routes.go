package main

import "net/http"

func (app *application) routes() http.Handler {
  
  mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/salvar", app.salvar)
	mux.HandleFunc("/excluir", app.excluir)
	mux.HandleFunc("/tarefa", app.mostrarTarefa)
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

  return app.logRequest( secureHeaders( mux ) )
}