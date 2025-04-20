package main

import (
	"context"
	_ "fmt"
	_ "html/template"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
)

func main() {

	conn, err := pgx.Connect(context.Background(), "postgres://postgres:Sed60056005@localhost:5432/kurs")
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}
	defer conn.Close(context.Background())

	// Обработчик страницы входа
	http.HandleFunc("/", ShowLoginPage)
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
        ParsingLoginPage(w, r, conn)
    })
	http.HandleFunc("/admin", ShowAdminPage)

	// Запуск сервера
	http.ListenAndServe(":8080", nil)
}

