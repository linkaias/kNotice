package help

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

//FilteredSQLInject sql注入判断
func FilteredSQLInject(oldStr string) string {
	ok := true
	str := `(?:')|(?:--)|(/\\*(?:.|[\\n\\r])*?\\*/)|(\b(select|update|and|or|delete|insert|trancate|char|chr|into|substr|ascii|declare|exec|count|master|into|drop|execute)\b)`
	re, err := regexp.Compile(str)
	if err != nil {
		//panic(err.Error())
		ok = false
		return ""
	}
	ok = re.MatchString(oldStr)
	if ok {
		return ""
	} else {
		return oldStr
	}
}

//CreateValidateCode 生成验证码
//width 验证码长度
func CreateValidateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}
