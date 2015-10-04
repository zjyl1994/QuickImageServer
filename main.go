package main

import(
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/mux"
)
	
func main(){
	fmt.Println("Quick Image Server.")
	fmt.Println("Author:zjyl1994@outlook.com.")
	LoadConf()
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/",UploadHandler).Methods("POST")
	r.HandleFunc("/{imgid}",DownloadHandler).Methods("GET")
    err := http.ListenAndServe(conf.ListenAddr, r)
    if err != nil {
        log.Fatal("ListenAndServe error: ", err)
    }
}
