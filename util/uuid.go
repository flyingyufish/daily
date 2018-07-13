package util

import (
	"math/rand"

	uuid "github.com/satori/go.uuid"
)

// GetUUID 生成uuid
func GetUUID() (string, error) {
	u, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return u.String(), nil
}

// RandInt 在0<= x <=max 的一个随机数
func RandInt(max int) int {
	return rand.Intn(max)
}
