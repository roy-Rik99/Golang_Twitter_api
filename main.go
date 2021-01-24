package main

import (
	"GolangTwitterapi/twitterapi"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var (
	tpl  *template.Template
	user twitterapi.Userprofile
)

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}
func checklogin(uname string) int {
	if user.Username == uname {
		return 0
	}
	return 1

}
func main() {

	http.HandleFunc("/", index)

	http.HandleFunc("/profile", profile) //base-->login-->profile

	http.HandleFunc("/register", register)        //base-->register
	http.HandleFunc("/register/welcome", welcome) //base-->register-->welcome-->profile

	http.HandleFunc("/twitteregister", twitteregister)         //base-->twitteregister-->welcome
	http.HandleFunc("/twitteregister/welcome", twitterwelcome) //base-->twitteregister-->welcome-->profile

	//http.HandleFunc("/try", test)
	log.Fatalln((http.ListenAndServe("177.186.149.2:8080", nil)))
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}
func profile(w http.ResponseWriter, r *http.Request) {
	var err int
	usrname := r.FormValue("usrname")
	if usrname == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		fmt.Println("\nALERT :--> USER-Name Field is Empty!\n\tPlease enter valid USER-Name.")
		return
	}
	user, err = twitterapi.Viewprofile(usrname)
	if err != 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		fmt.Println("\nALERT :--> USER NOT FOUND")
		return
	}

	tpl.ExecuteTemplate(w, "profile.html", user)
}
func register(w http.ResponseWriter, r *http.Request) {
	var err int
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	usrname := r.FormValue("usrname")
	if usrname == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		fmt.Println("\nALERT :--> USER-Name Field is Empty!\n\tPlease enter valid USER-Name.")
		return
	}

	_, err = twitterapi.Viewprofile(usrname)
	if err == 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		fmt.Println("\nALERT :--> USER-Name already exists!")
		return
	}
	user.Username = usrname

	tpl.ExecuteTemplate(w, "register.html", user)
}
func welcome(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	email := r.FormValue("email")
	gender := r.FormValue("gender")
	url := r.FormValue("url")
	desc := r.FormValue("status")
	loc := r.FormValue("location")

	user.Fullname = name
	user.Emailid = email
	user.Gender = gender
	user.URL = url
	user.Status = desc
	user.Location = loc
	user.Twitterlinked = "NO"
	twitterapi.CreateNewUserProfile(user) //function to push to database!
	tpl.ExecuteTemplate(w, "welcome.html", user)

}

func twitteregister(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	usrname := r.FormValue("usrname")
	if usrname == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		fmt.Println("\nALERT :--> USER-Name Field is Empty!\n\tPlease enter valid USER-Name.")
		return
	}
	var err int
	_, err = twitterapi.Viewprofile(usrname)
	if err == 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		fmt.Println("\nALERT :--> USER-Name already exists!")
		return
	}

	cred, err := twitterapi.TwitterCredbyUName(usrname)
	if err != 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		fmt.Println("\nALERT :--> USER Credentials NOT FOUND!")
		return
	}
	//fmt.Printf("%v :- ", usrname)
	//fmt.Println(cred)
	tinfo := twitterapi.RequestUserDetails(cred)
	user.ID = tinfo.TwitterID
	user.Username = tinfo.ScreenName
	user.Fullname = tinfo.Name
	user.URL = tinfo.URL
	user.Status = tinfo.Desc
	user.Location = tinfo.Location
	user.Twitterlinked = "YES"
	tpl.ExecuteTemplate(w, "twitter_register.html", user)
}

func twitterwelcome(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	gender := r.FormValue("gender")

	user.Emailid = email
	user.Gender = gender

	twitterapi.CreateNewUserProfile(user) //function to push to database!
	tpl.ExecuteTemplate(w, "welcome.html", user)

}
