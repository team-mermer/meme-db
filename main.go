package main

import (
	"meme-db/app"
	"net/http"
	"os"
)

func main() {
	config, err := app.GetDBConfig(os.Args[1])
	if err != nil {
		panic(err)
	}

	// add func to handle url request
	http.Handle("/get_meme_details", app.ServeWrapper(config, app.GetMemeDetails))
	http.Handle("/get_trending_memes", app.ServeWrapper(config, app.GetTrendingMemes))
	http.Handle("/get_meme_without_tags", app.ServeWrapper(config, app.GetMemeWithoutTags))
	http.Handle("/insert_meme_without_tags", app.ServeWrapper(config, app.InsertMemeWithoutTags))
	http.Handle("/insert_meme_abouts_and_tags", app.ServeWrapper(config, app.InsertMemeAboutsAndTags))
	http.Handle("/increment_meme_click", app.ServeWrapper(config, app.IncrementMemeClick))

	// listen and serve
	if err := http.ListenAndServe(":8070", nil); err != nil {
		panic(err)
	}
}
