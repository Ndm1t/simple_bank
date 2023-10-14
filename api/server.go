package api

import (
	db "bankingApp/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests
type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	router.DELETE("/accounts/:id", server.deleteAccount)
	router.PUT("/accounts", server.updateAccount)

	server.router = router
	return server
}

func (server *Server) Start(adress string) error {
	return server.router.Run(adress)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
