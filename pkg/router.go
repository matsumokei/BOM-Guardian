package server

import "github.com/go-chi/chi/v5"


//ServerInterface wrrap
func NewRouter(si ServerInterface) *chi.Mux{
	r := chi.NewRouter()

	r.Route("/bom", func(r chi.Router)  {
		r.Get("/serch", si.HandleSearchPackages)
		//r.Post("/", si.HandleImportBOM)
		//r.Delete("/", si.HandleDeleteComponent)
	})
	return r
}



