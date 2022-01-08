package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"net/smtp"
	"strconv"
)

var tpl *template.Template
var x int = 0

func init() {
	tpl, _ = template.ParseGlob("templates/*.html")
}
func handleregister(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "register.html", nil)
}
func verifyreg(w http.ResponseWriter, r *http.Request) {
	from := "mywork.p98@gmail.com"
	password := ""
	to := []string{
		r.FormValue("email"),
	}
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	if x == 0 {
		tpl.ExecuteTemplate(w, "code.html", nil)
		x = rand.Intn(1000000)
		z := strconv.Itoa(x)
		message := []byte(z)
		auth := smtp.PlainAuth("", from, password, smtpHost)
		smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
		fmt.Println("email is sent")
		fmt.Println(x)
		username := r.FormValue("email")
		fmt.Println("dafa end point hit:::")
		fmt.Println(username)
	} else {
		tpl.ExecuteTemplate(w, "code.html", nil)
		// r.ParseForm()
		username := r.FormValue("email")
		fmt.Println("dafa end point hit:::")
		fmt.Println(username)
	}
}
func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	z := strconv.Itoa(x)
	if r.FormValue("code") == z {
		x = 0
		z = ""
		fmt.Fprintf(w, "welcome")
	} else {
		http.Redirect(w, r, "/verifyreg", http.StatusSeeOther)
	}
}

func handleRequests() {
	r := http.NewServeMux()
	r.HandleFunc("/", handleregister)
	r.HandleFunc("/verifyreg", verifyreg)
	r.HandleFunc("/login", login)
	log.Fatal(http.ListenAndServe(":5000", r))
}

func main() {
	handleRequests()
}
