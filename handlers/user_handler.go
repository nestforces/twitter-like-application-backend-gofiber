package handlers

import (
    "context"
    "twitter-like-backend/database"
    "github.com/gofiber/fiber/v2"
)

func ChangeUsername(c *fiber.Ctx) error {
    userID := c.Locals("user_id").(int)
    newUsername := c.FormValue("username")

    if newUsername == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "New username cannot be empty"})
    }

    sql := `UPDATE users SET username = $1 WHERE user_id = $2`
    _, err := database.DB.Exec(context.Background(), sql, newUsername, userID)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update username"})
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Username updated successfully"})
}

package handlers

import (
    "context"
    "twitter-like-backend/database"
    "github.com/gofiber/fiber/v2"
)

func UpdateProfile(c *fiber.Ctx) error {
    userID := c.Locals("user_id").(int)
    bio := c.FormValue("bio")

    // Handle profile picture upload
    file, err := c.FormFile("profile_picture")
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing or invalid profile picture"})
    }

    fileName := file.Filename
    err = c.SaveFile(file, "./public/images/profile/"+fileName)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save profile picture"})
    }

    sql := `UPDATE users SET bio = $1, profile_picture = $2 WHERE user_id = $3`
    _, err = database.DB.Exec(context.Background(), sql, bio, fileName, userID)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update profile"})
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Profile updated successfully"})
}

