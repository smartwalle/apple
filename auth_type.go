package apple

type User struct {
	Id             string `json:"id"`
	Issuer         string `json:"issuer"`
	BundleId       string `json:"bundle_id"`
	Email          string `json:"email"`
	EmailVerified  string `json:"email_verified"`
	IsPrivateEmail bool   `json:"is_private_email"`
	RealUserStatus int    `json:"real_user_status"`
	Nonce          string `json:"nonce"`
	AuthTime       int64  `json:"auth_time"`
	IssuedAt       int64  `json:"issued_at"`
	ExpiresAt      int64  `json:"expires_at"`
}
