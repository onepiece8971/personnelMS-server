package Api

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"personnelMS-server/app/Models"
)

// 获取多个文章标签
func Index(c *gin.Context) {
	dbUser := Models.Users{Id: 1}
	user := dbUser.GetUser()
	c.JSON(http.StatusOK, user)
}
