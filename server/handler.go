package main

import (
	_ "fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/jackc/pgx/v5"
)

func ShowLoginPage(w http.ResponseWriter, r *http.Request) {
	// Собираем абсолютный путь к шаблону
	tmplPath := filepath.Join("..","html", "registration.html")
	
	tmpl, err := template.ParseFiles(tmplPath)// сам шаблон входа
	if err != nil { 
		http.Error(w, "Ошибка загрузки страницы: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Ошибка отображения страницы: "+err.Error(), http.StatusInternalServerError)
	}
}

func ParsingLoginPage(w http.ResponseWriter, r *http.Request, conn *pgx.Conn) {
	if r.Method != "POST" {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return 
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Ошибка при разборе формы", http.StatusBadRequest)
		return
	}

	Fullname := r.FormValue("fio")
	Password := r.FormValue("password")	

	user, err := VerifyUser(Fullname, Password, conn)
    if err != nil {
        if err.Error() == "неверные учетные данные" {
            http.Error(w, err.Error(), http.StatusUnauthorized)
        } else {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        return
    }

	switch user.Role {
	case "admin":
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	case "user":
		http.Redirect(w, r, "/user", http.StatusSeeOther)
	default:
		http.Error(w, "Неизвестная роль пользователя", http.StatusForbidden)
	}

}


func ShowAdminPage(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("..","html", "UserPage.html")
	
	tmpl, err := template.ParseFiles(tmplPath)// сам шаблон входа
	if err != nil { 
		http.Error(w, "Ошибка загрузки страницы: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Ошибка отображения страницы: "+err.Error(), http.StatusInternalServerError)
	}
}