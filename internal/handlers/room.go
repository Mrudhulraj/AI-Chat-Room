package handlers

import (
	"fmt"

	"github.com/gofiber/contrib/v3/websocket"
	"github.com/gofiber/fiber/v3"
	gguid "github.com/google/uuid"
)

func RoomCreate(c fiber.Ctx) error {
	return c.Redirect().To(fmt.Sprintf("/room/%s", gguid.New().String()))
}

func Room(c fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		c.Status(400)
		return nil
	}
	uuid, suuid, _ := createOrGetRoom(uuid)

}

func RoomWebsocket(c *websocket.Conn) {
	uuid := c.Params("uuid")
	if uuid == "" {
		return
	}

	_, _, room := createOrGetRoom(uuid)
}

func createOrGetRoom(uuid string) (string, string, Room) {

}
