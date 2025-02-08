package graph

import (
	"context"
	"ozon_intern/internal/models"

	"github.com/jmoiron/sqlx"
)

type Resolver struct {
	db *sqlx.DB
}

func NewResolver(db *sqlx.DB) *Resolver {
	return &Resolver{db: db}
}

func (r *Resolver) Query_posts(ctx context.Context) ([]*models.Post, error) {
	var posts []*models.Post
	err := r.db.Select(&posts, "SELECT * FROM posts")
	return posts, err
}

func (r *Resolver) Mutation_createPost(ctx context.Context, input models.NewPost) (*models.Post, error) {
	var post models.Post
	query := `INSERT INTO posts (title, content, comments_enabled) VALUES ($1, $2, $3) RETURNING id, title, content, comments_enabled`
	err := r.db.QueryRow(query, input.Title, input.Content, input.CommentsEnabled).Scan(&post.ID, &post.Title, &post.Content, &post.CommentsEnabled)
	return &post, err
}
