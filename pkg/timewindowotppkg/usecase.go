package timewindowotppkg

import (
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"time"
)

const Secret = "JBSWY3DPEHPK3PX3P"

// generateOTP تولید کد OTP با استفاده از کلید مخفی و Time Window
func GenerateOTP(secret string) (string, error) {
	otp, err := totp.GenerateCode(secret, time.Now())
	if err != nil {
		return "", err
	}
	return otp, nil
}

// validateOTP بررسی کد OTP ارائه شده با کد تولید شده در محدوده زمانی 5 دقیقه
func ValidateOTP(secret, code string) (bool, error) {
	valid := totp.Validate(code, secret)
	if valid {
		return true, nil
	}

	// بررسی پنجره زمانی 5 دقیقه بعد از زمان کنونی
	for i := 1; i <= 5; i++ {
		futureTime := time.Now().Add(time.Duration(i) * time.Minute)
		valid, _ = totp.ValidateCustom(code, secret, futureTime, totp.ValidateOpts{
			Period:    30,
			Skew:      1,
			Digits:    otp.DigitsSix,
			Algorithm: otp.AlgorithmSHA1,
		})
		if valid {
			return true, nil
		}
	}

	return false, nil
}
