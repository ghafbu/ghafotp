package tsnotp

import (
	"fmt"
	"github.com/ghafbu/ghafotp/utl"
	"github.com/gofiber/fiber/v3"
	"time"
)

var KeyDB = make(map[string]string)
var sequenceNumber = uint64(1)

func Router(paramApp *fiber.App) {
	//network print
	//ip
	var ip, _ = getLocalIP()
	fmt.Println("ip: ", ip)
	//macid
	var macId, _ = getMACID()
	fmt.Println("macId: ", macId)
	//ssid
	var networkCounter, _ = getNetworkCount()
	fmt.Printf("networkCounter:%d ", networkCounter)

	var app = paramApp.Group("/tsnotp")

	//get
	app.Post("/get", func(c fiber.Ctx) error {
		//mobile = c.Params("mobile")
		var req = &RequestGetStruct{}
		var errs = c.Bind().JSON(req)
		if errs != nil {
			return c.SendString("error request bind....")
		}

		if req.Mobile == "" {
			return c.JSON(map[string]any{
				"error":  "mobile no available ....",
				"mobile": req.Mobile,
			})
		}

		//fmt.Println("mobile param:", req.Mobile)
		//fmt.Println("reflect type:", reflect.TypeOf(req.Mobile))

		//generation key
		secretKey, err := utl.GenerationSecretKey("tsnotp", req.Mobile)
		if err != nil {
			fmt.Println("generation key error:", err)
		}

		fmt.Println("secretKey create:", secretKey.Secret())
		//save key
		KeyDB[req.Mobile] = secretKey.Secret()

		// check KeyDB
		//for k, v := range KeyDB {
		//	fmt.Printf("loop => KeyDB[%s] = %s\n", k, v)
		//}

		//generation code
		publicIP, networkCount, macID, err := getNetworkInfo()
		if err != nil {
			fmt.Println("Error getting network info:", err)
		}
		networkInfo := fmt.Sprintf("%s%d%s", publicIP, networkCount, macID)

		//now := time.Now().UTC()
		timestamp := time.Now().Unix() / 30
		code := generate(secretKey.Secret(), timestamp, sequenceNumber, networkInfo)

		if err != nil {
			return c.JSON(map[string]any{
				"error":   "Error generating code ....",
				"details": err.Error(),
			})
		}

		//return
		return c.JSON(map[string]any{
			//"keydb":        KeyDB,
			"code":      code,
			"secretKey": secretKey.Secret(),
		})
	})

	//verify
	app.Post("/verify", func(c fiber.Ctx) error {
		fmt.Println(time.Now(), time.Now().UTC())
		var req = &RequestVerifyStruct{}
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
		//fmt.Println("reflect type verify:", reflect.TypeOf(req.Mobile))

		secretKey, ok := KeyDB[req.Mobile]
		//secretKey := fmt.Sprintf("%s", valueMobile)
		fmt.Println("secretKey get:", secretKey)
		if !ok {
			fmt.Println("No secret key found for mobile:", req.Mobile)
			fmt.Println("keyDB::", KeyDB)
			return c.JSON(map[string]any{
				"error":  "No secret key found for this mobile number.",
				"mobile": req.Mobile,
				"KeyDB":  KeyDB,
			})
		}

		//validation
		valid, err := validate(req.Otp, secretKey, sequenceNumber)
		if err != nil {
			return c.JSON(map[string]any{
				"error":   "validation error.",
				"details": err.Error(),
			})
		}

		return c.JSON(map[string]any{
			"valid": valid,
			"KeyDB": KeyDB,
		})
	})
}
