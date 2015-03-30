package main

import (
    "fmt"
    "log"
    "net/http"
    "os" 
)

func main() {
//    f, err := os.OpenFile("D:\\home\\site\\wwwroot\\testlogfile", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
    f, err := os.OpenFile("testlogfile", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
    
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Error: %v", err)
    })

    if err != nil {
        http.ListenAndServe(":"+os.Getenv("HTTP_PLATFORM_PORT"), nil)
    } else {
        defer f.Close()
        log.SetOutput(f)
        log.Println("--->   UP @ " + os.Getenv("HTTP_PLATFORM_PORT") +"  <------")
    }

    http.ListenAndServe(":"+os.Getenv("HTTP_PLATFORM_PORT"), nil)
}

