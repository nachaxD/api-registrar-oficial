package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"main.go/rutas"

	firebase "firebase.google.com/go"
	"github.com/gorilla/mux"
	"google.golang.org/api/option"
)

func main() {
	// Inicializa la configuración de Firebase
	ctx := context.Background()
	opt := option.WithCredentialsFile(".env")
	config := &firebase.Config{ProjectID: "test-5eebf"}
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatalln(err)
	}

	// Inicializa el cliente de Firestore
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	// Cerrar el cliente de Firestore cuando ya no se necesite
	defer client.Close()

	router := mux.NewRouter()
	const port string = ":8080"
	router.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "UP and running...") // imprime la respuesta del cliente
	})

	router.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		rutas.RegisterUser(w, r, app)
	}).Methods("POST")

	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		rutas.LoginUser(w, r, app)
	}).Methods("POST")

	router.HandleFunc("/update-password", func(w http.ResponseWriter, r *http.Request) {
		rutas.UpdatePassword(w, r, app)
	}).Methods("PUT")

	log.Println("Server listening on port", port) // imprime en el servidor
	log.Fatal(http.ListenAndServe(port, router))

}
