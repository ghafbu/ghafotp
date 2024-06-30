package tsotpbackup

import (
	"crypto/aes"
	"fmt"
	"github.com/pquerna/otp/totp"
)

// محاسبه OTP بر اساس تایم‌استمپ و شماره توالی
func GenerateOTP(timestamp int64, sequence int, secretKey string) (string, error, string, string) {
	// محاسبه Nonce بر اساس تایم‌استمپ و شماره توالی
	nonce := fmt.Sprintf("%d%d", sequence, timestamp)

	//fmt.Println("nonce:", nonce)

	// رمزنگاری Nonce با استفاده از AES (شبیه‌سازی)
	encryptedNonce, err := encryptAES(nonce, secretKey)
	if err != nil {
		return "", err, "", ""
	}
	// تراز کردن پویا (Dynamic Truncation) برای تولید OTP به طول 8 رقم
	//fmt.Println("nonce => AES:", fmt.Sprintf("%d", encryptedNonce))
	var aes = fmt.Sprintf("%d", encryptedNonce)

	otp := dynamicTruncate(encryptedNonce)
	return otp, nil, aes, nonce
}

// شبیه‌سازی رمزنگاری AES (باید با رمزنگاری واقعی جایگزین شود)
func encryptAES(paramNonce string, secretKey string) ([]byte, error) {
	// تبدیل secretKey به بایت‌ها
	key := []byte(secretKey)

	// انتخاب nonce بر اساس timestamp و sequence number
	nonce := make([]byte, 16)
	copy(nonce[:16], []byte(paramNonce))

	// رمزنگاری nonce با استفاده از AES
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	encryptedNonce := make([]byte, 16)
	block.Encrypt(encryptedNonce, nonce)

	return encryptedNonce, nil
	//return input, nil
}

// تراز کردن پویا برای تولید OTP
func dynamicTruncate(input []byte) string {
	// تراشیدن پویا روی encryptedNonce به طول 8 بایت
	truncated := input[:8]

	// تبدیل به رشته اعداد
	otp := ""
	for _, b := range truncated {
		otp += fmt.Sprintf("%02d", int(b))
	}
	return otp[:8] // تضمین اینکه طول OTP دقیقاً 8 کاراکتر باشد
}

// تابع برای اعتبارسنجی OTP
func ValidateOTP(userOTP string, timestamp int64, sequence int, secretKey string) (bool, error) {
	valid := totp.Validate(userOTP, secretKey)
	return valid, nil

	//expectedOTP, err, _, _ := GenerateOTP(timestamp, sequence, secretKey)
	//if err != nil {
	//	return false, err
	//}
	//return userOTP == expectedOTP, nil
}
