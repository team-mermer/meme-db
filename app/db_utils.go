package app

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"

	_ "github.com/lib/pq" // postgres driver
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
	strIds := strings.Join(strMemeIds, ",")
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
		strIds)

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

// TODO @jeff: try to refactor this with getMemeByIDs
func getTopClickedMemes(db *sql.DB, numOfTop int) ([]memeDetail, error) {
	var memes []memeDetail

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
		ORDER BY
			meme.click
		DESC
		LIMIT %d
		`,
		numOfTop)

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

func insertMemes(db *sql.DB, memes []memeDetail) error {
	sqlStr := "INSERT INTO meme(title, image_path) VALUES "
	vals := []interface{}{}

	for _, meme := range memes {
		sqlStr += "(?, ?),"
		vals = append(vals, meme.Title, meme.ImageURL)
	}
	sqlStr = sqlStr[0 : len(sqlStr)-2]
	stmt, _ := db.Prepare(sqlStr)

	//format all vals at once
	if _, err := stmt.Exec(vals...); err != nil {
		log.Printf("DB statement execution with error: %v\n", err)
		return err
	}

	return nil
}

func insertMemeAbouts(db *sql.DB, memes []memeDetail) error {
	sqlStr := "INSERT INTO meme(about) VALUES "
	vals := []interface{}{}

	for _, meme := range memes {
		sqlStr += "(?,),"
		vals = append(vals, meme.About)
	}
	sqlStr = sqlStr[0 : len(sqlStr)-2]
	stmt, _ := db.Prepare(sqlStr)

	//format all vals at once
	if _, err := stmt.Exec(vals...); err != nil {
		log.Printf("DB statement execution with error: %v\n", err)
		return err
	}

	return nil
}

func insertMemeTags(db *sql.DB, memes []memeDetail) error {
	newTagIDMap, err := insertAndGetNewTagIDMap(db, memes)
	if err != nil {
		log.Printf("insert and get new tag id map with error: %v\n", newTagIDMap)
		return err
	}

	memeIDs := make([]int, 0)
	tagIDs := make([]int, 0)
	for _, meme := range memes {
		for _, tag := range meme.Tags {
			id, _ := newTagIDMap[tag]
			memeIDs = append(memeIDs, meme.ID)
			tagIDs = append(tagIDs, id)
		}
	}

	sqlStr := "INSERT INTO meme_tag(meme_id, tag_id) VALUES "
	vals := []interface{}{}

	for i, memeID := range memeIDs {
		sqlStr += "(?, ?),"
		vals = append(vals, memeID, tagIDs[i])
	}
	stmt, _ := db.Prepare(sqlStr)
	if _, err := stmt.Exec(vals...); err != nil {
		log.Printf("DB statement execution with error: %v\n", err)
		return err
	}

	return nil
}

func insertAndGetNewTagIDMap(db *sql.DB, memes []memeDetail) (map[string]int, error) {
	newTagIDMap := make(map[string]int)

	tags := make([]string, 0)
	for _, meme := range memes {
		tags = append(tags, meme.Tags...)
	}
	sqlQuery := fmt.Sprintf("SELECT id FROM tag WHERE tag.name IN(%s)", strings.Join(tags, ","))
	rows, err := db.Query(sqlQuery)
	if err != nil {
		log.Printf("fail to insert and get new tag ids: %v\n", err)
		return newTagIDMap, err
	}
	defer rows.Close()

	existingTags := make(map[string]bool)
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Printf("fail to scan row: %v\n", err)
			return newTagIDMap, err
		}
		existingTags[name] = true
	}
	if err := rows.Err(); err != nil {
		log.Printf("error occurs when iterating through rows: %v\n", err)
		return newTagIDMap, err
	}

	newTags := make([]string, 0)
	for _, tag := range tags {
		_, ok := existingTags[tag]
		if !ok {
			newTags = append(newTags, tag)
		}
	}

	// insert new tags
	insertSQL := "INSERT INTO tag(name) VALUES "
	vals := []interface{}{}
	for _, tag := range newTags {
		insertSQL += "(?,),"
		vals = append(vals, tag)
	}
	insertSQL = insertSQL[0 : len(insertSQL)-2]
	stmt, _ := db.Prepare(insertSQL)

	if _, err := stmt.Exec(vals...); err != nil {
		log.Printf("DB statement execute with error: %v\n", err)
		return newTagIDMap, err
	}

	// get new tag ids
	getNewIDSQL := fmt.Sprintf(
		"SELECT id, name FROM tag WHERE name IN(%s)",
		strings.Join(newTags, ","),
	)

	rows, err = db.Query(getNewIDSQL)
	if err != nil {
		log.Printf("DB query with error: %v\n", err)
		return newTagIDMap, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id   int
			name string
		)
		if err := rows.Scan(&id, &name); err != nil {
			log.Printf("fail to iterate rows with error: %v\n", err)
			return newTagIDMap, err
		}
		newTagIDMap[name] = id
	}

	return newTagIDMap, nil
}
