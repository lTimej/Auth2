package md5

import (
	"crypto/md5"
)

func Md5(s string) string {
	h := md5.New()
	return string(h.Sum([]byte(s)))
}
