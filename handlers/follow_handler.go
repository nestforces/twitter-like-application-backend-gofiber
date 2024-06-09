package handlers

import (
    "context"
    "time"
    "twitter-like-backend/database"
    "twitter-like-backend/models"

    "github.com/gofiber/fiber/v2"
)

func FollowUser(c *fiber.Ctx) error {
    followerID := c.Locals("user_id").(int)
    followeeID, err := c.ParamsInt("id")
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
    }

    follow := models.Follow{
        FollowerID: followerID,
        FolloweeID: followeeID,
        FollowedAt: time.Now(),
    }

    sql := `INSERT INTO follows (follower_id, followee_id, followed_at) VALUES ($1, $2, $3)`
    _, err = database.DB.Exec(context.Background(), sql, follow.FollowerID, follow.FolloweeID, follow.FollowedAt)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return c.Status(fiber.StatusCreated).JSON(follow)
}

func UnfollowUser(c *fiber.Ctx) error {
    followerID := c.Locals("user_id").(int)
    followeeID, err := c.ParamsInt("id")
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
    }

    sql := `DELETE FROM follows WHERE follower_id = $1 AND followee_id = $2`
    _, err = database.DB.Exec(context.Background(), sql, followerID, followeeID)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Unfollowed successfully"})
}
