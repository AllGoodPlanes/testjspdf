package main

import (

    "net/http"
    "os"
    "fmt"
"path/filepath"
"html/template"
"log"
)

func main() {

    mux := http.NewServeMux()
	mux.HandleFunc("/", Pdf)
    wpdf:= http.FileServer(http.Dir("web"))
	mux.Handle("/web/",http.StripPrefix("/web", wpdf))
    bpdf:= http.FileServer(http.Dir("build"))
	mux.Handle("/build/",http.StripPrefix("/build", bpdf))
    img:= http.FileServer(http.Dir("images"))
	mux.Handle("/images/",http.StripPrefix("/images", img))
     
    http.ListenAndServe(GetPort(), mux)

    }

func Pdf(w http.ResponseWriter, req *http.Request) {

      lp := filepath.Join("mypage.html")

	t, err := template.ParseFiles(lp)
	if err!=nil{
		fmt.Println("it's here1")
		log.Fatalln(err)
	}

err =  t.ExecuteTemplate(w, "mypage.html", "")
if err != nil {
		fmt.Println("it's here2")
	log.Fatalln(err)
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


