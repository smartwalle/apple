package auth

import "github.com/golang-jwt/jwt/v5"

type Header struct {
	Kid string `json:"kid"`
	Alg string `json:"alg"`
}

type Claims struct {
	jwt.RegisteredClaims
	CHash          string `json:"c_hash"`
	AuthTime       int    `json:"auth_time"`
	Nonce          string `json:"nonce"`
	NonceSupported bool   `json:"nonce_supported"`
	Email          string `json:"email"`
	EmailVerified  string `json:"email_verified"`
	IsPrivateEmail bool   `json:"is_private_email"`
	RealUserStatus int    `json:"real_user_status"`
	TransferSub    string `json:"transfer_sub"`
}

type Key struct {
	Kty string `json:"kty"`
	Kid string `json:"kid"`
	Use string `json:"use"`
	Alg string `json:"alg"`
	N   string `json:"n"`
	E   string `json:"e"`
}
