package util

import (
	"log"
	"net/http"
	"net/url"
	"src/config"
	"strings"
	"time"
)

func GetCheckInList(userinfo config.UserInfo, classid string) *http.Response {
	CheckInListUrl := config.CheckInListApi
	formData := url.Values{
		"clazz_course_id": {classid},
		"page":            {"1"},
		"role_id":         {"2"},
	}
	body := strings.NewReader(formData.Encode())
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	req, err := http.NewRequest(http.MethodPost, CheckInListUrl, body)
	if err != nil {
		log.Fatalln(err)
	}
	ti := GetDate()
	req.Header.Set("User-Agent", "Dalvik/2.1.0 (Linux; U; Android 8.1.0; ONE A2001 Build/OPM7.181205.001)")
	req.Header.Set("X-scheme", "https")
	req.Header.Set("X-app-id", "MTANDROID")
	req.Header.Set("X-app-version", "5.1.1")
	req.Header.Set("X-dpr", "2.7")
	req.Header.Set("X-app-machine", "ONE A2001")
	req.Header.Set("X-app-system-version", "8.1.0")
	req.Header.Set("Host", config.MainHost)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Date", ti)
	req.Header.Set("X-mssvc-signature", GetCheckInListSignature(CheckInListUrl, userinfo.UserId, ti, userinfo.AccessSecret, classid))
	req.Header.Set("X-mssvc-sec-ts", userinfo.LastSecUpdateTsS)
	req.Header.Set("X-mssvc-access-id", userinfo.AccessId)
	rsp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	return rsp
}

func CheckIn(userinfo config.UserInfo, classid string) *http.Response {
	CheckInUrl := config.CheckInApi
	//formData := url.Values{
	//	"checkin_id":      {checkinid},
	//	"report_pos_flag": {"Y"},
	//	"lat":             {lat},
	//	"lng":             {lng},
	//}
	// 发现现在签到已经不需要location了
	formData:=url.Values{
		"cc_id" : {classid},
	}
	body := strings.NewReader(formData.Encode())
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	req, err := http.NewRequest(http.MethodPost, CheckInUrl, body)
	if err != nil {
		log.Fatalln(err)
	}
	ti := GetDate()
	req.Header.Set("User-Agent", "Dalvik/2.1.0 (Linux; U; Android 8.1.0; ONE A2001 Build/OPM7.181205.001)")
	req.Header.Set("X-scheme", "https")
	req.Header.Set("X-app-id", "MTANDROID")
	req.Header.Set("X-app-version", "5.1.1")
	req.Header.Set("X-dpr", "2.7")
	req.Header.Set("X-app-machine", "ONE A2001")
	req.Header.Set("X-app-system-version", "8.1.0")

	req.Header.Set("Host", config.MainHost)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Date", ti)
	req.Header.Set("X-mssvc-signature", GetCheckInSignature(CheckInUrl, userinfo.UserId, ti, userinfo.AccessSecret,classid))
	req.Header.Set("X-mssvc-sec-ts", userinfo.LastSecUpdateTsS)
	req.Header.Set("X-mssvc-access-id", userinfo.AccessId)
	//fmt.Println(req.Header)

	rsp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	return rsp
}

func IsCheckInOpen(userinfo config.UserInfo,classid string) *http.Response {
	IsCheckInOpenUrl := config.IsCheckInOPen
	formData := url.Values{
		"clazz_course_id": {classid},
	}
	body := strings.NewReader(formData.Encode())
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	req, err := http.NewRequest(http.MethodPost, IsCheckInOpenUrl, body)
	if err != nil {
		log.Fatalln(err)
	}
	ti := GetDate()
	req.Header.Set("User-Agent", "Dalvik/2.1.0 (Linux; U; Android 8.1.0; ONE A2001 Build/OPM7.181205.001)")
	req.Header.Set("X-scheme", "https")
	req.Header.Set("X-app-id", "MTANDROID")
	req.Header.Set("X-app-version", "5.1.1")
	req.Header.Set("X-dpr", "2.7")
	req.Header.Set("X-app-machine", "ONE A2001")
	req.Header.Set("X-app-system-version", "8.1.0")

	req.Header.Set("Host", config.MainHost)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Date", ti)
	req.Header.Set("X-mssvc-signature", GetIsCheckInOpenSignature(IsCheckInOpenUrl, userinfo.UserId, ti, userinfo.AccessSecret,classid))
	req.Header.Set("X-mssvc-sec-ts", userinfo.LastSecUpdateTsS)
	req.Header.Set("X-mssvc-access-id", userinfo.AccessId)
	//fmt.Println(req.Header)

	rsp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	return rsp
}