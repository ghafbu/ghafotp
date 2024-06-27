package main

import (
	"bufio"
	"fmt"
	"github.com/ghafbu/ghafotp/pkg/timewindowotppkg"
	"os"
	"strings"
)

func main() {
	timewindowotppkg.Run()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("INSERT OTP CODE ::")
	code, _ := reader.ReadString('\n')
	code = strings.TrimSpace(code)

	// اعتبارسنجی کد OTP
	valid, err := timewindowotppkg.ValidateOTP(timewindowotppkg.Secret, code)
	if err != nil {
		fmt.Println("otp-validate:", err)
		return
	}
	if valid {
		fmt.Println("otp is valid")
	} else {
		fmt.Println("otp notValid")
	}
}
