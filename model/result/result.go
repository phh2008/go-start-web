package result

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const success = "200"
const fail = "500"

type Result struct {
	Code      string      `json:"code"`
	Message   string      `json:"message"`
	TimeStamp int64       `json:"timeStamp"`
	Data      interface{} `json:"data"`
}

func Success() Result {
	return Result{Code: success, Message: "SUCCESS", TimeStamp: time.Now().UnixMilli(), Data: nil}
}

func Error() Result {
	return Result{Code: fail, Message: "ERROR", TimeStamp: time.Now().UnixMilli(), Data: nil}
}

func NewResult(code string, msg string, data interface{}) Result {
	return Result{Code: code, Message: msg, Data: data}
}

func writeJSON(data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, data)
}

func Response(code string, msg string, data interface{}, c *gin.Context) {
	writeJSON(NewResult(code, msg, data), c)
}

func Ok(c *gin.Context) {
	writeJSON(Success(), c)
}

func OkMsg(msg string, c *gin.Context) {
	writeJSON(Success().SetMsg(msg), c)
}

func OkData(data interface{}, c *gin.Context) {
	writeJSON(Success().SetData(data), c)
}

func OkDetail(data interface{}, msg string, c *gin.Context) {
	writeJSON(Success().SetData(data).SetMsg(msg), c)
}

func Fail(c *gin.Context) {
	writeJSON(Error(), c)
}

func FailMsg(msg string, c *gin.Context) {
	writeJSON(Error().SetMsg(msg), c)
}

func FailDetail(data interface{}, msg string, c *gin.Context) {
	writeJSON(Error().SetData(data).SetMsg(msg), c)
}

func (a Result) IsSuccess() bool {
	return a.Code == success
}

func (a Result) SetCode(code string) Result {
	a.Code = code
	return a
}

func (a Result) SetMsg(msg string) Result {
	a.Message = msg
	return a
}

func (a Result) SetData(data interface{}) Result {
	a.Data = data
	return a
}

func (a Result) Response(c *gin.Context) {
	c.JSON(http.StatusOK, a)
}
