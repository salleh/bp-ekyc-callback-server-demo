package handler

import (
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func IDOnly(c *fiber.Ctx) error {
	log.Debugln(c.Request().String())

	return c.JSON(fiber.Map{"status": "success", "message": "ID Only Callback", "data": nil})
}

func IDFace(c *fiber.Ctx) error {
	log.Debugln(c.Request().String())

	return c.JSON(fiber.Map{"status": "success", "message": "ID Face Callback", "data": nil})
}

func IDLive(c *fiber.Ctx) error {
	log.Debugln(c.Request().String())

	return c.JSON(fiber.Map{"status": "success", "message": "ID Live Callback", "data": nil})
}

func IDReflect(c *fiber.Ctx) error {
	log.Debugln(c.Request().String())

	return c.JSON(fiber.Map{"status": "success", "message": "ID Reflect Callback", "data": nil})
}
