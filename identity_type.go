package apple

type User struct {
	Id             string `json:"id"`
	BundleId       string `json:"bundle_id"`
	AuthTime       int    `json:"auth_time"`
	Email          string `json:"email"`
	EmailVerified  string `json:"email_verified"`
	IsPrivateEmail bool   `json:"is_private_email"`
	RealUserStatus int    `json:"real_user_status"`
}
