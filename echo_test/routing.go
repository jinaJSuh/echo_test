package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//URL(Uniform Resource Locators) : http (protocol)://localhost(host):8080(port)/path/server/(resource path)?a=b&x=y(query)
	//Protocol :서버와 클라이언트가 정보를 주고 받는 규칙  (HTTP SMTP FTP DHCP telnet DNS  등등)
	//IDL(Interface Definition Language)  : 정보를 저장하는 규칙 (XML , JSON , proto) HTTP + XML 조합 ,  HTTP + RESTful API + JSON 조합

	//	Get:		존재하는 자원에 대한 요청
	//	PUT:		존재하는 자원에 대한 변경
	//	DELETE:		존재하는 자원에 대한 삭제
	//	POST:		새로운 자원을 생성

	// Routes
	e.GET("/", hello)
	e.GET("/users", getUser)
	e.PUT("/users", putUser)
	e.POST("/users", postUser)
	e.DELETE("/users", deleteUser)

	//e.POST example
	e.POST("/save", postSave)

	// Request  : client 가 보내는 data 를 server 가 받는 법
	//1. URL Path로 보내는 법
	e.GET("/users/:name", getUserPath)
	//2. FORM으로 보내는 법 :put http://127.0.0.1:1323/users?name=joe
	e.PUT("/users", putUserForm)
	//3. Query로 보내는 법
	e.DELETE("/users", deleteUserQuery)

	//Response : Server 가 Client에게 전달하는 Data 생성
	//1. client 가 GET 으로 접속할때  Text 로 보내는 법
	e.GET("/users", getUserText)
	//2. clent 가 PUT 으로 접속할때 HTML 로 보내는 법
	e.PUT("/users", putUserHTML)

	// static  : GET http://127.0.0.1:1323/static/file.txt
	// 실제 directory 는 assets 이고 클라이언트 에서는 static 으로 알고 있다.
	e.Static("/static", "assets")

	e.File("/", "public/index.html")

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func getUser(c echo.Context) error {
	return c.String(http.StatusOK, "getUser")
}
func putUser(c echo.Context) error {
	return c.String(http.StatusOK, "putUser")
}
func postUser(c echo.Context) error {
	return c.String(http.StatusOK, "postUser")
}
func deleteUser(c echo.Context) error {
	return c.String(http.StatusOK, "deleteUser")
}

func getUserPath(c echo.Context) error {
	name := c.Param("name")
	return c.String(http.StatusOK, "getUserPath:"+name)
}

func putUserForm(c echo.Context) error {
	name := c.FormValue("name")
	return c.String(http.StatusOK, "putUserForm:"+name)
}

func deleteUserQuery(c echo.Context) error {
	name := c.QueryParam("name")
	return c.String(http.StatusOK, "deleteUserQuery:"+name)
}

func getUserText(c echo.Context) error {
	return c.String(http.StatusOK, "getUserText : response server ")
}

func putUserHTML(c echo.Context) error {
	return c.HTML(http.StatusOK, "<strong> putUserHTML </strong>")
}

// e.POST("/save", postSave) http://127.0.0.1:1323/save?name=Joe Smith&email=joe@labstack.com
func postSave(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+", email:"+email)
}
