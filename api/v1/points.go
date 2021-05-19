package v1

import (
	"net/http"

	"acgfate/services/points"
	"github.com/gin-gonic/gin"
)

// PointsSign 普通签到
// @Summary 普通签到
// @Description 普通签到接口
// @Tags Points
// @Produce  application/json
// @Param Authorization header string true "用户令牌"
// @Param form body points.SignService true "用户信息"
// @Success 0 {object} serializer.UserPointsResponse "msg: "Success"
// @Failure 60001 {object} serializer.Response "msg: 参数错误"
// @Router /sign [post]
func PointsSign(c *gin.Context) {
	var form points.SignService
	res := form.DoSign(c)
	c.JSON(http.StatusOK, res)
}
