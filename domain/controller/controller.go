package controller

import (
	"faceit/domain/service"
	"faceit/infrastructure/database"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type IUsersController interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Remove(c *gin.Context)
	Get(c *gin.Context)
}

type UsersController struct {
	service service.IUserService
	store   database.IDatabase
}

func NewUserController(service service.IUserService) *UsersController {
	return &UsersController{service: service}
}

func (u *UsersController) Run(port string) *http.Server {
	// init gin
	gin.SetMode(gin.DebugMode)
	router := gin.New()

	router.GET("/health", u.HealthCheck)

	user := router.Group("/user")
	{
		user.POST("/create", u.Create)
		user.POST("/update", u.Update)
		user.DELETE("/:id", u.Remove)
		user.POST("/get", u.Get)
	}

	// gin middleware config
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "HEAD", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "X-Requested-With", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"status": false, "data": gin.H{"status": false, "message": fmt.Sprintf("Page not found: %s, method: %s", c.Request.URL, c.Request.Method)}})
	})

	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"status": false, "data": gin.H{"status": false, "message": "Method not found"}})
	})

	// Note: we use http server to have graceful shutdown
	server := &http.Server{
		Addr:         port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 20 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	go func() {
		log.Printf("Listening and serving HTTP on %s\n", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("gin sever stoped with err: %s \n", err)
		}
	}()

	return server
}

func (u *UsersController) Create(c *gin.Context) {

}

func (u *UsersController) Update(c *gin.Context) {

}

func (u *UsersController) Remove(c *gin.Context) {

}

func (u *UsersController) Get(c *gin.Context) {

}

// HealthCheck checks database ping
func (u *UsersController) HealthCheck(c *gin.Context) {
	health := map[string]interface{}{
		"store": "up",
	}

	if err := u.store.Ping(); err != nil {
		health["database"] = "down"
		u.ginResponse(c, http.StatusInternalServerError, health)
		return
	}

	u.ginResponse(c, http.StatusOK, health)
}

func (u *UsersController) ginResponse(c *gin.Context, status int, payload interface{}) {
	type Response struct {
		Status  int         `json:"status"`
		Payload interface{} `json:"payload"`
	}

	response := Response{
		Status:  status,
		Payload: payload,
	}

	c.Header("Content-Type", "application/json")
	c.Status(status)

	c.JSON(status, response)
}
