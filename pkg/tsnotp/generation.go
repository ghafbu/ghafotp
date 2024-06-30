package tsnotp

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

// تابع تولید TSNOTP
func generate(secret string, timestamp int64, sequenceNumber uint64, networkInfo string) string {
	key, _ := hex.DecodeString(secret)

	// ترکیب timestamp و شماره ترتیبی و اطلاعات شبکه
	combined := make([]byte, 24)
	binary.BigEndian.PutUint64(combined[:8], uint64(timestamp))
	binary.BigEndian.PutUint64(combined[8:], sequenceNumber)
	copy(combined[16:], []byte(networkInfo)[:8]) // استفاده از 8 بایت اول اطلاعات شبکه

	// ایجاد HMAC-SHA1 از مقدار ترکیب شده
	hmacSha1 := hmac.New(sha1.New, key)
	hmacSha1.Write(combined)
	hash := hmacSha1.Sum(nil)

	// برش پویا برای گرفتن کد 6 رقمی OTP
	offset := hash[len(hash)-1] & 0xf
	binaryCode := (int(hash[offset])&0x7f)<<24 |
		(int(hash[offset+1])&0xff)<<16 |
		(int(hash[offset+2])&0xff)<<8 |
		(int(hash[offset+3]) & 0xff)

	otp := binaryCode % 1000000

	return fmt.Sprintf("%06d", otp)
}