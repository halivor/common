package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	mrand "math/rand"
	"time"

	p "github.com/halivor/common/golang/packet"
	_ "github.com/halivor/goutil/bufferpool"
)

var KEY = ""

func InitCrypto(key string) {
	KEY = key
}

func GeneToken(uid int64) *p.Token {
	key := make([]byte, 16) // TODO: 考虑使用固定key，或每个用户固定key
	sign := make([]byte, 16)
	io.ReadFull(rand.Reader, key)
	io.ReadFull(rand.Reader, sign)

	exp := time.Hour*24*30 + time.Duration(mrand.Int63n(3600))*time.Second

	return &p.Token{
		Uid:  uid,
		Key:  hex.EncodeToString(key),
		Sign: hex.EncodeToString(sign),
		Exp:  time.Now().Add(exp).Unix(),
		Rand: mrand.Int63(),
	}
}

func EncData(key, data []byte) string {
	switch {
	case len(key) != 16:
		panic(fmt.Sprintf("加密密钥[%s]长度[%d]错误",
			string(key), len(key)))
	case len(data) == 0:
		return ""
	case len(data)%aes.BlockSize != 0:
		data = append(data, bytes.Repeat([]byte(" "), aes.BlockSize-len(data)%aes.BlockSize)...)
	}
	c, _ := aes.NewCipher(key)

	enc_data := make([]byte, len(data))
	copy(enc_data, data[:aes.BlockSize])
	mode := cipher.NewCBCEncrypter(c, enc_data[:aes.BlockSize])
	mode.CryptBlocks(enc_data[aes.BlockSize:], data[aes.BlockSize:])
	return fmt.Sprintf("%x", enc_data)
}

func DecData(key, data []byte) string {
	switch {
	case len(key) != 16:
		panic(fmt.Sprintf("解密密钥[%s]长度[%d]错误",
			string(key), len(key)))
	case len(data) == 0:
		return ""
	case len(data) < aes.BlockSize, len(data)%aes.BlockSize != 0:
		panic(fmt.Sprintf("解密密钥[%s]长度[%d]错误",
			string(key), len(key)))
	}

	c, _ := aes.NewCipher(key)
	mode := cipher.NewCBCDecrypter(c, data[:aes.BlockSize])
	dec_data := make([]byte, len(data[aes.BlockSize:]))
	mode.CryptBlocks(dec_data, data[aes.BlockSize:])
	return string(dec_data)
}
