package totp

import (
	"fmt"
	"github.com/pquerna/otp/totp"
)

// تابع برای اعتبارسنجی OTP
func ValidateOTP(userOTP string, secretKey string) (bool, error) {
	fmt.Println("secretKey valid:", secretKey)
	valid := totp.Validate(userOTP, secretKey)
	return valid, nil
}
