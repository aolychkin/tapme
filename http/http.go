package http

import (
	"fmt"
	"log"
	"net/http"
)

func startServe(){
	http.HandleFunc("/", sayHello) // Устанавливаем роутер
	err := http.ListenAndServe(":8080", nil) // устанавливаем порт веб-сервера
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Привет!")
}