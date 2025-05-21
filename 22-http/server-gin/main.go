// Расширенный пример роут-карты на Gin.
// go get github.com/gin-gonic/gin
//
// Запуск: go run main.go
//
// Демо-маршруты:
//  GET  /ping                  -> {"message":"pong"}
//  GET  /hello/:name           -> {"hello":"<name>"}
//  GET  /search?q=term         -> {"q":"term"}
//  POST /login (JSON)          -> {"user":"<user>","status":"logged in"}
//  GET  /api/v1/users/42       -> {"id":"42"}
//  GET  /static/*filepath      -> {"path":"/anything/you/like"}

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Структура для примера JSON-запроса/ответа
type Login struct {
	User string `json:"user" binding:"required"`
	Pass string `json:"pass" binding:"required"`
}

func main() {
	r := gin.New()        // «чистый» движок без логгера/рестора
	r.Use(gin.Logger())   // добавим логгер
	r.Use(gin.Recovery()) // автo-recover после panic

	// 1. Простейший маршрут
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// 2. Параметр пути :name  =>  /hello/alex
	r.GET("/hello/:name", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"hello": c.Param("name")})
	})

	// 3. Query-параметр ?q=term  =>  /search?q=go
	r.GET("/search", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"q": c.Query("q")})
	})

	// 4. POST с JSON-телом. Gin валидирует и мапит в структуру
	r.POST("/login", func(c *gin.Context) {
		var req Login
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"user": req.User, "status": "logged in"})
	})

	// 5. Группировка префикса /api/v1
	api := r.Group("/api/v1")
	{
		api.GET("/users/:id", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"id": c.Param("id")})
		})
	}

	// 6. «Подстановочный» маршрут — всё после /static
	r.GET("/static/*filepath", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"path": c.Param("filepath")})
	})

	r.Run(":8989") // старт сервера
}

/*
Ключевые отличия от net/http:

  • Маршруты описываются декларативно методами r.GET/POST и т.д.
  • Параметры пути (:id) и wildcard (*filepath) разбирает Gin, нет
    необходимости писать парсер.
  • Группы (r.Group) упрощают версионирование API.
  • Helpers вроде c.JSON(), c.Query(), c.Param() убирают шаблонный код.
  • Middleware подключается одной строкой r.Use().
*/

// curl -uri http://localhost:8989/ping -method get
// curl -uri http://localhost:8989/hello/gopher -method get
// curl -uri http://localhost:8989/search?q=123 -method get
