package main

import (
	"meme-db/app"
	"net/http"
)

func main() {
	// add func to handle url request
	http.HandleFunc("/get_meme_details", app.GetMemeDetails)

	// listen and serve
	if err := http.ListenAndServe(":8070", nil); err != nil {
		panic(err)
	}
}
