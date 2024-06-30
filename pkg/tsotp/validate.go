package tsotp

import (
	"time"
)

// تابع اعتبارسنجی TOTP
func validate(otp string, secret string, sequenceNumber uint64) bool {
	currentTimestamp := time.Now().Unix() / 30 // تبدیل زمان فعلی به timestamp 30 ثانیه‌ای
	//tolerance := int64(0) // بازه زمانی برای اعتبارسنجی، 0 به معنی فقط زمان فعلی

	// بررسی کد برای زمان فعلی
	timestamp := currentTimestamp
	generatedOTP := generate(secret, timestamp, sequenceNumber) // تولید OTP جدید
	if generatedOTP == otp {                                    // بررسی تطابق کد تولید شده با کد ورودی
		return true
	}

	// بررسی کد برای 30 ثانیه قبلی و 30 ثانیه بعدی (تطابق با بازه 30 ثانیه‌ای)
	for i := int64(-1); i <= 1; i++ {
		timestamp := currentTimestamp + i
		generatedOTP := generate(secret, timestamp, sequenceNumber) // تولید OTP جدید
		if generatedOTP == otp {                                    // بررسی تطابق کد تولید شده با کد ورودی
			return true
		}
	}

	return false
}
