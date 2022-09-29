package server

import (
	"exam/entity"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	_ "exam/docs"
)

type ServerHandler struct {
	service Service
}

func NewServiceStruct(service Service) ServerHandler {
	return ServerHandler{service: service}
}

// TakeWord
// @Summary      Take Word
// @Description
// @Tags         exam
// @Accept       json
// @Produce      json
// @Param        request body InsertWordRequest true "Create"
// @Success      200 
// @Failure      400
// @Failure      500
// @Router       /new-word [Post]
func (s ServerHandler) TakeWord(c *gin.Context) {
	var words InsertWordRequest
	if err := c.BindJSON(&words); err != nil {
		c.JSON(400, gin.H{"error0:": err})
		return
	}
	var w entity.Words
	var wl entity.WordList
	for k, v := range words.Words {
		k = strings.ToLower(k)
		w = entity.NewWord(k, v)
		wl = append(wl, w)
	}
	if err := s.service.InsertWords(c.Request.Context(), wl); err != nil {
		c.JSON(500, gin.H{"error1": err})
		return
	}
	c.JSON(200, gin.H{"success": true})
}

// GiveAllWord
// @Summary      Give All Word
// @Description
// @Tags         exam
// @Accept       json
// @Produce      json
// @Success      200 {object} []entity.Words
// @Failure      400
// @Failure      500
// @Router       /take-all-word [Get]
func (s ServerHandler) GiveAllWord(c *gin.Context) {
	words, err := s.service.GetAllWords(c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{
			"error11": err,
		})
		return
	}
	c.JSON(200, words)
}

// GiveAllWord
// @Summary      Give All Word
// @Description
// @Tags         exam
// @Accept       json
// @Produce      json
// @Param        page query int32 true "GET"
// @Param        limit query int32 true "GET"
// @Success      200 {object} []entity.Words
// @Failure      400
// @Failure      500
// @Router       /take-limited-word [Get]
func (s ServerHandler) GiveLimitedWord(c *gin.Context) {
	page := c.Query("page")
	limit := c.Query("limit")
	page2, err := strconv.Atoi(page)
	limit2, err := strconv.Atoi(limit)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
	}
	wordslist, err := s.service.GetLimitedWords(c.Request.Context(), int32(page2), int32(limit2))
	if err != nil {
		c.JSON(500, gin.H{
			"Error": err,
		})
	}
	c.JSON(200, wordslist)
}
