package users

import (
	"fmt"
	"net/http" //require to send specific http status

	"github.com/arya0618/my-bookstore-api/domain/users"
	"github.com/arya0618/my-bookstore-api/services"
	"github.com/gin-gonic/gin"
)

var (
	counter int
)

// GetUser is
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

// CreatetUser is
func CreatetUser(c *gin.Context) {
	var user users.User
	fmt.Println(user)
	if err := c.ShouldBindJSON(&user); err != nil {
		return
	}
	/*bytes, err := ioutil.ReadAll(c.Request.Body)
	fmt.Println(err)
	if err != nil {
		// error handling part
		return
	}
	if err := json.Unmarshal(bytes, &user); err != nil {
		fmt.Println(err.Error())
		return
	}*/
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		return
	}

	c.JSON(http.StatusCreated, result)
}

// // SearchtUser is
// func SearchtUser(c *gin.Context) {
// 	c.String(http.StatusNotImplemented , "implement me!")
// }
