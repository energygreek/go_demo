package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

type Result struct{
	Ret int
	Reason string
	Data interface{}
}
type User struct {
	UserName string
}
type adminController struct {
}

func (this *adminController)IndexAction(w http.ResponseWriter, r *http.Request, user string) {
	t, err := template.ParseFiles("template/html/admin/index.html")
	if err != nil{
		log.Println(err)
	}
	t.Execute(w, &User{user})
}

type loginController struct {
}
func (this *loginController)IndexAction(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/html/login/index.html")
	if err != nil {
		log.Println(err)
	}
	log.Println("Index of login")
	t.Execute(w, nil)
}


type ajaxController struct {
}

func (this * ajaxController)IndexAction(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type","application/json")
	err := r.ParseForm()
	if err != nil{
		log.Println("pase error")
	}
	admin_name := r.FormValue("admin_name")
	admin_password := r.FormValue("admin_password")

	log.Println("passwd", admin_password)
	//param validate skipped
	// db skipped

	//expiration := time.Unix(1,0)
	cookie := http.Cookie{Name:"admin_name", Value:admin_name, Path:"/"}
	http.SetCookie(w,&cookie)

	OutputJson(w,1,"Success",nil)
}

func OutputJson(w http.ResponseWriter, ret int, reason string, i interface{}){
	out := &Result{ret,reason,i}
	b,err := json.Marshal(out)
	if err != nil{
		return
	}
	w.Write(b)
}
