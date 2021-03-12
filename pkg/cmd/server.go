package cmd

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"

	service "github.com/elorusso/vendor-health-be/pkg/service/v1"
)

const (
	portNumber = "80"
)

func RunServer() error {

	//create rounter
	router := chi.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	router.Use(c.Handler)

	fmt.Println("Registering and Serving at :" + portNumber)

	//register routes
	router.HandleFunc("/v1/amazon-status", service.AmazonStatusRequestHandler)
	router.HandleFunc("/v1/google-status", service.GoogleStatusRequestHandler)
	router.HandleFunc("/v1/all-status", service.AllStatusRequestHandler)

	return http.ListenAndServe(":"+portNumber, router)
}
