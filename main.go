package main

import (
	"fmt"
	"net/http"

	authcontroller "appgo/controllers"
)

func main() {
	http.HandleFunc("/", authcontroller.Index)
	http.HandleFunc("/login", authcontroller.Login)
	http.HandleFunc("/logout", authcontroller.Logout)
	http.HandleFunc("/register", authcontroller.Register)

	fmt.Println("Server Running: http://localhost:4000")
	http.ListenAndServe(":4000", nil)
}
