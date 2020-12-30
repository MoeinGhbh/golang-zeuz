//package main
//
//import (
//	"fmt"
//	"io/ioutil"
//	"log"
//	"net/http"
//	"os"
//)
//
//func main() {
//	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
//
//	if err != nil {
//		fmt.Print(err.Error())
//		os.Exit(1)
//	}
//
//	responseData, err := ioutil.ReadAll(response.Body)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(string(responseData))
//
//}

package main

import (
	"html/template"
	"net/http"
	"os"
	"git.zeuz.io/sdk/golang/zeuzsdk"
)

var tpl = template.Must(template.ParseFiles("index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}


	zeuzsdk.IDGenerate(zeuzsdk.IDTypeInvalid)


	fs := http.FileServer(http.Dir("assets"))

	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)


}
