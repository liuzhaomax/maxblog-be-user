package core

import (
	"crypto/md5"
	"github.com/anaskhan96/go-password-encoder"
	"sync"
)

var ctx *Context
var once sync.Once

func init() {
	once.Do(func() {
		ctx = &Context{}
	})
}

func GetInstanceOfContext() *Context {
	return ctx
}

type Context struct {
	PwdEncodingOpts *password.Options
}

func SetPwdEncodingOpts() {
	ctx.PwdEncodingOpts = &password.Options{
		SaltLen:      16,
		Iterations:   64,
		KeyLen:       16,
		HashFunction: md5.New,
	}
}

func GetEncodedPwd(pwd string) (string, string) {
	salt, encodedPwd := password.Encode(pwd, ctx.PwdEncodingOpts)
	return salt, encodedPwd
}

func VerifyEncodedPwd(pwdHeldRaw string, salt string, pwdTarget string) bool {
	return password.Verify(pwdHeldRaw, salt, pwdTarget, ctx.PwdEncodingOpts)
}
