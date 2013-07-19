package goaws

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

type Auth struct {
	accessKey string
	secretKey string
	algorithm string
}

func (a *Auth) setCredentials(AccessKey string, SecretKey string) {
	a.accessKey = AccessKey
	a.secretKey = SecretKey
	a.algorithm = "HmacSHA256"
}

func (a *Auth) authorize(message string) string {
	mac := hmac.New(sha256.New, []byte(a.secretKey))
	mac.Write([]byte(message))
	cryptedString := mac.Sum(nil)
	data := base64.StdEncoding.EncodeToString(cryptedString)
	return string(data)
}

func (a *Auth) GetHeader(date string) (string, error) {
	return fmt.Sprintf("AWS3-HTTPS AWSAccessKeyId=%s,Algorithm=%s,Signature=%s", a.accessKey, a.algorithm, a.authorize(date)), nil
}
