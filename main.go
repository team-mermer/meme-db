package main

import (
	"database/sql"
	"io/ioutil"
	"log"
	"meme-db/app"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

func serveWrapper(config app.DBConfig, serveFunc func(*sql.DB, http.ResponseWriter, *http.Request)) http.Handler {
	db, err := app.ConnectDB(config)
	if err != nil {
		log.Println(err.Error())
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		serveFunc(db, w, r)
	})
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	var config app.DBConfig
	if err := yaml.Unmarshal([]byte(data), &config); err != nil {
		log.Fatalf("error: %v", err)
	}

	// add func to handle url request
	http.Handle("/get_meme_details", serveWrapper(config, app.GetMemeDetails))
	http.Handle("/get_trending_memes", serveWrapper(config, app.GetTrendingMemes))
	http.Handle("/get_meme_without_tags", serveWrapper(config, app.GetMemeWithoutTags))
	http.Handle("/insert_meme_without_tags", serveWrapper(config, app.InsertMemeWithoutTags))
	http.Handle("/insert_meme_abouts_and_tags", serveWrapper(config, app.InsertMemeAboutsAndTags))
	http.Handle("/increment_meme_click", serveWrapper(config, app.IncrementMemeClick))

	// listen and serve
	if err := http.ListenAndServe(":8070", nil); err != nil {
		panic(err)
	}
}
