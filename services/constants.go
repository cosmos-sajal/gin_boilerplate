package authservice

const (
	OTP_KEY_PREFIX         = "OTP_KEY_"
	OTP_KEY_EXPIRY         = 120 // in seconds
	OTP_MAX_ATTEMPT        = 5
	OTP_ATTEMPTS_PREFIX    = "OTP_ATTEMPTS_"
	OTP_ATTEMPT_KEY_EXPIRY = 3600 // in seconds
)
