package handlers

import (
    "context"
    "time"
    "twitter-like-backend/database"
    "twitter-like-backend/models"

    "github.com/gofiber/fiber/v2"
)

func LikeTweet(c *fiber.Ctx) error {
    userID := c.Locals("user_id").(int)
    tweetID, err := c.ParamsInt("id")
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid tweet ID"})
    }

    like := models.Like{
        UserID:  userID,
        TweetID: tweetID,
        LikedAt: time.Now(),
    }

    sql := `INSERT INTO likes (user_id, tweet_id, liked_at) VALUES ($1, $2, $3)`
    _, err = database.DB.Exec(context.Background(), sql, like.UserID, like.TweetID, like.LikedAt)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return c.Status(fiber.StatusCreated).JSON(like)
}

func UnlikeTweet(c *fiber.Ctx) error {
    userID := c.Locals("user_id").(int)
    tweetID, err := c.ParamsInt("id")
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid tweet ID"})
    }

    sql := `DELETE FROM likes WHERE user_id = $1 AND tweet_id = $2`
    _, err = database.DB.Exec(context.Background(), sql, userID, tweetID)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Unliked successfully"})
}
