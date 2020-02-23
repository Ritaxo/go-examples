package main

import "github.com/gofiber/fiber"

func main(){
  app := fiber.New()
  app.Get("/", func (c *fiber.Ctx)  {
    c.Send("Hellow, World!")

  })

  app.Listen(3000)
}
