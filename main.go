package main

import (
	"log"
	"meme-db/app"
	"net/http"
)

func serveWrapper(config app.DBConfig, serveFunc func(app.DBConfig, http.ResponseWriter, *http.Request)) http.Handler {
	db, err := app.ConnectDB(config)
	if err != nil {
		log.Println(err.Error())
	}

	return http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		serveFunc(db, w, r)
	})
}

func main() {
	// TODO @jeff: load in yaml config
	config := app.DBConfig{
		User:     "postgres",
		Password: "meme",
		Host:     "35.192.115.150",
		DBName:   "meme-db",
		SSLMode:  false,
	}

	// add func to handle url request
	http.HandleFunc("/get_meme_details", serveWrapper(config, app.GetMemeDetails))
	http.HandleFunc("/get_trending_memes", serveWrapper(config, app.GetTrendingMemes))
	http.HandleFunc("/get_meme_without_tags", serveWrapper(config, app.GetMemeWithoutTags))
	http.HandleFunc("/insert_meme_without_tags", serveWrapper(config, app.InsertMemeWithoutTags))
	http.HandleFunc("/insert_meme_abouts_and_tags", serveWrapper(config, app.InsertMemeAboutsAndTags))
	http.HandleFunc("/increment_meme_click", serveWrapper(config, app.IncrementMemeClick))

	// listen and serve
	if err := http.ListenAndServe(":8070", nil); err != nil {
		panic(err)
	}
}
