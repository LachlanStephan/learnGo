package validator

import (
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

type Validator struct {
	FormErrors    map[string]string
	NonFormErrors []string
}

var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func (v *Validator) Valid() bool {
	return len(v.FormErrors) == 0 && len(v.NonFormErrors) == 0
}

func (v *Validator) AddNonFieldError(message string) {
	v.NonFormErrors = append(v.NonFormErrors, message)
}

func (v *Validator) AddFieldError(key, message string) {
	if v.FormErrors == nil {
		v.FormErrors = make(map[string]string)
	}

	if _, exists := v.FormErrors[key]; !exists {
		v.FormErrors[key] = message
	}
}

func (v *Validator) CheckField(ok bool, key, message string) {
	if !ok {
		v.AddFieldError(key, message)
	}
}

func NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

func MinChars(value string, n int) bool {
	return utf8.RuneCountInString(value) >= n
}

func MaxChars(value string, n int) bool {
	return utf8.RuneCountInString(value) <= n
}

func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

func PermittedValue[T comparable](targetValue T, permittedValues ...T) bool {
	for i := range permittedValues {
		if targetValue == permittedValues[i] {
			return true
		}
	}
	return false
}

func CastUserId(userId string) int {
	uid, err := strconv.Atoi(userId)
	if err != nil {
		return 0
	}
	return uid
}

func ValidUserId(userId int, validUsers []int) bool {
	if userId == 0 || !isInt(userId) {
		return false
	}

	return PermittedValue(userId, validUsers...)
}

func isInt(n any) bool {
	t := reflect.TypeOf(n).Kind()
	return t == reflect.Int
}
