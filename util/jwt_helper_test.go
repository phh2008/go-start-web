package util

import (
	"com.phh/start-web/pkg/config"
	"encoding/json"
	"fmt"
	"github.com/cristalhq/jwt/v4"
	"testing"
	"time"
)

// 生成 token
func TestCreateToken(t *testing.T) {
	config := config.NewConfig("../config")
	jwtHelper := NewJwtHelper(config)
	claims := UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        "1000",
			Subject:   "tom",
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Minute * 1)},
		},
		Phone: "18975391618",
	}
	token, err := jwtHelper.CreateToken(claims)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Token     %s\n", token.String())
	fmt.Printf("Algorithm %s\n", token.Header().Algorithm)
	fmt.Printf("Type      %s\n", token.Header().Type)
	fmt.Printf("Claims    %s\n", token.Claims())
	fmt.Printf("HeaderPart    %s\n", token.HeaderPart())
	fmt.Printf("ClaimsPart    %s\n", token.ClaimsPart())
	fmt.Printf("PayloadPart   %s\n", token.PayloadPart())
	fmt.Printf("SignaturePart %s\n", token.SignaturePart())
}

// 校验token
func TestVerifyToken(t *testing.T) {
	config := config.NewConfig("../config")
	var tk = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiIxMDAwIiwic3ViIjoidG9tIiwiZXhwIjoxNjUxOTA4MjcyLCJQaG9uZSI6IjE4OTc1MzkxNjE4IiwiUnVsZSI6IiJ9.B9jIbQ0C74KcKOx5I0w0tyOfc2Yc0D5mQdbbccS-M9c"
	jwtHelper := NewJwtHelper(config)
	token, err := jwtHelper.VerifyToken(tk)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Token     %s\n", token.String())
	fmt.Printf("Algorithm %s\n", token.Header().Algorithm)
	fmt.Printf("Type      %s\n", token.Header().Type)
	fmt.Printf("Claims    %s\n", token.Claims())
	fmt.Printf("HeaderPart    %s\n", token.HeaderPart())
	fmt.Printf("ClaimsPart    %s\n", token.ClaimsPart())
	fmt.Printf("PayloadPart   %s\n", token.PayloadPart())
	fmt.Printf("SignaturePart %s\n", token.SignaturePart())
	var user UserClaims
	json.Unmarshal(token.Claims(), &user)
	fmt.Println(user)
	if !user.IsValidExpiresAt(time.Now()) {
		fmt.Println("expire .......")
	}
}
