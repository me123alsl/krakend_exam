package application

import (
	"fmt"
	"jwk_sample/api"
	"jwk_sample/authorize"
	"net/http"
	"os"

	"github.com/labstack/echo"
	md "github.com/labstack/echo/middleware"
)

// 서버 구조체
type Server struct {
	Port string
}

// 서버 라우터 설정 함수
func (s *Server) SetRouter(e *echo.Echo) {
	e.GET("/api/v1/authorize", api.Authorize)
	e.GET("/api/v1/authorize/verify", api.Verify)
	e.GET(
		"/api/v1/authorize/verifycookie",
		api.VerifyCookie(),
		md.JWTWithConfig(md.JWTConfig{
			ErrorHandler: func(error) error {
				return echo.NewHTTPError(http.StatusUnauthorized, api.Result{Success: false, Message: "Unauthorized", Data: "Unauthorized"})
			},
			SigningKey:  []byte(authorize.TOKEN_SECRET),
			Claims:      md.DefaultJWTConfig.Claims,
			TokenLookup: "cookie:access-token",
			BeforeFunc: func(ctx echo.Context) {
				fmt.Println(ctx.Request().Header)
				cookie := ctx.Request().Header.Get("Cookie")
				fmt.Println(cookie)
			},
		},
		))
	e.GET("/api/v1/authorize/refresh", api.Refresh)
}

// 서버 실행 함수
func (s *Server) RunServer() {

	f, err := os.OpenFile("logfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("error opening file: %v", err))
	}
	defer f.Close()

	// echo framework 초기화
	e := echo.New()

	// 미들웨어 설정
	e.Use(md.Recover())
	e.Use(md.LoggerWithConfig(md.LoggerConfig{
		Format: "${time_rfc3339} ${status} ${method} ${uri} ${latency_human} ${bytes_in} ${bytes_out} ${error}\n",
		Output: f,
	}))
	// 라우터 설정
	s.SetRouter(e)
	// 서버 실행
	if err := e.Start(":" + s.Port); err != nil {
		e.Logger.Fatal(err)
	}

}
