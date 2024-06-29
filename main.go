package main

import (
	"github.com/ghafbu/ghafotp/pkg/totp"
	"github.com/ghafbu/ghafotp/pkg/tsotp"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

func main() {
	//fiber add
	app := fiber.New()
	app.Use(
		recover.New(),
	)
	var KeyDB = make(map[string]string, 0)
	totp.Router(app, KeyDB)
	tsotp.Router(app)
	//log.Fatal(app.ListenTLS(":443", certFile, keyFile))
	app.Listen(":3000")
}
