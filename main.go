package main

import (
    "fmt"
    "net/http"
    "path/filepath"
    "html/template"
    "os"
    "log"
)

func Home(w http.ResponseWriter, r *http.Request) {
//  lp := filepath.Join("templates","base.html")
//fp := filepath.Join("templates", "index.html")
fp := filepath.Join("index.html")
	t, err := template.ParseFiles(fp)
	if err!=nil{
		fmt.Println("it's here1")
		log.Fatalln(err)
	}

err =  t.ExecuteTemplate(w, "index.html", "")
if err != nil {
		fmt.Println("it's here2")
	log.Fatalln(err)
}
}

func main() {
   mux := http.NewServeMux()

	fmt.Println("Listening...")
	mux.HandleFunc("/", Home)
    http.ListenAndServe(GetPort(), mux)
    
	}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Info: No port detected in the environment, defaulting to :" + port)

	return ":" + port
}

