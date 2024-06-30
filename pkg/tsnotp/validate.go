package tsnotp

import (
	"fmt"
	"time"
)

// تابع اعتبارسنجی TSNOTP
func validate(otp string, secret string, sequenceNumber uint64) (bool, error) {
	currentTimestamp := time.Now().Unix() / 30 // تبدیل زمان فعلی به timestamp 30 ثانیه‌ای

	publicIP, networkCount, macID, err := getNetworkInfo()
	if err != nil {
		return false, err
	}
	networkInfo := fmt.Sprintf("%s%d%s", publicIP, networkCount, macID)

	// بررسی کد برای زمان فعلی
	generatedOTP := generate(secret, currentTimestamp, sequenceNumber, networkInfo)
	if generatedOTP == otp {
		return true, nil
	}

	// بررسی کد برای 30 ثانیه قبلی و 30 ثانیه بعدی (تطابق با بازه 30 ثانیه‌ای)
	for i := int64(-1); i <= 1; i++ {
		timestamp := currentTimestamp + i
		generatedOTP := generate(secret, timestamp, sequenceNumber, networkInfo)
		if generatedOTP == otp {
			return true, nil
		}
	}

	return false, nil
}
