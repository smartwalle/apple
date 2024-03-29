package storekit

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/smartwalle/ncrypto"
	"strings"
)

const kRootPEM = `-----BEGIN CERTIFICATE-----
MIICQzCCAcmgAwIBAgIILcX8iNLFS5UwCgYIKoZIzj0EAwMwZzEbMBkGA1UEAwwS
QXBwbGUgUm9vdCBDQSAtIEczMSYwJAYDVQQLDB1BcHBsZSBDZXJ0aWZpY2F0aW9u
IEF1dGhvcml0eTETMBEGA1UECgwKQXBwbGUgSW5jLjELMAkGA1UEBhMCVVMwHhcN
MTQwNDMwMTgxOTA2WhcNMzkwNDMwMTgxOTA2WjBnMRswGQYDVQQDDBJBcHBsZSBS
b290IENBIC0gRzMxJjAkBgNVBAsMHUFwcGxlIENlcnRpZmljYXRpb24gQXV0aG9y
aXR5MRMwEQYDVQQKDApBcHBsZSBJbmMuMQswCQYDVQQGEwJVUzB2MBAGByqGSM49
AgEGBSuBBAAiA2IABJjpLz1AcqTtkyJygRMc3RCV8cWjTnHcFBbZDuWmBSp3ZHtf
TjjTuxxEtX/1H7YyYl3J6YRbTzBPEVoA/VhYDKX1DyxNB0cTddqXl5dvMVztK517
IDvYuVTZXpmkOlEKMaNCMEAwHQYDVR0OBBYEFLuw3qFYM4iapIqZ3r6966/ayySr
MA8GA1UdEwEB/wQFMAMBAf8wDgYDVR0PAQH/BAQDAgEGMAoGCCqGSM49BAMDA2gA
MGUCMQCD6cHEFl4aXTQY2e3v9GwOAEZLuN+yRhHFD/3meoyhpmvOwgPUnPWTxnS4
at+qIxUCMG1mihDK1A3UT82NQz60imOlM27jbdoXt2QfyFMm+YhidDkLF1vLUagM
6BgD56KyKA==
-----END CERTIFICATE-----
`

type Header struct {
	Alg string   `json:"alg"`
	X5C []string `json:"x5c"`
}

func DecodeClaims(payload string, claims jwt.Claims) error {
	headerBytes, err := base64.RawStdEncoding.DecodeString(strings.Split(payload, ".")[0])
	if err != nil {
		return err
	}

	//var data, _ = base64.RawStdEncoding.DecodeString(strings.Split(payload, ".")[1])
	//fmt.Println(string(data))

	var header *Header
	if err = json.Unmarshal(headerBytes, &header); err != nil {
		return err
	}

	rootCert, err := ncrypto.DecodeCertificate([]byte(kRootPEM))
	if err != nil {
		return err
	}

	intermediateCert, err := ncrypto.DecodeCertificate([]byte(header.X5C[1]))
	if err != nil {
		return err
	}

	cert, err := ncrypto.DecodeCertificate([]byte(header.X5C[0]))
	if err != nil {
		return err
	}

	if err = verifyCert(rootCert, intermediateCert, cert); err != nil {
		return err
	}

	if _, err = jwt.ParseWithClaims(payload, claims, func(token *jwt.Token) (interface{}, error) {
		switch publicKey := cert.PublicKey.(type) {
		case *ecdsa.PublicKey:
			return publicKey, nil
		default:
			return nil, errors.New("key is not a valid *ecdsa.PublicKey")
		}
	}); err != nil {
		return err
	}
	return nil
}

func verifyCert(rootCert, intermediateCert, cert *x509.Certificate) error {
	var roots = x509.NewCertPool()
	roots.AddCert(rootCert)

	var intermediates = x509.NewCertPool()
	intermediates.AddCert(intermediateCert)
	intermediates.AddCert(rootCert)

	var opts = x509.VerifyOptions{
		Roots:         roots,
		Intermediates: intermediates,
	}
	if _, err := cert.Verify(opts); err != nil {
		return err
	}
	return nil
}
