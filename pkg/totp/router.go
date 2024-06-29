package totp

import (
	"fmt"
	"github.com/ghafbu/ghafotp/utl"
	"github.com/gofiber/fiber/v3"
	"github.com/pquerna/otp/totp"
	"reflect"
	"time"
)

type RequestStruct struct {
	Otp    string `json:"otp"`
	Mobile string `json:"mobile"`
}

func Router(paramApp *fiber.App) {
	var app = paramApp.Group("/totp")
	
	//get
	app.Get("/get/:mobile", func(c fiber.Ctx) error {
		var mobile string = c.Params("mobile")
		fmt.Println("mobile param:", mobile)
		fmt.Println("reflect type:", reflect.TypeOf(mobile))

		//generation key
		secretKey, err := utl.GenerationSecretKey(mobile, "totp")
		if err != nil {
			fmt.Println("generation key error:", err)
		}

		fmt.Println("secretKey create:", secretKey.Secret())
		//save key
		//KeyDB["0911"] = secretKey.Secret()
		KeyDB[mobile] = secretKey.Secret()
		//KeyDB.Store(mobile, secretKey.Secret())

		fmt.Println("KeyDB:", KeyDB)

		//generation code
		now := time.Now().UTC()
		code, err := totp.GenerateCode(secretKey.Secret(), now)
		//code, err := totp.GenerateCodeCustom(secretKey.Secret(), now, totp.ValidateOpts{
		//	Period:    120,
		//	Skew:      1,
		//	Digits:    otp.DigitsSix,
		//	Algorithm: otp.AlgorithmSHA1,
		//})

		if err != nil {
			return c.JSON(map[string]any{
				"error":   "Error generating code ....",
				"details": err.Error(),
			})
		}

		//return
		return c.JSON(map[string]any{
			"code":         code,
			"secretKey":    secretKey.Secret(),
			"secretKeyURL": secretKey.URL(),
		})
	})

	//verify
	app.Post("/verify", func(c fiber.Ctx) error {
		fmt.Println(time.Now(), time.Now().UTC())
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
		//valueMobile, _ := KeyDB.Load(req.Mobile)
		//secretKey := KeyDB["0911"]
		fmt.Println("reflect type verify:", reflect.TypeOf(req.Mobile))

		secretKey, ok := KeyDB[req.Mobile]
		//secretKey := fmt.Sprintf("%s", valueMobile)
		fmt.Println("secretKey get:", secretKey)
		if !ok {
			fmt.Println("No secret key found for mobile:", req.Mobile)
			fmt.Println("keydb::", KeyDB)
			return c.JSON(map[string]any{
				"error":  "No secret key found for this mobile number.",
				"mobile": req.Mobile,
				"KeyDB":  KeyDB,
			})
		}

		//validation
		valid, err := ValidateOTP(req.Otp, secretKey)
		if err != nil {
			return c.JSON(map[string]any{
				"error":   "error validation otp....",
				"details": err.Error(),
			})
		}

		return c.JSON(map[string]any{
			"valid": valid,
			"KeyDB": KeyDB,
		})
	})
}
