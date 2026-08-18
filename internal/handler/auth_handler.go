package handler

import (
	"net/http"

	"github.com/aditya-singh-finbox/todo-api/internal/model"
	"github.com/aditya-singh-finbox/todo-api/internal/service"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userService *service.UserService
	authService *service.AuthService
}

func NewAuthHandler(
	userService *service.UserService,
	authService *service.AuthService,
) *AuthHandler {

	return &AuthHandler{
		userService: userService,
		authService: authService,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.userService.Register(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})

}

func (h *AuthHandler) Login(c *gin.Context) {

	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	accessToken, refreshToken, user, err :=
		h.authService.Login(
			req.Email,
			req.Password,
		)

	if err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",

		"access_token": accessToken,

		"refresh_token": refreshToken,

		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}

func (h *AuthHandler) Refresh(c *gin.Context) {

	var req struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	accessToken, refreshToken, err :=
		h.authService.Refresh(req.RefreshToken)

	if err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func (h *AuthHandler) Logout(c *gin.Context) {

	var req struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	err := h.authService.Logout(req.RefreshToken)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not logout",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "logged out successfully",
	})
}
