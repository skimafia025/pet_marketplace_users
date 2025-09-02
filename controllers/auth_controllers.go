package controllers

import (
	"net/http"
	"pet_marketplace_users/config"
	"pet_marketplace_users/logging"
	"pet_marketplace_users/models"
	"pet_marketplace_users/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Register(c *gin.Context) {
	var req models.RegisterRequest
	log := logging.Log(c)

	log.WithFields(logrus.Fields{
		"name":  req.Name,
		"email": req.Email,
	}).Info("Создаем пользователя")

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var existingUser models.User
	if err := config.DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "A user with this email address already exists",
		})
		return
	}

	user := models.User{
		EMAIL:    req.Email,
		PASSWORD: req.Password,
		Name:     req.Name,
	}

	if err := user.Hashpassword(); err != nil {
		log.Error("Error hashing password:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error create user",
		})
	}
	if err := config.DB.Create(&user).Error; err != nil {
		log.Error("Error create user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error create user",
		})
		return
	}

	token, err := utils.GenerateToken(user.ID, user.EMAIL)
	if err != nil {
		log.Error("Token generation error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Token generation error",
		})
		return
	}

	log.Info("Registered user:", user.EMAIL)

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data": gin.H{
			"user":  user,
			"token": token,
		},
	})
}

func Login(c *gin.Context) {
	log := logging.Log(c)

	var req models.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var user models.User

	if err := config.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	if !user.CheckPassword(req.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	token, err := utils.GenerateToken(user.ID, user.EMAIL)
	if err != nil {
		log.Error("Token generation error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Token generation error",
		})
		return
	}
	log.Info("Authorized user:", req.Email)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"user":  user,
			"token": token,
		},
	})
}
