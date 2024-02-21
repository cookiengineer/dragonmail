package structs

type Account struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	OTP      struct {
		// TODO? No Idea
	} `json:"otp"`
	Config Config `json:"config"`
}
