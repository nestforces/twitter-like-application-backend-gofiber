package handlers

import (
    "context"
    "twitter-backend/database"
    "twitter-backend/models"
    "github.com/gofiber/fiber/v2"
)

func SendDirectMessage(c *fiber.Ctx) error {
    senderID := c.Locals("user_id").(int)
    receiverID := c.FormValue("receiver_id")
    content := c.FormValue("content")

    if content == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Message content cannot be empty"})
    }

    sql := `INSERT INTO direct_messages (sender_id, receiver_id, content) VALUES ($1, $2, $3)`
    _, err := database.DB.Exec(context.Background(), sql, senderID, receiverID, content)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to send message"})
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Message sent successfully"})
}

func GetDirectMessages(c *fiber.Ctx) error {
    userID := c.Locals("user_id").(int)
    otherUserID := c.Params("id")

    sql := `SELECT * FROM direct_messages WHERE (sender_id = $1 AND receiver_id = $2) OR (sender_id = $2 AND receiver_id = $1) ORDER BY created_at`
    rows, err := database.DB.Query(context.Background(), sql, userID, otherUserID)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get messages"})
    }
    defer rows.Close()

    messages := []models.DirectMessage{}
    for rows.Next() {
        var dm models.DirectMessage
        if err := rows.Scan(&dm.ID, &dm.SenderID, &dm.ReceiverID, &dm.Content, &dm.CreatedAt); err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to scan messages"})
        }
        messages = append(messages, dm)
    }

    return c.Status(fiber.StatusOK).JSON(messages)
}
