package utl

import (
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

func GenerationSecretKey(appName string, mobile string) (*otp.Key, error) {
	secretKey, err := totp.Generate(totp.GenerateOpts{
		Issuer:      appName,
		AccountName: mobile,
		//Period:      120,
	})
	if err != nil {
		return secretKey, err
		//fmt.Println("generation key error:", err)
	}

	return secretKey, nil
}
