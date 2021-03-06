package users

import (
	"fmt"
	"net/http" //require to send specific http status
	"strconv"

	"github.com/arya0618/my-bookstore-api/domain/users"
	"github.com/arya0618/my-bookstore-api/services"
	"github.com/arya0618/my-bookstore-api/utils/errors"
	"github.com/gin-gonic/gin"
)

var (
	counter int
)

// GetUser is
func GetUser(c *gin.Context) {
	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64) //base:10 basesize=64
	if userErr != nil {
		err := errors.NewBadRequestError("User id should be number")
		c.JSON(err.Status, err)
		return
	}
	user, getErr := services.GetUser(userID)
	fmt.Println("--->", user)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user)

}

// CreatetUser is
func CreatetUser(c *gin.Context) {
	var user users.User
	fmt.Println(user)
	if err := c.ShouldBindJSON(&user); err != nil {
		// restErr := errors.RestErr{
		// 	Message: "invalid json body",
		// 	Status:  http.StatusBadRequest,
		// 	Error:   "bad_request",
		// }
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
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
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

// // SearchtUser is
// func SearchtUser(c *gin.Context) {
// 	c.String(http.StatusNotImplemented , "implement me!")
// }
