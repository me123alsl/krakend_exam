package api

import (
	"fmt"
	"jwk_sample/authorize"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

type Result struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Authorize(c echo.Context) error {
	if jwtToken, err := authorize.CreateToken(authorize.TOKEN_SECRET); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			Result{Success: false, Message: "Internal Server Error", Data: err.Error()},
		)
		return err
	} else {

		cookies := http.Cookie{
			Name:     "access-token",
			Value:    jwtToken,
			HttpOnly: true,
			Expires:  time.Now().Add(time.Minute * 5),
		}
		c.SetCookie(&cookies)

		c.JSON(
			http.StatusOK,
			Result{Success: true, Message: "Success", Data: jwtToken},
		)
		return nil
	}
}

func Verify(c echo.Context) error {
	// header 전체 출력
	fmt.Println(c.Request().Header)

	if claimMap, err := authorize.VerifyTokenWithClaims(c.Request().Header.Get("Authorization"), authorize.TOKEN_SECRET); err != nil {
		c.JSON(
			http.StatusUnauthorized,
			Result{Success: false, Message: "Unauthorized", Data: err.Error()},
		)
		return err
	} else {
		c.JSON(
			http.StatusOK,
			Result{Success: true, Message: "Success", Data: claimMap},
		)
		return nil
	}
}

func VerifyCookie() echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("Cookie값 Authorized 되었습니다.")

		cookie, err := c.Cookie("access-token")
		if err != nil || cookie.Value == "" {
			return echo.ErrUnauthorized
		}
		if claimMap, err := authorize.VerifyTokenWithClaims(cookie.Value, authorize.TOKEN_SECRET); err != nil {
			c.JSON(
				http.StatusUnauthorized,
				Result{Success: false, Message: "Unauthorized", Data: err.Error()},
			)
		} else {
			c.JSON(
				http.StatusOK,
				Result{Success: true, Message: "Success", Data: claimMap},
			)
			return nil
		}
		return nil
	}
}

func Refresh(c echo.Context) error {
	fmt.Println(c.Request().Header)
	if jwtToken, err := authorize.RefreshToken(c.Request().Header.Get("Authorization"), authorize.TOKEN_SECRET); err != nil {
		c.JSON(
			http.StatusUnauthorized,
			Result{Success: false, Message: "Internal Server Error", Data: err.Error()},
		)
		return err
	} else {
		c.JSON(
			http.StatusOK,
			Result{Success: true, Message: "Success", Data: jwtToken},
		)
		return nil
	}
}
