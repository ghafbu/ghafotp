package timewindowotppkg

import "fmt"

func Run() {
	// تولید کد OTP
	otp, err := GenerateOTP(Secret)
	if err != nil {
		fmt.Println("otp-generate:", err)
		return
	}
	fmt.Println("otp code=", otp)
}
