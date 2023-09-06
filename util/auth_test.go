package yizuutil

import "testing"

func TestAuthCode(t *testing.T) {
	// 验证通过
	phoneNum := "12345"
	captcha := "666666"
	SendAuthCode(phoneNum, captcha)
}
