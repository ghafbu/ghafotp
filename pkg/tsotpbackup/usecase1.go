package tsotpbackup

//
//import (
//	"crypto/aes"
//	"fmt"
//	"time"
//)
//
//// GenerateOTP تابع برای تولید یک OTP بر اساس الگوریتم TSOTP
//func GenerateOTP(secretKey string) (string, error) {
//	// زمان فعلی به عنوان timestamp
//	timestamp := time.Now().Unix()
//
//	// شماره دنباله به طور تصادفی (در اینجا فقط برای نمایش)
//	sequenceNumber := 1234
//
//	// تبدیل secretKey به بایت‌ها
//	key := []byte(secretKey)
//
//	// انتخاب nonce بر اساس timestamp و sequence number
//	nonce := make([]byte, 16)
//	copy(nonce[:8], int64ToBytes(timestamp))
//	copy(nonce[8:], intToBytes(sequenceNumber))
//
//	// رمزنگاری nonce با استفاده از AES
//	block, err := aes.NewCipher(key)
//	if err != nil {
//		return "", err
//	}
//	encryptedNonce := make([]byte, 16)
//	block.Encrypt(encryptedNonce, nonce)
//
//	// تراشیدن پویا (Dynamic Truncation) روی encryptedNonce
//	otp := DynamicTruncation(encryptedNonce)
//
//	return otp, nil
//}
//
//// ValidateOTP بررسی اعتبار OTP بر اساس shared secret و کد OTP ارائه شده
//func ValidateOTP(secretKey string, otp string, expectedOtp string) (bool, error) {
//	// تبدیل secretKey به بایت‌ها
//	key := []byte(secretKey)
//
//	// تبدیل expectedOtp به بایت‌ها
//	expectedOtpBytes := []byte(expectedOtp)
//
//	// رمزگشایی expectedOtp با استفاده از AES
//	block, err := aes.NewCipher(key)
//	if err != nil {
//		return false, err
//	}
//	decrypted := make([]byte, 16)
//	block.Decrypt(decrypted, expectedOtpBytes)
//
//	// تراشیدن پویا (Dynamic Truncation) روی decrypted
//	calculatedOtp := DynamicTruncation(decrypted)
//
//	// مقایسه calculatedOtp با otp ارائه شده
//	return calculatedOtp == otp, nil
//}
//
//// DynamicTruncation تابع برای تراشیدن پویا روی encryptedNonce و تولید رشته 8 کاراکتری از اعداد
//func DynamicTruncation(encryptedNonce []byte) string {
//	// تراشیدن پویا روی encryptedNonce به طول 8 بایت
//	truncated := encryptedNonce[:8]
//
//	// تبدیل به رشته اعداد
//	otp := ""
//	for _, b := range truncated {
//		otp += fmt.Sprintf("%02d", int(b))
//	}
//	return otp[:8] // تضمین اینکه طول OTP دقیقاً 8 کاراکتر باشد
//}
//
//// int64ToBytes تابعی برای تبدیل int64 به بایت‌ها
//func int64ToBytes(i int64) []byte {
//	b := make([]byte, 8)
//	for j := 0; j < 8; j++ {
//		b[j] = byte(i >> uint((7-j)*8))
//	}
//	return b
//}
//
//// intToBytes تابعی برای تبدیل int به بایت
//func intToBytes(i int) []byte {
//	b := make([]byte, 4)
//	b[0] = byte(i >> 24)
//	b[1] = byte(i >> 16)
//	b[2] = byte(i >> 8)
//	b[3] = byte(i)
//	return b
//}
