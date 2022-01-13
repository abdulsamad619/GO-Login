package main

import (
	"fmt"
	"context"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	// "net/smtp"
	"strconv"
	"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var tpl *template.Template
var x int = 0
var em string=""
var ps string=""
// client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
// if err != nil {
//                 panic(err)
//         }
func init() {
	tpl, _ = template.ParseGlob("templates/*.html")
}
func handleregister(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "register.html", nil)
}
func verifyreg(w http.ResponseWriter, r *http.Request) {
	if r.Method=="POST"{
		em=r.FormValue("email")
		ps=r.FormValue("password")
		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
		if err != nil {
                panic(err)
        }
		collection := client.Database("GoDb").Collection("users")
		oneDoc := MongoFields{
		Email: em,
		Password: ps,
		}
		collection.InsertOne(context.TODO(), oneDoc)
		fmt.Println(em," and ",ps)
		// from := "mywork.p98@gmail.com"
		// password := "gallardo.98"
		// to := []string{
		// 	r.FormValue("email"),
		// }
		// smtpHost := "smtp.gmail.com"
		// smtpPort := "587"
		if x == 0 {
			tpl.ExecuteTemplate(w, "code.html", nil)
			x = rand.Intn(1000000)
			// z := strconv.Itoa(x)
			// message := []byte(z)
			// auth := smtp.PlainAuth("", from, password, smtpHost)
			// smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
			// fmt.Println("email is sent")
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
	}else{
		http.Redirect(w, r, "/", http.StatusSeeOther)
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
type MongoFields struct {
Email string `json:"email"`
Password string `json:"password"`

}
func main() {
	

	handleRequests()

}
