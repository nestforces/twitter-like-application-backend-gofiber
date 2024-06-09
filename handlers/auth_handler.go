package handlers

import (
    "context"
    "time"
    "twitter-like-backend/database"
    "twitter-like-backend/models"
    "twitter-like-backend/utils"

    "github.com/gofiber/fiber/v2"
    "golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
    user := new(models.User)
    if err := c.BodyParser(user); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    user.PasswordHash = string(hashedPassword)
    user.CreatedAt = time.Now()

    sql := `INSERT INTO users (username, email, password_hash, name, bio, location, website, created_at)
            VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING user_id`
    err = database.DB.QueryRow(context.Background(), sql, user.Username, user.Email, user.PasswordHash, user.Name, user.Bio, user.Location, user.Website, user.CreatedAt).Scan(&user.UserID)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    token, err := utils.GenerateJWT(user.UserID)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return c.Status(fiber.StatusCreated).JSON(fiber.Map{"token": token, "user": user})
}

func Login(c *fiber.Ctx) error {
    input := new(struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    })
    if err := c.BodyParser(input); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

    var user models.User
    sql := `SELECT user_id, password_hash FROM users WHERE email = $1`
    err := database.DB.QueryRow(context.Background(), sql, input.Email).Scan(&user.UserID, &user.PasswordHash)
    if err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
    }

    token, err := utils.GenerateJWT(user.UserID)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": token})
}
