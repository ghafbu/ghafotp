package utl

import "time"

// تابع برای محاسبه مقادیر timestamp تا 120 ثانیه قبل
func OldTimeStamp(initialTimestamp int64) []int64 {
	// تبدیل timestamp اولیه به زمان معمولی (DateTime)
	t := time.Unix(initialTimestamp, 0)

	// تعریف یک آرایه برای نگهداری مقادیر timestamp ها
	timestamps := make([]int64, 0)

	// افزودن timestamp اولیه به آرایه
	timestamps = append(timestamps, initialTimestamp)

	// محاسبه timestamp های 120 ثانیه قبل
	for i := 1; i <= 120; i++ {
		t = t.Add(-time.Second)
		timestamps = append(timestamps, t.Unix())
	}

	return timestamps
}

//
//// مقدار timestamp اولیه
//initialTimestamp := int64(1719594812)
//
//// فراخوانی تابع برای دریافت مقادیر timestamp ها
//result := getTimestamps(initialTimestamp)
//
//// چاپ نتیجه
//fmt.Printf("Timestamp اولیه: %d\n", result[0])
//fmt.Printf("Timestamp 1 ثانیه قبل: %d\n", result[1])
//fmt.Printf("Timestamp 120 ثانیه قبل: %d\n", result[120])
