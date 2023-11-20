package rutas

import (
	"bytes"

	"fmt"
	"io/ioutil"
	"net/http"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

func LoginUser(resp http.ResponseWriter, req *http.Request, app *firebase.App) {

	email := req.FormValue("email")
	password := req.FormValue("password")

	// Define la URL del punto de conexión de Firebase Identity Toolkit para iniciar sesión con contraseña
	url := "https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=AIzaSyAN5HyOCaS1NMqiVKgcYaN1s6fq3oJWbMw"

	// Construye la carga útil de la solicitud JSON
	payload := fmt.Sprintf(`{
        "email": "%s",
        "password": "%s",
        "returnSecureToken": true
    }`, email, password)

	// Realiza la solicitud HTTP POST
	response, err := http.Post(url, "application/login", bytes.NewBuffer([]byte(payload)))

	if err != nil {
		fmt.Fprintln(resp, "Error en la solicitud HTTP:", err)
		return
	}
	defer response.Body.Close()

	// Lee y procesa la respuesta
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Fprintln(resp, "Error al leer la respuesta:", err)
		return
	}

	// Verifica el código de estado de la respuesta
	if response.StatusCode == http.StatusOK {
		// Autenticación exitosa
		fmt.Fprintln(resp, "Usuario autenticado con éxito.")
	} else {
		// Autenticación fallida
		fmt.Fprintln(resp, "Error de autenticación:", string(responseBody))
	}

	// Puedes procesar la respuesta de Firebase según tus necesidades

	// No es necesario devolver nada aquí, ya que la respuesta se maneja en el lugar.
}

func UpdatePassword(resp http.ResponseWriter, req *http.Request, app *firebase.App) {
	// Obtén el ID del usuario que desea actualizar la contraseña
	userID := req.FormValue("userID")

	// Obtén la nueva contraseña
	newPassword := req.FormValue("newPassword")

	// Inicia una instancia del cliente de autenticación de Firebase
	authClient, err := app.Auth(req.Context())

	if err != nil {
		http.Error(resp, "Error al iniciar el cliente de autenticación", http.StatusInternalServerError)
		return
	}

	// Actualiza la contraseña del usuario
	_, err = authClient.UpdateUser(req.Context(), userID, (&auth.UserToUpdate{}).Password(newPassword))

	if err != nil {
		http.Error(resp, "Error al actualizar la contraseña del usuario", http.StatusInternalServerError)
		return
	}

	// Contraseña actualizada con éxito
	fmt.Fprintln(resp, "Contraseña actualizada con éxito.")
}
