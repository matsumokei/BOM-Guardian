package main

import (
	"net/http"

	server "github.com/matsumokei/BOM-Guardian/pkg"
)



func main() {
	router := server.NewRouter()
	//router.Mount("/api", restApiHandler)
	server := http.Server{
		Addr: "8080",
		Handler: router,
	}
	server.ListenAndServe()
}

