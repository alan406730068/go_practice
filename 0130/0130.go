package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

var store *sessions.CookieStore

func init() {
	store = sessions.NewCookieStore([]byte("secret-key"))
}

func main() {
	http.HandleFunc("/home", home)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	fmt.Println("session server run on port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["auth"] = nil
	session.Options.MaxAge = -1
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "logged out.")
}

func login(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if session.IsNew {
		//session.Options.MaxAge = -1
		session.Values["auth"] = true
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintln(w, "logged in")
	} else {
		http.Redirect(w, r, "/home", 301)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if session.IsNew {
		http.Redirect(w, r, "/login", 301)
	} else {
		auth := session.Values["auth"]
		if auth != nil {
			isAuth, ok := auth.(bool)
			if ok && isAuth {
				fmt.Fprintln(w, "Home Page")
			} else {
				http.Error(w, "unauthorizeed", http.StatusUnauthorized)
				return
			}
		} else {
			http.Error(w, "unauthorizeed", http.StatusUnauthorized)
			return
		}
	}

}
