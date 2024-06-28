package main

import (
	"crypto/rand"
	"fmt"
	"github.com/ghafbu/ghafotp/pkg/findnetworkpkg"
	"github.com/ghafbu/ghafotp/pkg/tsotp"
)

// تابع برای تولید یک رشته تصادفی با طول مشخص
func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	rand.Read(b)
	for i := range b {
		b[i] = charset[b[i]%byte(len(charset))]
	}
	return string(b)
}

func main() {
	findnetworkpkg.Get()
	secretKey := generateRandomString(16) // به طول 16 بایت
	// Example usage:
	otp, err := tsotp.GenerateOTP(secretKey)
	if err != nil {
		fmt.Println("error OTP:", err)
		return
	}
	fmt.Println("OTP code:", otp)

	//timewindowotppkg.Run()
	//
	//reader := bufio.NewReader(os.Stdin)
	//fmt.Print("INSERT OTP CODE ::")
	//code, _ := reader.ReadString('\n')
	//code = strings.TrimSpace(code)
	//
	//// اعتبارسنجی کد OTP
	//valid, err := timewindowotppkg.ValidateOTP(timewindowotppkg.Secret, code)
	//if err != nil {
	//	fmt.Println("otp-validate:", err)
	//	return
	//}
	//if valid {
	//	fmt.Println("otp is valid")
	//} else {
	//	fmt.Println("otp notValid")
	//}
}
