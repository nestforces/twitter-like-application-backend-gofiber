package handlers

import (
    "context"
    "time"
    "twitter-like-backend/database"
    "twitter-like-backend/models"

    "github.com/gofiber/fiber/v2"
)

func Retweet(c *fiber.Ctx) error {
    userID := c.Locals("user_id").(int)
    tweetID, err := c.ParamsInt("id")
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid tweet ID"})
    }

    retweet := models.Retweet{
        UserID:     userID,
        TweetID:    tweetID,
        RetweetedAt: time.Now(),
    }

    sql := `INSERT INTO retweets (user_id, tweet_id, retweeted_at) VALUES ($1, $2, $3)`
    _, err = database.DB.Exec(context.Background(), sql, retweet.UserID, retweet.TweetID, retweet.RetweetedAt)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return c.Status(fiber.StatusCreated).JSON(retweet)
}

func RemoveRetweet(c *fiber.Ctx) error {
    userID := c.Locals("user_id").(int)
    tweetID, err := c.ParamsInt("id")
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid tweet ID"})
    }

    sql := `DELETE FROM retweets WHERE user_id = $1 AND tweet_id = $2`
    _, err = database.DB.Exec(context.Background(), sql, userID, tweetID)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Retweet removed successfully"})
}
