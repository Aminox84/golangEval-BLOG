package services

import (
	"blog/internal/models"
	"blog/internal/repositories"
)

type ArticleService struct {
	articleRepository *repositories.ArticleRepository
}

func NewArticleService(articleRepository *repositories.ArticleRepository) *ArticleService {
	return &ArticleService{articleRepository}
}

func (s *ArticleService) GetArticles() ([]models.Article, error) {
	return s.articleRepository.GetArticles()
}

func (s *ArticleService) CreateArticle(article *models.Article) (int64, error) {
	return s.articleRepository.CreateArticle(article)
}

func (s *ArticleService) LikeArticle(articleID int) error {
	return s.articleRepository.LikeArticle(articleID)
}

func (s *ArticleService) DislikeArticle(articleID int) error {
	return s.articleRepository.DislikeArticle(articleID)
}

func (s *ArticleService) DeleteArticle(articleID int) error {
	return s.articleRepository.DeleteArticle(articleID)
}
