package main

import(
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Message struct{
	Text string `json:"text"`
}

func main(){

	router := mux.NewRouter()
	router.HandleFunc("/",helloHandler).Methods("GET")
	router.HandleFunc("/api/greet",greetHandler).Methods("POST")

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080",router)

}


func helloHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to Gopher's Playground")
}

func greetHandler(w http.ResponseWriter, r *http.Request){
	var message Message 
	
	decoder:= json.NewDecoder(r.Body)
	err := decoder.Decode(&message)
	if err != nil{
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	response := Message{Text: "Greetings " + message.Text}

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(response)

}

