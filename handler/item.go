package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"gin_practice/service/item"
	"gin_practice/service/obj"
)

type Item struct {
	item item.IService
}

func NewItem(item item.IService) *Item {
	return &Item{
		item: item,
	}
}

func (h *Item) Items() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request struct {
			Username string `json:"username" binding:"required"`
			Type     string `json:"type" binding:"required"`
		}
		if err := c.ShouldBindQuery(&request); err != nil {
			c.JSON(400, gin.H{
				"status":      "failed",
				"description": "Invalid form data: " + err.Error(),
			})
			return
		}

		items, err := h.item.Items(
			c,
			item.FilterOfItems{
				Usernames: []string{request.Username},
				Type:      request.Type,
			},
		)
		if err != nil {
			switch {
			case errors.Is(err, item.ErrInvalidUsernames):
				c.JSON(
					http.StatusBadRequest,
					gin.H{
						"status":      "failed",
						"description": "invalid usernames",
					},
				)
				return
			default:
				c.JSON(
					http.StatusInternalServerError,
					gin.H{
						"status":      "failed",
						"description": "internal server error",
					},
				)
				return
			}
		}

		c.JSON(
			http.StatusOK,
			gin.H{
				"status": "success",
				"item":   items,
			},
		)
	}
}

func (h *Item) BulkInsert() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request struct {
			Items []*obj.Item `json:"items" binding:"required"`
		}
		if err := c.ShouldBindBodyWithJSON(&request); err != nil {
			c.JSON(400, gin.H{
				"status":      "failed",
				"description": "Invalid form data: " + err.Error(),
			})
			return
		}

		if err := h.item.BulkInsert(c, request.Items); err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"status":      "failed",
					"description": "internal server error",
				},
			)
			return
		}

		c.JSON(
			http.StatusOK,
			gin.H{
				"status": "success",
			},
		)
	}
}
