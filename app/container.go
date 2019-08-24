package app

type memeIcon struct {
	ImageURL string `json:"image_url"`
	Title    string `json:"title"`
	ItemURL  string `json:"item_url"`
}

type memeIDInput struct {
	IDs         []int `json:"meme_ids,omitempty"`
	NumOfResult int   `json:"n_result,omitempty"`
}

type trendingInput struct {
	NumOfResult int `json:"n_result"`
}

type memeDetail struct {
	ID       int      `json:"id,omitempty"`
	Title    string   `json:"title"`
	ImageURL string   `json:"image_url"`
	About    string   `json:"about,omitempty"`
	Tags     []string `json:"tags,omitempty"`
}

type memeDetailsInput struct {
	MemeDetails []memeDetail `json:"meme_details"`
}
