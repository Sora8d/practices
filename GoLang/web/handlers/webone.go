package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Println(strings.Repeat("-", 20))
	fmt.Fprintf(w, "Hello astaxie!")
}

func BasicHandle() {
	http.HandleFunc("/", sayhelloName)
	err := http.ListenAndServe(":9092", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// More complex handlers
//http.HandleFunc("/complete/", views.CompleteTaskFunc)
//http.HandleFunc("/delete/", views.DeleteTaskFunc)

func ComplexHandle() {
	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/es", ShowAllTasksFunc)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9006", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func ShowAllTasksFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, "GET i suppose")
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form["assd"])
	if r.Method == "GET" {
		t, _ := template.ParseFiles("templates/login.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println("username: ", r.Form["username"])
		fmt.Println("password: ", r.Form["password"])
	}
}
