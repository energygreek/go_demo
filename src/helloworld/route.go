package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
)

func loginHandler(w http.ResponseWriter, r *http.Request)  {
	pathInfo := strings.Trim(r.URL.Path,"/")
	parts := strings.Split(pathInfo,"/")
	log.Println("login Handel")
	var action = ""
	if len(parts) > 1 {
		action = strings.Title(parts[1] )+ "Action"
	}
	login := &loginController{}
	controller := reflect.ValueOf(login)
	method := controller.MethodByName(action)

	if !method.IsValid() {
		method = controller.MethodByName(strings.Title("index") + "Action")
	}
	requestValue := reflect.ValueOf(r)
	responseValue := reflect.ValueOf(w);

	method.Call([]reflect.Value{responseValue,requestValue})
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	// 获取cookie
	cookie, err := r.Cookie("admin_name")
	if err != nil || cookie.Value == ""{
		http.Redirect(w, r, "/login/index", http.StatusFound)
	}

	pathInfo := strings.Trim(r.URL.Path,"/")
	parts := strings.Split(pathInfo,"/");
	var action = ""
	if len(parts) > 1 {
		action = strings.Title(parts[1] )+ "Action"
	}

	admin := &adminController{}
	controller := reflect.ValueOf(admin)
	method := controller.MethodByName(action)
	if !method.IsValid() {
		method = controller.MethodByName(strings.Title("index") + "Action")
	}
	requestValue := reflect.ValueOf(r)
	responseValue := reflect.ValueOf(w)
	userValue := reflect.ValueOf(cookie.Value)
	fmt.Println("userValue",userValue)
	method.Call([]reflect.Value{responseValue, requestValue, userValue})
}

func ajaxHandler(w http.ResponseWriter, r *http.Request) {
	pathInfo := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(pathInfo, "/")
	var action = ""
	if len(parts) > 1 {
		action = strings.Title(parts[1]) + "Action"
	}

	login := &ajaxController{}
	controller := reflect.ValueOf(login)
	method := controller.MethodByName(action)
	if !method.IsValid() {
		method = controller.MethodByName(strings.Title("index") + "Action")
	}
	requestValue := reflect.ValueOf(r)
	responseValue := reflect.ValueOf(w)
	method.Call([]reflect.Value{responseValue, requestValue})
}