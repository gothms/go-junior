package service

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"testing"
	"time"
	"unicode/utf8"
)

func TestPasswordEncrypt(t *testing.T) {
	password := []byte("123456#hello")
	// 限制密码不能超过 72
	encrypted, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	assert.NoError(t, err)
	println(string(encrypted))

	err = bcrypt.CompareHashAndPassword(encrypted, password)
	assert.NoError(t, err)
}

func TestStringLength(t *testing.T) {
	s := "Go语言"
	length := len(s)
	runeCount := utf8.RuneCountInString(s)
	t.Logf("len=%d,runeCount=%d \n", length, runeCount)
	assert.NotEqual(t, length, runeCount, "长度len != 字节数")
}
func TestTimeParse(t *testing.T) {
	const format = "2006-01-02"
	//const format = time.DateOnly
	date := "2008-09-12"
	parse, err := time.Parse(format, date)
	t.Logf("parse: %v, error: %v", parse, err)
}
