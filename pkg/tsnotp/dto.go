package tsnotp

type RequestGetStruct struct {
	Mobile string `json:"mobile"`
}

type RequestVerifyStruct struct {
	Otp    string `json:"otp"`
	Mobile string `json:"mobile"`
}
