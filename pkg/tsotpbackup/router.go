package tsotpbackup

import (
	"fmt"
	"github.com/ghafbu/ghafotp/utl"
	"github.com/gofiber/fiber/v3"
	"time"
)

type RequestStruct struct {
	Otp    string `json:"otp"`
	Mobile string `json:"mobile"`
}

var KeyDB = map[string]string{}

func Router(paramApp *fiber.App) {
	var app = paramApp.Group("/tsotp")
	//fetch from db real time
	sequence := 1

	//get
	app.Get("/get/:mobile", func(c fiber.Ctx) error {
		mobile := c.Params("mobile")
		timestamp := time.Now().Unix()
		//secretKey := utl.GenerationSecretKey(16)

		//generation key
		secretKey, err := utl.GenerationSecretKey(mobile, "tsotp")
		if err != nil {
			fmt.Println("generation key error:", err)
		}

		//generation otp
		otp, err, aes, nonce := GenerateOTP(timestamp, sequence, secretKey.Secret())
		if err != nil {
			fmt.Println("generation otp error:", err)
		}

		KeyDB[mobile] = secretKey.Secret()

		return c.JSON(map[string]any{
			"timestamp":    timestamp,
			"sequence":     sequence,
			"nonce":        nonce,
			"aes":          aes,
			"secretKey":    secretKey.Secret(),
			"otp":          otp,
			"secretKeyURL": secretKey.URL(),
		})
	})

	//verify
	app.Post("/verify", func(c fiber.Ctx) error {
		timestamp := time.Now().Unix()

		var req = &RequestStruct{}
		var errs = c.Bind().JSON(req)
		if errs != nil {
			return c.SendString("error request bind....")
		}

		if req.Mobile == "" || req.Otp == "" {
			return c.JSON(map[string]any{
				"error":  "otp || mobile no available ....",
				"mobile": req.Mobile,
				"otp":    req.Otp,
			})
		}

		//fetch key
		secretKey := KeyDB[req.Mobile]

		//validation
		valid, err := ValidateOTP(req.Otp, timestamp, sequence, secretKey)
		if err != nil {
			return c.JSON(map[string]any{
				"error":   "error validation otp....",
				"details": err.Error(),
			})
		}

		return c.JSON(map[string]any{
			"valid": valid,
		})
	})
}
