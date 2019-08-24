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

// GetMemeWithoutTags api func to return meme details which has no tags
func GetMemeWithoutTags(w http.ResponseWriter, r *http.Request) {
	db, connectErr := connectDB()
	if connectErr != nil {
		log.Println(connectErr.Error())
		http.Error(w, "connect db error", http.StatusBadRequest)
	}

	limit := 100
	memes, _ := getMemesWithoutAbout(db, limit)
	jsonString, _ := json.Marshal(memes)

	if _, err := w.Write(jsonString); err != nil {
		log.Println(err.Error())
		log.Printf("json content:\n %s\n", jsonString)
		http.Error(w, "can't write json string to response", http.StatusBadRequest)
	}
}

// InsertMemeWithoutTags api to insert meme's basic info besides tags and about
func InsertMemeWithoutTags(w http.ResponseWriter, r *http.Request) {
	db, connectErr := connectDB()
	if connectErr != nil {
		log.Println(connectErr.Error())
		http.Error(w, "connect db error", http.StatusBadRequest)
	}

	memes := []memeDetail{
		memeDetail{
			Title:    "corgi-1",
			ImageURL: "http://placecorgi.com/600/600",
			About:    "",
			Tags:     nil,
		},
		memeDetail{
			Title:    "corgi-2",
			ImageURL: "http://placecorgi.com/600/600",
			About:    "",
			Tags:     nil,
		},
	}

	if err := insertMemes(db, memes); err != nil {
		log.Println(err.Error())
		http.Error(w, "fail to insert memes without tags", http.StatusBadRequest)
	}
}

// InsertMemeAboutsAndTags api to insert meme's advanced info, i.e. tags and about
func InsertMemeAboutsAndTags(w http.ResponseWriter, r *http.Request) {
	db, connectErr := connectDB()
	if connectErr != nil {
		log.Println(connectErr.Error())
		http.Error(w, "connect db error", http.StatusBadRequest)
	}

	memes := []memeDetail{
		memeDetail{
			Title:    "corgi-1",
			ImageURL: "http://placecorgi.com/600/600",
			About:    "",
			Tags:     nil,
		},
		memeDetail{
			Title:    "corgi-2",
			ImageURL: "http://placecorgi.com/600/600",
			About:    "",
			Tags:     nil,
		},
	}

	if err := insertMemeAbouts(db, memes); err != nil {
		log.Println(err.Error())
		http.Error(w, "fail to insert meme abouts", http.StatusBadRequest)
	}

	if err := insertMemeTags(db, memes); err != nil {
		log.Println(err.Error())
		http.Error(w, "fail to insert meme tags", http.StatusBadRequest)
	}
}

// IncrementClick api to increment meme's click
func IncrementClick(w http.ResponseWriter, r *http.Request) {
	db, connectErr := connectDB()
	if connectErr != nil {
		log.Println(connectErr.Error())
		http.Error(w, "connect db error", http.StatusBadRequest)
	}

	memeIDs := []int{1, 2, 3}

	if err := incrementClick(db, memeIDs); err != nil {
		log.Println(err.Error())
		http.Error(w, "fail to increment meme click", http.StatusBadRequest)
	}
}
