package handlers

import (
    "context"
    "time"
    "twitter-like-backend/database"
    "twitter-like-backend/models"

    "github.com/gofiber/fiber/v2"
)

func CreateTweet(c *fiber.Ctx) error {
    tweet := new(models.Tweet)
    if err := c.BodyParser(tweet); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

    tweet.CreatedAt = time.Now()
    sql := `INSERT INTO tweets (user_id, content, created_at) VALUES ($1, $2, $3) RETURNING tweet_id`
    err := database.DB.QueryRow(context.Background(), sql, tweet.UserID, tweet.Content, tweet.CreatedAt).Scan(&tweet.TweetID)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return c.Status(fiber.StatusCreated).JSON(tweet)
}

func GetTweets(c *fiber.Ctx) error {
    rows, err := database.DB.Query(context.Background(), "SELECT tweet_id, user_id, content, created_at FROM tweets ORDER BY created_at DESC")
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    defer rows.Close()

    var tweets []models.Tweet
    for rows.Next() {
        var tweet models.Tweet
        if err := rows.Scan(&tweet.TweetID, &tweet.UserID, &tweet.Content, &tweet.CreatedAt); err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
        }
        tweets = append(tweets, tweet)
    }

    return c.Status(fiber.StatusOK).JSON(tweets)
}
