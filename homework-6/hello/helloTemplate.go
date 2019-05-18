package main

/*
** Homework-6: Дополнить функцию hello() нашего http сервера так, чтобы принять и отобразить один GET параметр.
** Khamenev Sergei
** 18.05.2019
 */

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	Name := req.URL.Query().Get("name")
	// http://localhost:8080/hello?name=Sergei - пример URL
	t, err := template.New("foo").Parse(`
	{{define "ViewName"}}
	<doctype html>
	<html>
		<head>
			<title>Hello {{ . }}</title>
		</head>
		<body>
			Hello {{ . }}!
		</body>
	</html>
	{{end}}`)

	if err != nil {
		log.Println("Не удалось распарсить шаблон")
		return
	}

	if err := t.ExecuteTemplate(res, "ViewName", Name); err != nil {
		io.WriteString(res, "Параметр name не получен")
		log.Println("Параметр name не получен")
	}
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8080", nil)
}
