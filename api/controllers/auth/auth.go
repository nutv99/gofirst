package auth

import (
	orm "melivecode/jwt-api/orm"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	_ "github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var hmacSampleSecret []byte

type RegisterBody struct {
	UserName string `json:"username"  binding:"required"`
	Password string `json:"password" binding:"required"`
	FullName string `json:"FullName" binding:"required"`
	Avatar   string `json:"Avatar" binding:"required"`
}

type LoginBody struct {
	UserName string `json:"username"  binding:"required"`
	Password string `json:"password" binding:"required"`
	FullName string `json:"FullName" binding:"required"`
	Avatar   string `json:"Avatar" binding:"required"`
}

func Register(c *gin.Context) {

	var json RegisterBody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userExists orm.User
	// Get first matched record
	orm.Db.Where("user_name = ?", json.UserName).First(&userExists)
	if userExists.ID > 0 {
		c.JSON(http.StatusOK,
			gin.H{
				"status":  "Error",
				"message": "User Exists",
				"userID":  "No",
			})
		return

	}

	encryptPassword, _ := bcrypt.GenerateFromPassword([]byte(json.Password), 10)
	user := orm.User{UserName: json.UserName, Password: string(encryptPassword),
		FullName: json.FullName, Avatar: json.Avatar,
	}

	orm.Db.Create(&user) // pass pointer of data to Create
	if user.ID > 0 {
		c.JSON(http.StatusOK,
			gin.H{
				"status":          "ok",
				"message":         "User registered Success",
				"userID":          user.ID,
				"encryptPassword": encryptPassword,
			})
	} else {
		c.JSON(http.StatusOK,
			gin.H{
				"status":          "Error",
				"message":         "User registered Failed",
				"userID":          "No",
				"encryptPassword": encryptPassword,
			})

	}

}

func Login(c *gin.Context) {

	var json LoginBody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userExists orm.User
	// Get first matched record
	orm.Db.Where("user_name = ?", json.UserName).First(&userExists)
	if userExists.ID == 0 {
		c.JSON(http.StatusOK,
			gin.H{
				"status":  "Error",
				"message": "ไม่พบ User Exists",
				"userID":  "No",
			})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(userExists.Password), []byte(json.Password))

	// ExpireMinute := os.Getenv("EXPIRE_MINUTE")
	ExpireTime := time.Now().Add(time.Minute * 30).Unix()
	if err == nil {
		hmacSampleSecret = []byte(os.Getenv("JWT_SECRET_KEY"))
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"userId": userExists.ID,
			"exp":    time.Now().Add(time.Minute * 30).Unix(),
		})
		tokenString, err := token.SignedString(hmacSampleSecret)
		//fmt.Println(tokenString,err)

		c.JSON(http.StatusOK,
			gin.H{
				"status":  "Error",
				"message": "Login Success",
				"userID":  "No",
				"token":   tokenString,
				"Expires": ExpireTime,
				"err":     err,
			})
	} else {
		c.JSON(http.StatusOK,
			gin.H{
				"status":  "Error",
				"message": "Login Fail (Password)",
				"userID":  "No",
			})
	}

}
