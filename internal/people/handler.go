package people

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type handler struct {
	repository *repository
}

func (h *handler) create(c echo.Context) error {
	var p People
	if err := c.Bind(&p); err != nil {
		return c.NoContent(422)
	}

	if err := p.Validate(); err != nil {
		return c.NoContent(422)
	}

	layout := "2006-01-02"
	birthDate, err := time.Parse(layout, p.BirthDate)
	if err != nil {
		return c.NoContent(422)
	}

	entity := &PeopleEntity{
		ID:        uuid.NewString(),
		Name:      p.Name,
		Nickname:  p.Nickname,
		BirthDate: birthDate,
		Stack:     p.Stack,
		Search:    p.Name + " " + p.Nickname + " " + strings.Join(p.Stack, " "),
	}

	if err := h.repository.create(entity); err != nil {
		return err
	}

	c.Response().Header().Set("Location", "/pessoas/"+entity.ID)
	return c.NoContent(201)
}

func (h *handler) readByID(c echo.Context) error {
	id := c.Param("id")

	entity, err := h.repository.readByID(id)
	if err != nil {
		return err
	}

	p := People{
		ID:        entity.ID,
		Name:      entity.Name,
		Nickname:  entity.Nickname,
		BirthDate: entity.BirthDate.Format("2006-01-02"),
		Stack:     entity.Stack,
	}

	return c.JSON(200, p)
}

func (h *handler) searchByTerm(c echo.Context) error {
	term := c.QueryParam("t")
	if term == "" {
		return c.NoContent(400)
	}

	entities, err := h.repository.searchByTerm(term)
	if err != nil {
		return err
	}

	peoples := make([]People, 0, len(entities))
	for _, entity := range entities {
		p := People{
			ID:        entity.ID,
			Name:      entity.Name,
			Nickname:  entity.Nickname,
			BirthDate: entity.BirthDate.Format("2006-01-02"),
			Stack:     entity.Stack,
		}

		peoples = append(peoples, p)
	}

	return c.JSON(200, peoples)
}

func (h *handler) count(c echo.Context) error {
	count, err := h.repository.count()
	if err != nil {
		return err
	}

	return c.JSON(200, map[string]int64{"count": count})
}
