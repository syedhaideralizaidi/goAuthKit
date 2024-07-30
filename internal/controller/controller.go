package controller

import db "github.com/syedhaideralizaidi/authkit/internal/database/sqlc"

// Controller holds the dependencies for the user-related endpoints
type Controller struct {
	Queries *db.Queries
}

// NewController creates a new instance of Controller with dependencies
func NewController(queries *db.Queries) *Controller {
	return &Controller{
		Queries: queries,
	}
}
