package app

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func connectDB() (*sql.DB, error) {
	connectString := `
		user=postgres 
		password=meme 
		host=35.192.115.150 
		dbname=meme-db 
		sslmode=disable 
	`
	db, openErr := sql.Open("postgres", connectString)
	if openErr != nil {
		log.Fatal(openErr)
		return db, openErr
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
		return db, err
	}

	return db, nil
}

func getMemesByIds(db *sql.DB, memeIds []int) ([]memeDetail, error) {
	var memes []memeDetail

	strMemeIds := make([]string, len(memeIds))
	for i, memeID := range memeIds {
		strMemeIds[i] = strconv.Itoa(memeID)
	}
	sqlFmtStr := strings.Join(strMemeIds, ",")
	sqlQuery := fmt.Sprintf(
		`
		SELECT
			meme.id,
			meme.title,
			meme.image_path,
			meme.about,
			tag.name
		FROM
			meme
		LEFT JOIN
			meme_tag
		ON
			meme.id = meme_tag.meme_id
		INNER JOIN
			tag
		ON
			meme_tag.tag_id = tag.id
		WHERE
			meme.id IN(%s)
		`,
		sqlFmtStr)

	rows, queryErr := db.Query(sqlQuery)
	if queryErr != nil {
		log.Print(queryErr)
		return memes, queryErr
	}
	defer rows.Close()

	IDToMemeMap := make(map[int]memeDetail)

	for rows.Next() {
		var (
			id    int
			title string
			path  string
			about string
			tag   string
		)
		if err := rows.Scan(&id, &title, &path, &about, &tag); err != nil {
			log.Fatal(err)
			return memes, err
		}

		meme, exists := IDToMemeMap[id]
		if exists {
			meme.Tags = append(meme.Tags, tag)
		} else {
			tags := []string{tag}
			meme := memeDetail{
				ID:       id,
				Title:    title,
				ImageURL: path,
				About:    about,
				Tags:     tags,
			}
			IDToMemeMap[id] = meme
		}
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
		return memes, err
	}

	for _, meme := range IDToMemeMap {
		memes = append(memes, meme)
	}

	return memes, nil
}

func getMemesWithoutAbout(db *sql.DB, limit int) ([]memeDetail, error) {
	var memes []memeDetail

	sqlQuery := fmt.Sprintf(
		`
		SELECT
			id,
			title,
			image_path
		FROM
			meme
		WHERE
			about IS NULL
		LIMIT %d
		`,
		limit)

	var rows *sql.Rows
	rows, queryErr := db.Query(sqlQuery)
	if queryErr != nil {
		log.Print(queryErr)
		return memes, queryErr
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id    int
			title string
			path  string
		)
		if err := rows.Scan(&id, &title, &path); err != nil {
			log.Fatal(err)
			return memes, err
		}

		meme := memeDetail{
			ID:       id,
			Title:    title,
			ImageURL: path,
			About:    "",
			Tags:     nil,
		}

		memes = append(memes, meme)
	}

	return memes, nil
}
