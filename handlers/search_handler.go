package handlers

import (
    "context"
    "twitter-backend/database"
    "github.com/gofiber/fiber/v2"
)

func SearchTweets(c *fiber.Ctx) error {
    keyword := c.Query("keyword")

    sql := `SELECT * FROM tweets WHERE content LIKE '%' || $1 || '%'`
    rows, err := database.DB.Query(context.Background(), sql, keyword)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to search tweets"})
    }
    defer rows.Close()

    tweets := []fiber.Map{}
    for rows.Next() {
        var id, userID, content string
        var createdAt time.Time
        if err := rows.Scan(&id, &userID, &content, &createdAt); err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to scan tweets"})
        }
        tweets = append(tweets, fiber.Map{
            "id":         id,
            "user_id":    userID,
            "content":    content,
            "created_at": createdAt,
        })
    }

    return c.Status(fiber.StatusOK).JSON(tweets)
}
