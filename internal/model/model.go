package model

// NewsFullDetailed, NewsShortDetailed, Comment,

type NewsFullDetailed struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Comments    *[]Comment
}

type NewsShortDetailed struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Comment struct {
	ID       int64  `json:"id"`
	Text     string `json:"text"`
	Children []Comment
}
