package router

import (
	"encoding/json"
	"log"
	"net/http"
	"service"
)

func UserRouter(serveMux *http.ServeMux){
	serveMux.HandleFunc("/api/data", func(rw http.ResponseWriter, r *http.Request){
        rw.WriteHeader(http.StatusCreated)
        rw.Header().Set("Content-Type", "application/json")

        response, err := json.Marshal(service.GetUser())

        if err != nil{
            log.Fatal(err)
        }

        rw.Write([]byte(response))
    })
}