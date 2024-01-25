package repositories

import (
	"blog/internal/models"
	"database/sql"
)

type ArticleRepository struct {
	db *sql.DB
}

func NewArticleRepository(db *sql.DB) *ArticleRepository {
	return &ArticleRepository{db}
}

func (r *ArticleRepository) GetArticles() ([]models.Article, error) {
	var articles []models.Article

	rows, err := r.db.Query("SELECT id, content FROM articles")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var article models.Article
		if err := rows.Scan(&article.ID, &article.Content); err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}

	return articles, nil
}

func (r *ArticleRepository) CreateArticle(article *models.Article) (int64, error) {
	result, err := r.db.Exec(`
        INSERT INTO articles (content) 
        VALUES (?)`, article.Content)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *ArticleRepository) LikeArticle(articleID int) error {
	_, err := r.db.Exec("UPDATE articles SET likes = likes + 1 WHERE id = ?", articleID)
	return err
}

func (r *ArticleRepository) DislikeArticle(articleID int) error {
	_, err := r.db.Exec("UPDATE articles SET dislikes = dislikes + 1 WHERE id = ?", articleID)
	return err
}

func (r *ArticleRepository) DeleteArticle(articleID int) error {
	_, err := r.db.Exec("DELETE FROM articles WHERE id = ?", articleID)
	return err
}
