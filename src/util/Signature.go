package util

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"strings"
	"time"
)

func GetCourseSignature(url string, userid string, ftime string, accessSecret string) string {
	str := fmt.Sprintf("%v|%v|%v", url, userid, ftime)
	//fmt.Println(str)
	sign := GetHMACSha1(accessSecret, str)
	//fmt.Println(sign)
	return sign
}

func GetCheckInListSignature(url string, userid string, ftime string, accessSecret string, classid string) string {
	mdstr := strings.ToUpper(GetMD5(fmt.Sprintf("clazz_course_id=%v|page=1|role_id=2", classid)))
	str := fmt.Sprintf("%v|%v|%v|%v", url, userid, ftime, mdstr)
	//fmt.Println(str)
	sign := GetHMACSha1(accessSecret, str)
	//fmt.Println(sign)
	return sign
}

func GetCheckInSignature(url string, userid string, ftime string, accessSecret string, classid string) string {
	mdstr := strings.ToUpper(GetMD5(fmt.Sprintf("cc_id=%v", classid)))
	str := fmt.Sprintf("%v|%v|%v|%v", url, userid, ftime,mdstr)
	//fmt.Println(str)
	sign := GetHMACSha1(accessSecret, str)
	//fmt.Println(sign)
	return sign
}

func GetIsCheckInOpenSignature(url string, userid string, ftime string, accessSecret string, classid string) string {
	mdstr := strings.ToUpper(GetMD5(fmt.Sprintf("clazz_course_id=%v", classid)))
	str := fmt.Sprintf("%v|%v|%v|%v", url, userid, ftime, mdstr)
	//fmt.Println(str)
	sign := GetHMACSha1(accessSecret, str)
	//fmt.Println(sign)
	return sign
}

func GetDate() string {
	date := time.Now()
	t, _ := time.ParseDuration("-8h")
	tim := date.Add(t).Format("Mon, 02 Jan 2006 15:04:05") + " GMT+00:00"
	//fmt.Println(tim)
	return tim
}

func GetHMACSha1(key string, data string) string {
	hm := hmac.New(sha1.New, []byte(key))
	hm.Write([]byte(data))
	return hex.EncodeToString(hm.Sum(nil))
}

func GetMD5(data string) string {
	w := md5.New()
	_, _ = io.WriteString(w, data)
	bw := w.Sum(nil)
	md5str := hex.EncodeToString(bw)
	return md5str
}
