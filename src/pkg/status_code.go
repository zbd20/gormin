package pkg

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	HiErr string = "400-10000"
)

type RespResult struct {
	Code   int64       `json:"code"`
	Result interface{} `json:"result"`
}

func NewResult(result interface{}) RespResult {
	return RespResult{
		Code:   0,
		Result: result,
	}
}

func WriteResponse(c *gin.Context, code string, result interface{}) {
	httpCode, resp := NewResponseResult(code, result)
	c.JSON(httpCode, resp)
}

func NewResponseResult(code string, result interface{}) (int, RespResult) {
	codes := strings.Split(code, "-")
	httpCode, _ := strconv.Atoi(codes[0])
	statusCode, _ := strconv.ParseInt(codes[1], 10, 64)

	return httpCode, RespResult{
		Code:   statusCode,
		Result: fmt.Sprint(result),
	}
}
