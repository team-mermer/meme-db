package app

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

const testConfigPath = "../test_config.yaml"

func testAPIUtils(
	method string,
	APIUrl string,
	handler http.Handler,
	t *testing.T,
	expectedString string,
	bodyBuffer *bytes.Buffer,
) {

	req, err := http.NewRequest(method, APIUrl, bodyBuffer)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf(
			"handler returned wrong status code: got %v want %v",
			status,
			http.StatusOK,
		)
	}

	APIBodyString := rr.Body.String()

	if APIBodyString != expectedString {
		t.Errorf(
			"handler returned unexpected body: got %v want %v",
			APIBodyString,
			expectedString,
		)
	}
}

func TestGetMemeDetails(t *testing.T) {
	// TODO @jeff: reset DB and insert example data
	config, err := GetDBConfig(testConfigPath)
	if err != nil {
		t.Fatal(err)
	}
	handler := ServeWrapper(config, GetMemeDetails)

	// expected result
	expectMemes := make([]memeDetail, 0)
	expectMemes = append(
		expectMemes,
		memeDetail{
			ID:       0,
			Title:    "chuan",
			ImageURL: "www.blabla.com/blabla.jpg",
			About:    "ark chuan",
		})
	expectedBytes, _ := json.Marshal(expectMemes)
	expectedString := string(expectedBytes)

	testAPIUtils(
		"GET",
		"/get_meme_details",
		handler,
		t,
		expectedString,
		nil)
}

func TestGetMemeWithoutTags(t *testing.T) {
	// TODO @jeff: reset DB and insert example data
	config, err := GetDBConfig(testConfigPath)
	if err != nil {
		t.Fatal(err)
	}
	handler := ServeWrapper(config, GetMemeWithoutTags)

	// expected result
	expectMemes := make([]memeDetail, 0)
	expectMemes = append(
		expectMemes,
		memeDetail{
			ID:       0,
			Title:    "chuan",
			ImageURL: "www.blabla.com/blabla.jpg",
			About:    "ark chuan",
		})
	expectedBytes, _ := json.Marshal(expectMemes)
	expectedString := string(expectedBytes)

	testAPIUtils(
		"GET",
		"/get_meme_without_tags",
		handler,
		t,
		expectedString,
		nil)
}

func TestGetTrendingMemes(t *testing.T) {
	// TODO @jeff: reset DB and insert example data
	config, err := GetDBConfig(testConfigPath)
	if err != nil {
		t.Fatal(err)
	}
	handler := ServeWrapper(config, GetTrendingMemes)

	// expected result
	expectMemes := make([]memeDetail, 0)
	expectMemes = append(
		expectMemes,
		memeDetail{
			ID:       0,
			Title:    "chuan",
			ImageURL: "www.blabla.com/blabla.jpg",
			About:    "ark chuan",
		})
	expectedBytes, _ := json.Marshal(expectMemes)
	expectedString := string(expectedBytes)

	testAPIUtils(
		"GET",
		"/get_trending_memes",
		handler,
		t,
		expectedString,
		nil)
}

func TestInsertMemeWithoutTags(t *testing.T) {
	// TODO @jeff: reset DB and insert example data
	config, err := GetDBConfig(testConfigPath)
	if err != nil {
		t.Fatal(err)
	}
	handler := ServeWrapper(config, InsertMemeWithoutTags)

	// expected result
	expectMemes := make([]memeDetail, 0)
	expectMemes = append(
		expectMemes,
		memeDetail{
			ID:       0,
			Title:    "chuan",
			ImageURL: "www.blabla.com/blabla.jpg",
			About:    "ark chuan",
		})
	expectedBytes, _ := json.Marshal(expectMemes)
	expectedString := string(expectedBytes)

	// compose url body
	newMemes := make([]memeDetail, 0)
	newMemes = append(
		newMemes,
		memeDetail{
			ID:       0,
			Title:    "chuan",
			ImageURL: "www.blabla.com/blabla.jpg",
			About:    "ark chuan",
		})
	jsonStr, _ := json.Marshal(newMemes)

	testAPIUtils(
		"GET",
		"/insert_meme_without_tags",
		handler,
		t,
		expectedString,
		bytes.NewBuffer(jsonStr))
}

func TestInsertMemeAboutsAndTags(t *testing.T) {
	// TODO @jeff: reset DB and insert example data
	config, err := GetDBConfig(testConfigPath)
	if err != nil {
		t.Fatal(err)
	}
	handler := ServeWrapper(config, InsertMemeAboutsAndTags)

	// expected result
	expectMemes := make([]memeDetail, 0)
	expectMemes = append(
		expectMemes,
		memeDetail{
			ID:       0,
			Title:    "chuan",
			ImageURL: "www.blabla.com/blabla.jpg",
			About:    "ark chuan",
		})
	expectedBytes, _ := json.Marshal(expectMemes)
	expectedString := string(expectedBytes)

	// compose url body
	newMemes := make([]memeDetail, 0)
	newMemes = append(
		newMemes,
		memeDetail{
			ID:       0,
			Title:    "chuan",
			ImageURL: "www.blabla.com/blabla.jpg",
			About:    "ark chuan",
		})
	jsonStr, _ := json.Marshal(newMemes)

	testAPIUtils(
		"GET",
		"/insert_meme_abouts_and_tags",
		handler,
		t,
		expectedString,
		bytes.NewBuffer(jsonStr))
}

func TestIncrementMemeClick(t *testing.T) {
	// TODO @jeff: reset DB and insert example data
	config, err := GetDBConfig(testConfigPath)
	if err != nil {
		t.Fatal(err)
	}
	handler := ServeWrapper(config, InsertMemeWithoutTags)

	// expected result
	expectMemes := make([]memeDetail, 0)
	expectMemes = append(
		expectMemes,
		memeDetail{
			ID:       0,
			Title:    "chuan",
			ImageURL: "www.blabla.com/blabla.jpg",
			About:    "ark chuan",
		})
	expectedBytes, _ := json.Marshal(expectMemes)
	expectedString := string(expectedBytes)

	// compose url body
	IDs := make([]int, 0)
	IDs = append(IDs, 0)
	memeIDs := memeIDInput{IDs: IDs}
	jsonStr, _ := json.Marshal(memeIDs)

	testAPIUtils(
		"GET",
		"/increment_meme_click",
		handler,
		t,
		expectedString,
		bytes.NewBuffer(jsonStr),
	)
}
