package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Request struct {
	Player1 string `json:"player1"`
	Player2 string `json:"player2"`
}

type Response struct {
	Result string `json:"result"`
}

func main() {
	http.HandleFunc("/play", playHandler)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

func playHandler(w http.ResponseWriter, r *http.Request) {
	var req Request
	var res Response
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	player1 := conversion(req.Player1)
	player2 := conversion(req.Player2)
	comparision := Comparision(player1, player2)
	res.Result = comparision
	//w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
