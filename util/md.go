package util

import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"
)

// Md5 加密
func Md5(str string) string {
	m := md5.New()
	io.WriteString(m, str)
	return strings.ToLower(fmt.Sprintf("%x", m.Sum(nil)))
}
