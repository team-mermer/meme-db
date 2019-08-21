package main

import (
	"meme-db/app"
	"net/http"
)

func main() {
	// add func to handle url request
	http.HandleFunc("/get_meme_details", app.GetMemeDetails)
	http.HandleFunc("/get_meme_without_tags", app.GetMemeWithoutTags)
	http.HandleFunc("/insert_meme_without_tags", app.InsertMemeWithoutTags)
	http.HandleFunc("/insert_meme_abouts_and_tags", app.InsertMemeAboutsAndTags)

	// listen and serve
	if err := http.ListenAndServe(":8070", nil); err != nil {
		panic(err)
	}
}
