package main

import (
  "fmt"
  "log"
  "net/http"
)


func formHandle(w http.ResponseWriter , r *http.Request){
  if err := r.ParseForm(); err != nil{
    fmt.Fprintf(w, "ParseForm() err: %v", err)
    return
  }

  fmt.Fprintf(w, "Post request successful\n")
  name := r.FormValue("name")
  address := r.FormValue("address")

  fmt.Fprintf(w, "Name = %s\n", name)
  fmt.Fprintf(w, "Address= %s\n", address)
  
}

func helloHandler(w http.ResponseWriter , r *http.Request){
  if r.URL.Path != "/hello" {
    http.Error(w, "404 not found ", http.StatusNotFound)
    return
  }

  if r.Method != "GET"{
    http.Error(w, "Method not found", http.StatusNotFound)
    return
  }

  fmt.Fprintf(w, "hello!")


}

func main(){
  // create a http file server to serve static html files.
  fileServer := http.FileServer(http.Dir("./static"))
  http.Handle("/",fileServer)
  http.HandleFunc("/form",formHandle)
  http.HandleFunc("/hello", helloHandler)

  // server starting messege
  fmt.Println("Starting server at port 8080")
   err := http.ListenAndServe(":8080",nil)
  if err != nil {
    log.Fatal(err)
  }
}


