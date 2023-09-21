package people

import "github.com/labstack/echo/v4"

var (
	ErrNoPeopleFound = echo.NewHTTPError(404)
	ErrInvalidPeople = echo.NewHTTPError(422)
	ErrNicknameTaken = echo.NewHTTPError(422)
	ErrInternal      = echo.NewHTTPError(500)
)
