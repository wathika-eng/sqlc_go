package main

import (
	"context"
	"fmt"
	"log"
	"sql_c/pkg/repository"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

type CreateUserParams struct {
	Email        string `json:"email" validate:"required,email,min=5"`
	PhoneNumber  string `json:"phone_number" validate:"required,e164,min=10,max=15"`
	PasswordHash string `json:"password_hash" validate:"required,min=8"`
}

func (r *Repo) GetUsers(c *fiber.Ctx) error {
	log.Println(c.IP())
	u, err := r.db.GetAllUsers(context.Background(), repository.GetAllUsersParams{Limit: 10, Offset: 0})
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": true, "message": err.Error()})
	}

	return c.JSON(fiber.Map{"message": u})
}

func (r *Repo) CreateUser(c *fiber.Ctx) error {
	var u repository.CreateUserParams

	// Parse JSON body
	if err := c.BodyParser(&u); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "invalid request body",
			"detail":  err.Error(),
		})
	}

	// Validate struct
	if err := validate.Struct(&u); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   true,
				"message": "validation failed",
				"fields":  errs.Error(), // can format better later
			})
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}
	u.Email = strings.ToLower(strings.TrimSpace(u.Email))

	user, err := r.db.CreateUser(context.Background(), u)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    user,
	})
}

func (r *Repo) FindUser(c *fiber.Ctx) error {
	qEmail := strings.ToLower(strings.TrimSpace(c.Query("email", "")))
	if strings.Trim(qEmail, " ") == "" {
		return c.Redirect("/users")
	}
	user, err := r.db.GetUserByEmail(context.Background(), qEmail)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": fmt.Sprintf("user with email %v not found", qEmail)})
	}
	return c.JSON(user)
}
