package handler

import (
	"html/template"
	"net/http"
	"shinkansen_rest_api/helper"
	"shinkansen_rest_api/resources"
)

func ClientIndex(writer http.ResponseWriter,data interface{}) {
	tmpl, err := template.ParseFS(resources.Templates,"templates/*.html");
	helper.PanicIfError(err);
	tmpl.ExecuteTemplate(writer,"index.html",data);
}

func ClientSignUp(writer http.ResponseWriter,data interface{}) {
	tmpl, err := template.ParseFS(resources.Templates,"templates/*.html");
	helper.PanicIfError(err);
	tmpl.ExecuteTemplate(writer,"signup.html",data);
}

func ClientLogin(writer http.ResponseWriter,data interface{}) {
	tmpl, err := template.ParseFS(resources.Templates,"templates/*.html");
	helper.PanicIfError(err);
	tmpl.ExecuteTemplate(writer,"login.html",data);
}

func ClientDashboard(writer http.ResponseWriter,data interface{}) {
	tmpl, err := template.ParseFS(resources.Templates,"templates/*.html");
	helper.PanicIfError(err);
	tmpl.ExecuteTemplate(writer,"dashboard.html",data);
}
func ClientDocs(writer http.ResponseWriter,data interface{}) {
	tmpl, err := template.ParseFS(resources.Templates,"templates/*.html");
	helper.PanicIfError(err);
	tmpl.ExecuteTemplate(writer,"docs.html",data);
}
