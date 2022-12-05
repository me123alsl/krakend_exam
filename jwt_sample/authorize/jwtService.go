package authorize

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const TOKEN_SECRET = "secret"

func CreateToken(key string) (string, error) {
	// 토큰 생성
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":  "golang",
		"admin": true,
		"exp":   time.Now().Add(time.Minute * 5).Unix(),
	})

	// 토큰 서명
	token.Header["kid"] = TOKEN_SECRET
	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string, key string) (bool, error) {
	// 토큰 검증
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})

	if err != nil {
		return false, err
	}

	if token.Valid {
		return true, nil
	} else {
		return false, nil
	}
}

// JWT 토큰 검증 함수 (토큰 검증 후 토큰 정보 추출)
func VerifyTokenWithClaims(tokenString string, key string) (jwt.MapClaims, error) {
	// 토큰 검증
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, nil
	}
}

// 요청 cookie값으로 access-token 토큰값으로 cliam 추출
func GetClaimsFromCookie(cookie string) (jwt.MapClaims, error) {
	// 토큰 검증
	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(TOKEN_SECRET), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, nil
	}
}

// function refresh token
func RefreshToken(tokenString string, key string) (string, error) {
	// 토큰 검증
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})

	if err != nil {
		return "", err
	}

	// 토큰 재생성
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		claims["exp"] = time.Now().Add(time.Minute * 5).Unix()
		newToken := jwt.NewWithClaims(jwt.SigningMethodPS256, claims)
		newTokenString, err := newToken.SignedString([]byte(key))
		if err != nil {
			return "", err
		}
		return newTokenString, nil
	} else {
		return "", nil
	}
}
