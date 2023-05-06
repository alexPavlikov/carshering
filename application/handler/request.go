package handler

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/alexPavlikov/scud.ru/algorithms"
	"github.com/alexPavlikov/scud.ru/models"
)

var Done = true

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/index.html", "web/header.html", "web/footer.html")
	algorithms.CheckErr(err)

	data := map[string]interface{}{"Config": models.Config}
	header := map[string]interface{}{"Config": models.Config, "Done": Done}

	tmpl.ExecuteTemplate(w, "header", header)
	tmpl.ExecuteTemplate(w, "index", data)
	tmpl.ExecuteTemplate(w, "footer", data)

}

func subscribeHandler(w http.ResponseWriter, r *http.Request) { //подписка на рассылку сообщений по почте
	if r.Method == "POST" {
		r.ParseForm()
		mail := r.FormValue("mail")
		fmt.Println(mail)
		http.Redirect(w, r, "/#footer", http.StatusSeeOther)
	}
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/auth.html")
	algorithms.CheckErr(err)

	if r.Method == "POST" {
		r.ParseForm()
		email := r.FormValue("email")
		password := r.FormValue("password")
		fmt.Println(email, password) // запрос на проверку наличия пользователя
		Done = false
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	data := map[string]interface{}{"Config": models.Config}
	tmpl.ExecuteTemplate(w, "auth", data)
}

func regHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/registration.html")
	algorithms.CheckErr(err)

	if r.Method == "POST" {
		r.ParseForm()
		name := r.FormValue("name")
		email := r.FormValue("email")
		password := r.FormValue("password")
		fmt.Println(name, email, password) //запросы на добавление пользователя
		http.Redirect(w, r, "/auth", http.StatusSeeOther)
	}

	data := map[string]interface{}{"Config": models.Config}
	tmpl.ExecuteTemplate(w, "reg", data)
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/profile.html", "web/header.html", "web/footer.html")
	algorithms.CheckErr(err)

	data := map[string]interface{}{"Config": models.Config}
	header := map[string]interface{}{"Config": models.Config, "Done": Done}

	tmpl.ExecuteTemplate(w, "header", header)
	tmpl.ExecuteTemplate(w, "profile", data)
	tmpl.ExecuteTemplate(w, "footer", data)
}

func exitHandler(w http.ResponseWriter, r *http.Request) {
	Done = true
	http.Redirect(w, r, "/", http.StatusSeeOther)
	//http.Redirect(w, r, "/auth", http.StatusSeeOther)
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/product.html", "web/header.html", "web/footer.html")
	algorithms.CheckErr(err)

	data := map[string]interface{}{"Config": models.Config}
	header := map[string]interface{}{"Config": models.Config, "Done": Done}

	tmpl.ExecuteTemplate(w, "header", header)
	tmpl.ExecuteTemplate(w, "product", data)
	tmpl.ExecuteTemplate(w, "footer", data)
}

func cardHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/card.html", "web/header.html", "web/footer.html")
	algorithms.CheckErr(err)

	r.ParseForm()
	name := r.FormValue("name")
	fmt.Println(name)

	data := map[string]interface{}{"Config": models.Config}
	header := map[string]interface{}{"Config": models.Config, "Done": Done}

	tmpl.ExecuteTemplate(w, "header", header)
	tmpl.ExecuteTemplate(w, "card", data)
	tmpl.ExecuteTemplate(w, "footer", data)
}
