package tsnotp

import (
	"fmt"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"time"
)

// تابع برای اعتبارسنجی OTP
func ValidateOTP(userOTP string, secretKey string) (bool, error) {
	fmt.Println("secretKey valid:", secretKey)
	rv, _ := totp.ValidateCustom(
		userOTP,
		secretKey,
		time.Now().UTC(),
		totp.ValidateOpts{
			Period:    30,
			Skew:      0,
			Digits:    otp.DigitsSix,
			Algorithm: otp.AlgorithmSHA1,
		},
	)

	//valid := totp.Validate(userOTP, secretKey)
	return rv, nil
}
