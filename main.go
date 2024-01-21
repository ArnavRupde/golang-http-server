package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserCreationResp struct {
	Id     string `json:"id"`
	Status string `json:"status"`
}

func userHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		var user User
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&user)
		if err != nil {
			http.Error(
				w,
				"Failed to parse request body",
				http.StatusBadRequest,
			)
		}
		fmt.Printf("Name: %s Age: %d \n", user.Name, user.Age)
		userCreationResp := UserCreationResp{
			Id:     "abc",
			Status: "Created",
		}
		jsonResponseData, err := json.Marshal(userCreationResp)
		if err != nil {
			http.Error(
				w,
				"Failed to marshal response",
				http.StatusBadRequest,
			)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponseData)
	default:
		http.Error(
			w,
			"Method not supported",
			http.StatusMethodNotAllowed,
		)

	}
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Status: OK")
}
func main() {
	fmt.Println("Starting HTTP Server on port 8000")
	http.HandleFunc("/health", healthCheckHandler)
	http.HandleFunc("/user", userHandler)
	http.ListenAndServe(":8000", nil)
}
