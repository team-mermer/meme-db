package app

import (
	"encoding/json"
	"log"
	"net/http"
)

// GetMemeDetails api func to return meme details
func GetMemeDetails(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Access-Control-Allow-Origin", "*")

	db, connectErr := connectDB()
	if connectErr != nil {
		log.Println(connectErr.Error())
		http.Error(w, "connect db error", http.StatusBadRequest)
	}
	memeIds := []int{2, 3}
	// memes, getMemeErr := getMemesByIds(db, memeIds)
	memes, _ := getMemesByIds(db, memeIds)
	jsonString, _ := json.Marshal(memes)

	if _, err := w.Write(jsonString); err != nil {
		log.Println(err.Error())
		log.Printf("json content:\n %s\n", jsonString)
		http.Error(w, "can't write json string to response", http.StatusBadRequest)
	}
}
