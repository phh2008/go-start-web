package middleware

import (
	"com.phh/start-web/pkg/global"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func IsPreFlightRequest(c *gin.Context) bool {
	request := c.Request
	return request.Method == "OPTIONS" &&
		request.Header.Get("Origin") != "" &&
		request.Header.Get("Access-Control-Request-Method") != ""
}

// Cors 放行所有跨域请求
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-Token,X-User-Id")
		c.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS,DELETE,PUT")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type, New-Token, New-Expires-At")
		c.Header("Access-Control-Allow-Credentials", "true")
		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

// CorsByRules 跨域请求
func CorsByRules() gin.HandlerFunc {
	cors := global.Profile.Cors
	allowedOriginPatterns := cors.AllowedOriginPatterns
	if allowedOriginPatterns == nil || len(allowedOriginPatterns) == 0 {
		return Cors()
	}
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		// is cors request
		if origin == "" {
			c.Next()
			return
		}
		var isAllow bool
		for _, v := range allowedOriginPatterns {
			if "*" == v || origin == v {
				isAllow = true
				break
			}
		}
		if isAllow {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Headers", cors.AllowedHeaders)
			c.Header("Access-Control-Allow-Methods", cors.AllowedMethods)
			c.Header("Access-Control-Expose-Headers", cors.ExposeHeaders)
			c.Header("Access-Control-Allow-Credentials", strconv.FormatBool(cors.AllowCredentials))
			c.Header("Access-Control-Max-Age", strconv.FormatInt(cors.MaxAge, 10))
			if c.Request.Method == "OPTIONS" {
				c.AbortWithStatus(http.StatusNoContent)
			}
		} else {
			c.AbortWithStatus(http.StatusForbidden)
		}
		c.Next()
	}
}
