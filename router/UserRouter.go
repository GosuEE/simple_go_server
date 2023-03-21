package router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"requestDTO"
	"service"
)

func UserRouter(serveMux *http.ServeMux){
	serveMux.HandleFunc("/api/user", func(rw http.ResponseWriter, r *http.Request){
        rw.WriteHeader(http.StatusCreated)
        rw.Header().Set("Content-Type", "application/json")
        switch r.Method{
        case http.MethodGet:
            response, err := json.Marshal(service.GetUser())

            if err != nil{
              log.Fatal(err)
            }
            rw.Write([]byte(response))
        case http.MethodPatch:
            u := requestDTO.UserRequestDTO{}
            json.NewDecoder(r.Body).Decode(&u)
            service.UpdateUser(u)
        case http.MethodPost:
            u := requestDTO.UserRequestDTO{}
            json.NewDecoder(r.Body).Decode(&u)
            service.AddUser(u)
        case http.MethodDelete:
            u := requestDTO.UserRequestDTO{}
            json.NewDecoder(r.Body).Decode(&u)
            fmt.Println(u)
            service.DeleteUser(u)
        }
    })
}