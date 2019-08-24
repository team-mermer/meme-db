package app

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

// GetMemeDetails api func to return meme details
func GetMemeDetails(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Access-Control-Allow-Origin", "*")

	decoder := json.NewDecoder(r.Body)
	var input memeIDInput
	if err := decoder.Decode(&input); err != nil {
		log.Println("cannot decode from request body")
		http.Error(w, "can't parse request body", http.StatusBadRequest)
	}

	memes, _ := getMemesByIds(db, input.IDs)
	jsonString, _ := json.Marshal(memes)

	if _, err := w.Write(jsonString); err != nil {
		log.Println(err.Error())
		log.Printf("json content:\n %s\n", jsonString)
		http.Error(w, "can't write json string to response", http.StatusBadRequest)
	}
}

// GetMemeWithoutTags api func to return meme details which has no tags
func GetMemeWithoutTags(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var input memeIDInput
	if err := decoder.Decode(&input); err != nil {
		log.Println("cannot decode from request body")
		http.Error(w, "can't parse request body", http.StatusBadRequest)
	}

	memes, _ := getMemesWithoutAbout(db, input.NumOfResult)
	jsonString, _ := json.Marshal(memes)

	if _, err := w.Write(jsonString); err != nil {
		log.Println(err.Error())
		log.Printf("json content:\n %s\n", jsonString)
		http.Error(w, "can't write json string to response", http.StatusBadRequest)
	}
}

// GetTrendingMemes api func to return trending meme details ordered by click
func GetTrendingMemes(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var input trendingInput
	if err := decoder.Decode(&input); err != nil {
		log.Println("cannot decode from request body")
		http.Error(w, "can't parse request body", http.StatusBadRequest)
	}

	memes, _ := getTopClickedMemes(db, input.NumOfResult)
	jsonString, _ := json.Marshal(memes)

	if _, err := w.Write(jsonString); err != nil {
		log.Println(err.Error())
		log.Printf("json content:\n %s\n", jsonString)
		http.Error(w, "can't write json string to response", http.StatusBadRequest)
	}
}

// InsertMemeWithoutTags api to insert meme's basic info besides tags and about
func InsertMemeWithoutTags(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var input memeDetailsInput
	if err := decoder.Decode(&input); err != nil {
		log.Println("cannot decode from request body")
		http.Error(w, "can't parse request body", http.StatusBadRequest)
	}

	if err := insertMemes(db, input.MemeDetails); err != nil {
		log.Println(err.Error())
		http.Error(w, "fail to insert memes without tags", http.StatusBadRequest)
	}
}

// InsertMemeAboutsAndTags api to insert meme's advanced info, i.e. tags and about
func InsertMemeAboutsAndTags(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var input memeDetailsInput
	if err := decoder.Decode(&input); err != nil {
		log.Println("cannot decode from request body")
		http.Error(w, "can't parse request body", http.StatusBadRequest)
	}

	if err := insertMemeAbouts(db, input.MemeDetails); err != nil {
		log.Println(err.Error())
		http.Error(w, "fail to insert meme abouts", http.StatusBadRequest)
	}

	if err := insertMemeTags(db, input.MemeDetails); err != nil {
		log.Println(err.Error())
		http.Error(w, "fail to insert meme tags", http.StatusBadRequest)
	}
}

// IncrementMemeClick api to increment meme's click
func IncrementMemeClick(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var input memeIDInput
	if err := decoder.Decode(&input); err != nil {
		log.Println("cannot decode from request body")
		http.Error(w, "can't parse request body", http.StatusBadRequest)
	}

	if err := incrementClick(db, input.IDs); err != nil {
		log.Println(err.Error())
		http.Error(w, "fail to increment meme click", http.StatusBadRequest)
	}
}
