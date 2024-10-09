package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mehedimayall/go-cqrs/internal/entities"
	"github.com/mehedimayall/go-cqrs/internal/handlers/commands"
	repositories "github.com/mehedimayall/go-cqrs/internal/repositories/abstractions"
)

type MovieWriteController struct {
	repo repositories.IWriteRepository[entities.Movie]
}

func NewMovieController(repo repositories.IWriteRepository[entities.Movie]) MovieWriteController {
	return MovieWriteController{
		repo: repo,
	}
}

// Save Movie
func (c *MovieWriteController) Add(ctx *fiber.Ctx) error {
	movie := entities.Movie{}

	if err := ctx.BodyParser(&movie); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	command := commands.NewAddMovieCommand(c.repo)
	movieId, err := command.Handle(&movie)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return Ok(ctx, movieId)
}

// Update Movie
func (c *MovieWriteController) Update(ctx *fiber.Ctx) error {
	movie := entities.Movie{}

	if err := ctx.BodyParser(&movie); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	command := commands.NewUpdateMovieCommand(c.repo)
	err := command.Handle(&movie)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return Ok(ctx, "")
}

// Delete Movie
func (c *MovieWriteController) Delete(ctx *fiber.Ctx) error {
	movieId := ctx.Params("id")

	if movieId == "" {
		return ctx.Status(http.StatusBadRequest).JSON("Please provide a valid movie id")
	}

	command := commands.NewDeleteMovieCommand(c.repo)
	err := command.Handle(movieId)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return Ok(ctx, "")
}
