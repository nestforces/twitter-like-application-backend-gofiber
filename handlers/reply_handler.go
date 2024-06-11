package handlers

import (
    "context"
    "twitter-backend/database"
    "github.com/gofiber/fiber/v2"
)

func ReplyToTweet(c *fiber.Ctx) error {
    userID := c.Locals("user_id").(int)
    tweetID := c.Params("id")
    replyContent := c.FormValue("content")

    sql := `INSERT INTO replies (user_id, tweet_id, content) VALUES ($1, $2, $3)`
    _, err := database.DB.Exec(context.Background(), sql, userID, tweetID, replyContent)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to reply to tweet"})
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Reply posted successfully"})
}
