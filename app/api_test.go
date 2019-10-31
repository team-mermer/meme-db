package app

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

const testConfigPath = "../test_config.yaml"

func testAPIUtils(method string, APIUrl string, handler http.Handler, t *testing.T) string {
	req, err := http.NewRequest(method, APIUrl, nil)
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

	return rr.Body.String()
}

func TestGetMemeDetails(t *testing.T) {
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

	APIBodyString := testAPIUtils("GET", "/get_meme_details", handler, t)

	if APIBodyString != expectedString {
		t.Errorf(
			"handler returned unexpected body: got %v want %v",
			APIBodyString,
			expectedString,
		)
	}
}

func TestGetMemeWithoutTags(t *testing.T) {
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

	APIBodyString := testAPIUtils("GET", "/get_meme_without_tags", handler, t)

	if APIBodyString != expectedString {
		t.Errorf(
			"handler returned unexpected body: got %v want %v",
			APIBodyString,
			expectedString,
		)
	}
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

	APIBodyString := testAPIUtils("GET", "/get_trending_memes", handler, t)

	if APIBodyString != expectedString {
		t.Errorf(
			"handler returned unexpected body: got %v want %v",
			APIBodyString,
			expectedString,
		)
	}
}

func TestInsertMemeWithoutTags(t *testing.T) {
	// TODO @jeff: reset DB and insert example data
}

func TestInsertMemeAboutsAndTags(t *testing.T) {
	// TODO @jeff: reset DB and insert example data
}

func TestIncrementMemeClick(t *testing.T) {
	// TODO @jeff: reset DB and insert example data
}
