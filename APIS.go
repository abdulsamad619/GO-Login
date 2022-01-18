package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"net/smtp"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var tpl *template.Template
var x int = 0
var logs bool=false
// client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
// if err != nil {
//                 panic(err)
//         
func init() {
	tpl, _ = template.ParseGlob("templates/*.html")
}
func handleregister(w http.ResponseWriter, r *http.Request) {
	if r.Method=="POST"{
		logs=true
		em:=r.FormValue("email")
		ps:=r.FormValue("password")
		ls:=r.FormValue("number")
		if em!="" && ps!="" && ls!=""{
			add(em,ps)
			from := "mywork.p98@gmail.com"
			password := "gallardo.98"
			to := []string{
				r.FormValue("email"),
			}
			smtpHost := "smtp.gmail.com"
			smtpPort := "587"
			x = rand.Intn(2000000)
			z := strconv.Itoa(x)
			message := []byte(z)
			auth := smtp.PlainAuth("", from, password, smtpHost)
			smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
			fmt.Println("email is sent")
			http.Redirect(w, r, "/codepage", http.StatusSeeOther)
		}else{
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}

	}else{
	tpl.ExecuteTemplate(w, "register.html", nil)
	}
}
func codepage(w http.ResponseWriter, r *http.Request) {
	if logs{
		if r.Method=="POST"{
				code := r.FormValue("code")
				fmt.Println("dafa end point hit:::")
				fmt.Println(code)
				if code!=""{
					http.Redirect(w, r, "/login", http.StatusSeeOther)
				}else{
					http.Redirect(w, r, "/codepage", http.StatusSeeOther)
				}
		}else{
			tpl.ExecuteTemplate(w, "code.html", nil)
		}
	}else{
		http.Redirect(w, r, "/", http.StatusSeeOther)
		
	}
}
func add(em string,ps string){
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://mongo:gallardo@cluster0.l6jxp.mongodb.net/test"))
		if err != nil {
                panic(err)
        }
		collection := client.Database("GoDb").Collection("users")
		oneDoc := MongoFields{
		Email: em,
		Password: ps,
		}
		collection.InsertOne(context.TODO(), oneDoc)
		fmt.Println("Inserted")
}
func  find(em string,ps string) bool{
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://mongo:gallardo@cluster0.l6jxp.mongodb.net/test"))
		if err != nil {
                panic(err)
        }
		collection := client.Database("GoDb").Collection("users")
		results,err:=collection.Find(context.TODO(), bson.M{})
		if err != nil {
  			log.Fatal(err)
		}
		var perd []bson.M
		if err=results.All(context.TODO(),&perd);err!=nil{
			log.Fatal(err)
		}
		for _,docs := range perd{
			if docs["email"]==em && docs["password"]==ps{
				return true
			}
		}
		return false
		// if result.Email==em{
		// fmt.Println("finded")
		// }else{
		// 	fmt.Println("could not find")
		// }
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method=="POST"{
		em:=r.FormValue("email")
		ps:=r.FormValue("password")
		if em!="" && ps!=""{
			if find(em,ps){
			http.Redirect(w, r, "/dash", http.StatusSeeOther)}else{
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			}
		}else{				
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
		}else{
		tpl.ExecuteTemplate(w, "login.html", nil)
		}
}
func dash(w http.ResponseWriter, r *http.Request){
	tpl.ExecuteTemplate(w,"welcome.html",nil)
}
func handleRequests() {
	r := http.NewServeMux()
	r.HandleFunc("/", handleregister)
	r.HandleFunc("/codepage", codepage)
	r.HandleFunc("/login", login)
	r.HandleFunc("/dash", dash)
	log.Fatal(http.ListenAndServe(":5000", r))
}
type MongoFields struct {
Email string `json:"email"`
Password string `json:"password"`
}
func main() {
	handleRequests()
}
