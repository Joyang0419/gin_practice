package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"gin_practice/service/user"
)

type Auth struct {
	user user.IService
}

func NewAuth(user user.IService) *Auth {
	return &Auth{
		user: user,
	}
}

func (h *Auth) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request struct {
			Username string `json:"username" binding:"required"`
			Password string `json:"password" binding:"required"`
		}
		if err := c.ShouldBind(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":      "failed",
				"description": "Invalid form data: " + err.Error(),
			})
			return
		}

		if _, err := h.user.Login(c, request.Username, request.Password); err != nil {
			switch {
			case errors.Is(err, user.ErrInvalidInput):
				c.JSON(
					http.StatusBadRequest,
					gin.H{
						"status":      "failed",
						"description": "invalid input",
					},
				)
				return
			case errors.Is(err, user.ErrLoginFailed):
				c.JSON(
					http.StatusOK,
					gin.H{
						"status":      "failed",
						"description": "wrong password",
					},
				)
				return
			default:
				c.JSON(
					http.StatusInternalServerError,
					gin.H{
						"status":      "failed",
						"description": "Internal server error",
					},
				)
				return
			}
		}

		c.JSON(
			http.StatusOK,
			gin.H{
				"status":      "success",
				"description": "ok",
			},
		)
	}
}

func (h *Auth) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request struct {
			Username string `json:"username" binding:"required"`
			Password string `json:"password" binding:"required"`
		}
		if err := c.ShouldBind(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":      "failed",
				"description": "Invalid form data: " + err.Error(),
			})
			return
		}

		if _, err := h.user.Register(c, request.Username, request.Password); err != nil {
			switch {
			case errors.Is(err, user.ErrInvalidInput):
				c.JSON(
					http.StatusBadRequest,
					gin.H{
						"status":      "failed",
						"description": "invalid input",
					},
				)
				return
			case errors.Is(err, user.ErrDuplicateUsername):
				c.JSON(
					http.StatusOK,
					gin.H{
						"status":      "failed",
						"description": "username is already exist",
					},
				)
				return
			default:
				c.JSON(
					http.StatusInternalServerError,
					gin.H{
						"status":      "failed",
						"description": "Internal server error",
					},
				)
				return
			}
		}

		c.JSON(
			http.StatusOK,
			gin.H{
				"status":      "success",
				"description": "ok",
			},
		)
	}
}
