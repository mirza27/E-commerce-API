package middleware

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"context"
    "fmt"
    "time"
)

var (
	SessionUser int64 // dummy untuk session
)

func Authentication(key string) bool{
	op2 := rdb.Get(context.Background(), key)
    if err := op2.Err(); err != nil {
        return false 
    }
    _, err := op2.Result()
    if err != nil {
        return false
    }
	fmt.Println(SessionUser)
	return true
}


// autth login
func Login(c *gin.Context, id string, username string) {
	op1 := rdb.Set(context.Background(), id, username, time.Duration(3) * time.Minute)
	err := op1.Err()
    if  err != nil {
		panic(err)
    }
	c.JSON(http.StatusOK, gin.H{
		"message": "User Sign In successfully",
		})
}


