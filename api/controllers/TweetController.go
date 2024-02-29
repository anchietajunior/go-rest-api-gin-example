package controllers

import (
	"net/http"

	entities "github.com/anchietajunior/go-rest-api/api/entities"
	"github.com/gin-gonic/gin"
)

type TweetController struct {
	tweets []entities.Tweet
}

func NewTweetController() *TweetController {
	return &TweetController{}
}

// Apply a functions to the TweetController struct
func (t *TweetController) FindAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, t.tweets)
}

func (t *TweetController) Create(ctx *gin.Context) {
	tweet := entities.NewTweet()

	if err := ctx.BindJSON(&tweet); err != nil {
		return
	}

	t.tweets = append(t.tweets, *tweet)

	ctx.JSON(http.StatusOK, tweet)
}

func (t *TweetController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	for idx, tweet := range t.tweets {
		if tweet.ID == id {
			t.tweets = append(t.tweets[0:idx], t.tweets[idx+1:]...)
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{
		"error": "Tweet not found",
	})
}
