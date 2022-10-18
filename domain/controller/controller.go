package controller

import (
	"faceit/domain/dto"
	"faceit/domain/service"
	"faceit/infrastructure/database"
	"fmt"
	"log"
	"net/http"
	"strconv"
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

// NewUserController - Creates a new user controller with dependency injection
func NewUserController(service service.IUserService) *UsersController {
	return &UsersController{service: service}
}

// Run - Starts the gin engine and sets up the http routes
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

// Create - Handler to create a user with the given user information
func (u *UsersController) Create(c *gin.Context) {
	var request createRequest
	if err := c.BindJSON(&request); err != nil {
		u.ginResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userDTO := &dto.User{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		NickName:  request.NickName,
		Email:     request.Email,
		Country:   request.Country,
	}

	createdUserDTO, err := u.service.Create(c.Request.Context(), userDTO, request.Password)
	if err != nil {
		u.ginResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	u.ginResponse(c, http.StatusOK, createdUserDTO)
}

// Update - Handler to update the given user
func (u *UsersController) Update(c *gin.Context) {
	var request updateRequest
	if err := c.BindJSON(&request); err != nil {
		u.ginResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userDTO := &dto.User{
		ID:        request.ID,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		NickName:  request.NickName,
		Email:     request.Email,
		Country:   request.Country,
	}

	if err := u.service.Update(c.Request.Context(), userDTO, request.Password); err != nil {
		u.ginResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	u.ginResponse(c, http.StatusOK, nil)
}

// Remove - Handler to remove a user based on the provided user ID
func (u *UsersController) Remove(c *gin.Context) {
	ID := c.Param("id")
	IDint64, err := strconv.ParseInt(ID, 10, 64)
	if err != nil {
		u.ginResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := u.service.Remove(c.Request.Context(), IDint64); err != nil {
		u.ginResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	u.ginResponse(c, http.StatusOK, nil)
}

// Get - Handler for getting users based on the provided criteria in the URL parameters
func (u *UsersController) Get(c *gin.Context) {
	var filter *dto.Filter
	country, found := c.GetQuery("country")
	if found {
		filter.Country = country
	}

	createdAt, found := c.GetQuery("created_at")
	if found {
		timeLayout := "2006-01-02 15:04:05"
		parsedCreatedAt, err := time.Parse(timeLayout, createdAt)
		if err != nil {
			u.ginResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		filter.CreatedAt = parsedCreatedAt
	}

	updatedAt, found := c.GetQuery("created_at")
	if found {
		timeLayout := "2006-01-02 15:04:05"
		parsedUpdatedAt, err := time.Parse(timeLayout, updatedAt)
		if err != nil {
			u.ginResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		filter.UpdatedAt = parsedUpdatedAt
	}

	var request getRequest
	if err := c.BindJSON(&request); err != nil {
		u.ginResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userDTOs, count, err := u.service.Get(c.Request.Context(), filter, request.Page, request.PageSize)
	if err != nil {
		u.ginResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response := struct {
		users []*dto.User
		count int64
	}{
		users: userDTOs,
		count: count,
	}

	u.ginResponse(c, http.StatusOK, response)
}

// HealthCheck - Checks database health
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

// ginResponse - A simple helper function to prepare the response structure
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
