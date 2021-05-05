package main

import (
    "fmt"
    "net/http"
    "path/filepath"
    "html/template"
    "os"
    "log"
)

const STATIC_URL string = "/static/"
const STATIC_ROOT string = "static/"

func Home(w http.ResponseWriter, r *http.Request) {
  lp := filepath.Join("templates","base.html")
fp := filepath.Join("templates", "index.html")
	t, err := template.ParseFiles(lp, fp)
	if err!=nil{
		fmt.Println("it's here1")
		log.Fatalln(err)
	}

err =  t.ExecuteTemplate(w, "base", "")
if err != nil {
		fmt.Println("it's here2")
	log.Fatalln(err)
}
}

func main() {
   mux := http.NewServeMux()

	fmt.Println("Listening...")
	mux.HandleFunc("/", Home)
    fs:= http.FileServer(http.Dir("static"))
	mux.Handle("/static/",http.StripPrefix("/static/", fs))
    dn:= http.FileServer(http.Dir("pdf.js"))
	mux.Handle("/pdf.js/",http.StripPrefix("/pdf.js/", dn))
    http.ListenAndServe(GetPort(), mux)
}

func render(w http.ResponseWriter, tmpl string) {
	tmpl_list := []string{"templates/base.html",
		fmt.Sprintf("templates/%s.html", tmpl),
	}
	t, err := template.ParseFiles(tmpl_list...)
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err= t.ExecuteTemplate(w, "base", "")
	if err != nil {
		log.Print("template executing error: ", err)
	}

}


func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Info: No port detected in the environment, defaulting to :" + port)

	return ":" + port
}

