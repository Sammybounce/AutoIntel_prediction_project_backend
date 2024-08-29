package model

import (
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	Router *fiber.App
}

func (s *Server) Middleware(callback func(r *fiber.App)) {
	callback(s.Router)
}

func (s *Server) Routes(callback func(r *fiber.App)) {
	callback(s.Router)
}
