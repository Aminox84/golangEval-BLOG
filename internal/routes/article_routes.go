package routes

import (
	"blog/internal/models"
	"blog/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ArticleRoutes struct {
	articleService *services.ArticleService
}

func NewArticleRoutes(articleService *services.ArticleService) *ArticleRoutes {
	return &ArticleRoutes{articleService}
}

func (ar *ArticleRoutes) RegisterRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/home", ar.getArticles)
		v1.POST("/article", ar.createArticle)
		v1.POST("/articles/:articleID/like", ar.likeArticle)
		v1.POST("/articles/:articleID/dislike", ar.dislikeArticle)
		v1.DELETE("/articles/:articleID", ar.deleteArticle)
	}
}

func (ar *ArticleRoutes) getArticles(c *gin.Context) {
	articles, err := ar.articleService.GetArticles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch articles"})
		return
	}

	c.JSON(http.StatusOK, articles)
}

func (ar *ArticleRoutes) createArticle(c *gin.Context) {
	var article models.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	articleID, err := ar.articleService.CreateArticle(&article)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create article"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"articleID": articleID, "article content": article.Content})
}

func (ar *ArticleRoutes) likeArticle(c *gin.Context) {
	articleID, err := strconv.Atoi(c.Param("articleID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid article ID"})
		return
	}

	if err := ar.articleService.LikeArticle(articleID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like article"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Article liked successfully"})
}

func (ar *ArticleRoutes) dislikeArticle(c *gin.Context) {
	articleID, err := strconv.Atoi(c.Param("articleID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid article ID"})
		return
	}

	if err := ar.articleService.DislikeArticle(articleID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to dislike article"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Article disliked successfully"})
}

func (ar *ArticleRoutes) deleteArticle(c *gin.Context) {
	articleID, err := strconv.Atoi(c.Param("articleID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid article ID"})
		return
	}

	if err := ar.articleService.DeleteArticle(articleID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete article"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Article deleted successfully"})
}
