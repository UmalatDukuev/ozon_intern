package models

type Post struct {
	ID              int    `db:"id" json:"id"`
	Title           string `db:"title" json:"title"`
	Content         string `db:"content" json:"content"`
	CommentsEnabled bool   `db:"comments_enabled" json:"comments_enabled"`
}

type NewPost struct {
	Title           string
	Content         string
	CommentsEnabled bool
}
