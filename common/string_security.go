package common

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

type encryption struct {
	text string
}

// 字符串转MD5
func (s *encryption) EncryptMD5() string {
	str := &s.text
	ctx := md5.New()
	ctx.Write([]byte(*str))
	return hex.EncodeToString(ctx.Sum(nil))
}

// 字符串转SHA1
func (s *encryption) EncryptSHA1() string {
	str := &s.text
	ctx := sha1.New()
	ctx.Write([]byte(*str))
	return hex.EncodeToString(ctx.Sum(nil))
}

// 字符串转SHA256
func (s *encryption) EncryptSHA256() string {
	str := &s.text
	ctx := sha256.New()
	ctx.Write([]byte(*str))
	return hex.EncodeToString(ctx.Sum(nil))
}
