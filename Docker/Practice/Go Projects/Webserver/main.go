package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("hello")

	fileserver := http.FileServer(http.Dir("./Static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/home", hellomethod)
	http.HandleFunc("/form", formhandler)

	port := "8080"
	fmt.Printf("Server is running on %v", port)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}

func hellomethod(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/home" {
		http.Error(w, "404 NOT FOUND", http.StatusNotFound)

	}

	if r.Method != "GET" {
		http.Error(w, "Method is not found", http.StatusNotFound)
	}

	fmt.Fprint(w, "Hello this a route to /hello")
}

func formhandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Found an err %s\n", err)
		return
	}

	fmt.Fprintln(w, "POST request is Successfull")
	name := r.FormValue("name")
	email := r.FormValue("email")
	number := r.FormValue("number")

	fmt.Fprintf(w, "name is %s\n", name)
	fmt.Fprintf(w, "email is %s\n", email)
	fmt.Fprintf(w, "Number is %s\n", number)

	fmt.Fprintf(w, "conformation is send on %s to %s. if any queries contact via %s", email, name, number)
	// jsondata := json.Marshaler([name,email,number],err)

	// fmt.Printf(jsondata)
}
