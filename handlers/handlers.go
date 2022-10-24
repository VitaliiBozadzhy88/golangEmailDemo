package handlers

import (
	"GoUserEmail/model"
	"html/template"
	"net/http"
)

var actCode = 48092 //for example

func Handlers() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, "templates/index.html")
	})

	http.HandleFunc("/confirmation", func(writer http.ResponseWriter, request *http.Request) {
		emailFromField := request.FormValue("email")
		data := model.User{
			Email:          emailFromField,
			ActivationCode: actCode,
			Status:         "ERROR DURING EXECUTION",
		}

		if emailFromField == "" || request.FormValue("retype_email") == "" {
			tmpl, _ := template.ParseFiles("templates/emptyFields.html")
			tmpl.Execute(writer, data)
		} else {
			if (model.User{}.GetEmail(emailFromField)) != emailFromField {
				model.User{}.Insert(emailFromField, actCode)
				tmpl, _ := template.ParseFiles("templates/confirmation.html")
				tmpl.Execute(writer, data)
			} else {
				tmpl, _ := template.ParseFiles("templates/warning.html")
				tmpl.Execute(writer, data)
			}
		}
	})

	http.HandleFunc("/success", func(writer http.ResponseWriter, request *http.Request) {
		emailFromField := request.FormValue("email")
		emailFromDB := model.User{}.GetEmail(emailFromField)
		codeFromField := request.FormValue("code")

		data := model.User{
			Email:  model.User{}.GetEmail(emailFromField),
			Status: "REGISTERED",
		}

		if request.FormValue("email") == "" || request.FormValue("code") == "" {
			http.ServeFile(writer, request, "templates/emptyFields.html")
		} else {
			if codeFromField == "48092" {
				if request.FormValue("code") != "" && emailFromDB == emailFromField {
					tmpl, _ := template.ParseFiles("templates/success.html")
					tmpl.Execute(writer, data)
				}
			} else {
				http.ServeFile(writer, request, "templates/code.html")
			}
		}
	})

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic(err)
	}
}
