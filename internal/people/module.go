package people

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func RegisterModule(g *echo.Group, db *sqlx.DB) {
	h := &handler{
		repository: &repository{
			db: db,
		},
	}
	g.POST("/pessoas", h.create)
	g.GET("/pessoas/:id", h.readByID)
	g.GET("/pessoas", h.searchByTerm)
	g.GET("/contagem-pessoas", h.count)
}
