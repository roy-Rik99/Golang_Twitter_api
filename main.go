package main

import (
	"GolangTwitterapi/twitterapi"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template
var usrprofile twitterapi.Userprofile

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}
func freeusrprofile(data *twitterapi.Userprofile) {
	data.ID = 0
	data.Username = ""
	data.Fullname = ""
	data.Emailid = ""
	data.Gender = ""
	data.URL = ""
	data.Status = ""
	data.Location = ""
	data.Twitterlinked = ""
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}
func profile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("profile")
	usrname := r.FormValue("usrname")
	if usrname == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		fmt.Println("\nALERT :--> USER-Name Field is Empty!\n\tPlease enter valid USER-Name.")
		return
	}
	user, err := twitterapi.Viewprofile(usrname)
	if err != 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		fmt.Println("\nALERT :--> USER NOT FOUND")
		return
	}

	tpl.ExecuteTemplate(w, "profile.html", user)
}
func editprofile(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	var user twitterapi.Userprofile
	user.Username = r.FormValue("usrname")

	tpl.ExecuteTemplate(w, "edit_profile.html", user)
}
func savechanges(w http.ResponseWriter, r *http.Request) {

	var newuser twitterapi.Userprofile
	newuser.Username = r.FormValue("usrname")
	newuser.Fullname = r.FormValue("name")
	newuser.Emailid = r.FormValue("email")
	newuser.Gender = r.FormValue("gender")
	newuser.URL = r.FormValue("url")
	newuser.Status = r.FormValue("status")
	newuser.Location = r.FormValue("location")

	twitterapi.UpdateUserProfile(newuser) //function to push to database!
	tpl.ExecuteTemplate(w, "profileredirect.html", newuser)
}
func removeaccount(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/profile", http.StatusSeeOther)
		return
	}
	usrname := r.FormValue("usrname")
	err := twitterapi.Removeprofile(usrname)
	if err != nil {
		fmt.Printf("\nALERT :--> %v\n", err)
		http.Redirect(w, r, "/profile", http.StatusNotFound)
		return
	}
	fmt.Printf("\nALERT :--> User %v has been removed!\n", usrname)
	tpl.ExecuteTemplate(w, "delete_profile.html", nil)
}
func register(w http.ResponseWriter, r *http.Request) {
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

	_, err := twitterapi.Viewprofile(usrname)
	if err == 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		fmt.Println("\nALERT :--> USER-Name already exists!")
		return
	}
	var user twitterapi.Userprofile
	user.Username = usrname

	tpl.ExecuteTemplate(w, "register.html", user)
}
func welcome(w http.ResponseWriter, r *http.Request) {
	var user twitterapi.Userprofile
	user.Username = r.FormValue("usrname")
	user.Fullname = r.FormValue("name")
	user.Emailid = r.FormValue("email")
	user.Gender = r.FormValue("gender")
	user.URL = r.FormValue("url")
	user.Status = r.FormValue("status")
	user.Location = r.FormValue("location")
	user.Twitterlinked = "NO"
	twitterapi.CreateNewUserProfile(user) //function to push to database!
	tpl.ExecuteTemplate(w, "welcome.html", user)
}
func twitteregister(w http.ResponseWriter, r *http.Request) {
	fmt.Println("twitterregister")
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	usrname := r.FormValue("usrname")
	if usrname == "" {
		http.Redirect(w, r, "/", http.StatusBadRequest)
		fmt.Println("\nALERT :--> USER-Name Field is Empty!\n\tPlease enter valid USER-Name.")
		return
	}
	_, err := twitterapi.Viewprofile(usrname)
	if err == 0 {
		http.Redirect(w, r, "/", http.StatusLocked)
		fmt.Println("\nALERT :--> USER-Name already exists!")
		return
	}

	cred, err := twitterapi.TwitterCredbyUName(usrname)
	if err != 0 {
		http.Redirect(w, r, "/", http.StatusNoContent)
		fmt.Println("\nALERT :--> USER Credentials NOT FOUND!")
		return
	}

	tinfo := twitterapi.RequestUserDetails(cred)

	usrprofile.ID = tinfo.TwitterID
	usrprofile.Username = tinfo.ScreenName
	usrprofile.Fullname = tinfo.Name
	usrprofile.URL = tinfo.URL
	usrprofile.Status = tinfo.Desc
	usrprofile.Location = tinfo.Location
	usrprofile.Twitterlinked = "YES"

	tpl.ExecuteTemplate(w, "twitter_register.html", usrprofile)
}
func twitterwelcome(w http.ResponseWriter, r *http.Request) {
	usrprofile.Emailid = r.FormValue("email")
	usrprofile.Gender = r.FormValue("gender")

	twitterapi.CreateNewUserProfile(usrprofile) //function to push to database!
	tpl.ExecuteTemplate(w, "welcome.html", usrprofile)
	freeusrprofile(&usrprofile)
}

func main() {
	http.HandleFunc("/", index)

	http.HandleFunc("/profile", profile)                             //base-->login-->profile
	http.HandleFunc("/profile/editprofile", editprofile)             //base-->login-->profile-->editprofile
	http.HandleFunc("/profile/editprofile/savechanges", savechanges) //base-->login-->profile-->editprofile
	http.HandleFunc("/profile/removeprofile", removeaccount)         //base-->login-->profile-->removeprofile

	http.HandleFunc("/register", register)        //base-->register
	http.HandleFunc("/register/welcome", welcome) //base-->register-->welcome-->profile

	http.HandleFunc("/twitteregister", twitteregister)         //base-->twitteregister-->welcome
	http.HandleFunc("/twitteregister/welcome", twitterwelcome) //base-->twitteregister-->welcome-->profile

	//http.HandleFunc("/try", test)
	log.Fatalln((http.ListenAndServe("177.186.149.2:8080", nil)))
}
