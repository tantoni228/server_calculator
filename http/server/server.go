package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tantoni228/server_calculator/http/server/handler"
)

func Run() {

	for i := 1; i <= 2; i++ {
		go start_server(8080 + i)
	}
	go func() {
		r := http.NewServeMux()
		r.HandleFunc("/calculate/", handler.Handler_tasks)
		r.HandleFunc("/steps/", handler.Handler_steps)
		http.Handle("/", r)

		fmt.Println("Сервер запущен на порту 8080")
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()
	select {}
}

func start_server(port int) {
	r1 := http.NewServeMux()
	r1.HandleFunc("/", handler.Handler_operation)
	go func() { log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r1)) }()
	fmt.Printf("Сервер запущен на порту %d\n", port)
}
