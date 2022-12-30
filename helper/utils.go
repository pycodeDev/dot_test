package helper

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"hash"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func RemoveWhiteSpace(val string) string {
	s := strings.TrimSpace(val)
	return s
}

func ValidasiPass(pass string) error {

	if len(pass) < 8 {
		return errors.New("Mininmal password 8 karakter!")
	}

	var pass1, _ = regexp.Compile("[0-9]")
	pass1_check := pass1.MatchString(pass)
	if !pass1_check {
		return errors.New("Password setidaknya harus mengandung angka!")
	}

	var pass2, _ = regexp.Compile("[a-z]")
	pass2_check := pass2.MatchString(pass)
	if !pass2_check {
		return errors.New("Password setidaknya harus mengandung huruf!")
	}

	var pass3, _ = regexp.Compile("[A-Z]")
	pass3_check := pass3.MatchString(pass)
	if !pass3_check {
		return errors.New("Password setidaknya harus mengandung huruf besar!")
	}

	var pass4, _ = regexp.Compile(`[\W]`)
	pass4_check := pass4.MatchString(pass)
	if !pass4_check {
		return errors.New("Password setidaknya harus mengandung simbol!")
	}
	return nil
}

func SelisihWaktuPerMenit(value string) int {

	layoutFormat1 := "2006-01-02 15:04:05"
	value1 := fmt.Sprintf("%v", GetCurrentDateTime())
	date1, _ := time.Parse(layoutFormat1, value1)

	layoutFormat := "2006-01-02 15:04:05"
	date, _ := time.Parse(layoutFormat, value)
	daa := date1.Sub(date)
	return int(daa.Minutes())
}

func GetCurrentDate() string {
	//result example 2021-08-21
	now := time.Now()
	tanggal := fmt.Sprintf("%v", now.Format("06-01-02"))
	return tanggal
}

func GetCurrentDateFull() string {
	//result example 2021-08-21
	now := time.Now()
	tanggal := fmt.Sprintf("%v", now.Format("2006-01-02"))
	return tanggal
}

func GetCurrentTime() string {
	//ex : 15:02:59
	return fmt.Sprintf("%v", time.Now().Format("15:04:05"))
}

func GetCurrentDateTime() string {
	//ex : 2021-08-21 15:02:59
	return fmt.Sprintf("%v", time.Now().Format("2006-01-02 15:04:05"))
}

func CutStringV2(val string, jumlah int) string {
	res := val[jumlah:len(val)]
	return res
}

func StringContains(real string, s string) bool {
	if strings.Contains(real, s) {
		return true
	}
	return false
}

func GenerateSHA256(text string) (string, error) {
	s := text
	h := sha256.New()
	h.Write([]byte(s))
	sha256Hash := hex.EncodeToString(h.Sum(nil))
	return sha256Hash, nil
	//fmt.Println(s, sha1_hash)
}

func EncSha1(text string) string {
	h := sha1.New()
	h.Write([]byte(text))
	sha1_hash := hex.EncodeToString(h.Sum(nil))
	return sha1_hash
}

func SHA256(text string) string {
	// https://github.com/shomali11/util/blob/91c54758c87ba2b8006eda5b0d1b7293c4c60c40/xhashes/xhashes.go
	algorithm := sha256.New()
	return stringHasher(algorithm, text)
}

func stringHasher(algorithm hash.Hash, text string) string {
	algorithm.Write([]byte(text))
	return hex.EncodeToString(algorithm.Sum(nil))
}

func LengthString(val string) int {
	length := len([]rune(val))
	return length
}

func IntToString(val int) string {
	s := strconv.Itoa(val)
	return s
}

func StringToInt(val string) int {
	data, err := strconv.Atoi(val)
	err = err
	return data
}
func ExplodeString(val string, ex string) []string {
	s := strings.Split(val, ex)
	return s
}

func ReplaceAllString(val string, search string, rep string) string {
	length := len(val)
	fmt.Println(length)
	ret := strings.ReplaceAll(val, search, rep)
	return ret
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func HashPass(text string) string {
	old_hash := GetMD5Hash(text)
	var sha = sha1.New()
	sha.Write([]byte(old_hash))
	var encrypted = sha.Sum(nil)
	var encryptedString = fmt.Sprintf("dot%x", encrypted)
	password := GetMD5Hash(fmt.Sprintf("2k22%s%s", text, encryptedString))

	return password
}

func RandStringBytes(n int) string {
	letterBytes := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func RandStringNumber(n int) string {
	letterBytes := "1234567890"
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	// letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[seededRand.Intn(len(letterBytes))]
	}
	return string(b)
}
