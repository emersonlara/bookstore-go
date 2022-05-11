package users

import (
	"net/http"
	"strconv"
	"user-api/domain/users"
	"user-api/services"
	"user-api/utils/errors"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User

	// Pode ser substituido por o codigo abaixo com ShouldBindJSON
	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	//TODO: Handle error
	// 	return
	// }
	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	//TODO: handle json error
	// 	fmt.Println(err.Error())
	// 	return
	// }

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

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user)

}
