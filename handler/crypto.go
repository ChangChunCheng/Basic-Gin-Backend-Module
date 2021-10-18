// Package handler
package handler

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"reflect"
	"unsafe"

	"github.com/spf13/viper"
)

// byte16: AES-128, byte24: AES-192, byte32: ARS-256
var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

// GetSecret - encrypt the string (for password)
// response - secret string, err error
func GetSecret(text string) (secret string, err error) {
	// 1. Encode to base64
	debyte := base64Encode([]byte(text))
	// 2. Encrypt
	textSecret := debyte // []byte(text)
	c, err := aes.NewCipher([]byte(viper.GetString("auth.passwordKey")))
	if err != nil {
		return "", err
	}
	cfb := cipher.NewCFBEncrypter(c, commonIV)
	ciphertext := make([]byte, len(debyte))
	cfb.XORKeyStream(ciphertext, textSecret)
	// 3. Bytes to string
	secret = bytes2String(ciphertext)
	return secret, nil
}

func ParseSecret(secret string) (text string, err error) {
	// 1. Decrypt
	c, err := aes.NewCipher([]byte(viper.GetString("auth.passwordKey")))
	if err != nil {
		return "", err
	}
	cfb := cipher.NewCFBEncrypter(c, commonIV)
	plaintext := make([]byte, len(secret))
	cfb.XORKeyStream(plaintext, []byte(secret))
	// 2. Decode from base64
	enbyte, err := base64Decode(plaintext)
	if err != nil {
		return "", err
	}
	// 3. Bytes to string
	text = bytes2String(enbyte)
	return text, nil
}

// string2Bytes - Transfer string to bytes
func string2Bytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

// bytes2String - Transfer bytes to string
func bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// base64Encode - encode to base64
func base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

// base64Decode - decode from base64
func base64Decode(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(src))
}
