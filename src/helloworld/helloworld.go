package main

import (
	"net/http"
)

func main(){
	http.Handle("/css/", http.FileServer(http.Dir("template")))
	http.Handle("/js/", http.FileServer(http.Dir("template")))

	http.HandleFunc("/admin/",adminHandler)
	http.HandleFunc("/login/",loginHandler)
	http.HandleFunc("/ajax/",ajaxHandler)

	http.ListenAndServe(":8088",nil)
}