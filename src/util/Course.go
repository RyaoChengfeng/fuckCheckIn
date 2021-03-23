package util

import (
	"log"
	"net/http"
	"src/config"
	"time"
)

func GetUserCourses(userinfo config.UserInfo) *http.Response {
	courseUrl:=config.CourseApi
	userAgent := config.UserAgent
	println(userAgent)
	client := &http.Client{
		Timeout: 5*time.Second,
	}
	req ,err := http.NewRequest(http.MethodPost,courseUrl,nil)
	if err != nil {
		log.Fatalln(err)
	}
	ti:=GetDate()
	req.Header.Set("User-Agent","Dalvik/2.1.0 (Linux; U; Android 8.1.0; ONE A2001 Build/OPM7.181205.001)")
	req.Header.Set("X-scheme","https")
	req.Header.Set("X-app-id","MTANDROID")
	req.Header.Set("X-app-version","5.1.1")
	req.Header.Set("X-dpr","2.7")
	req.Header.Set("X-app-machine","ONE A2001")
	req.Header.Set("X-app-system-version","8.1.0")
	req.Header.Set("Host",config.MainHost)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Date",ti)
	req.Header.Set("X-mssvc-signature",GetCourseSignature(courseUrl,userinfo.UserId,ti,userinfo.AccessSecret))
	req.Header.Set("X-mssvc-sec-ts",userinfo.LastSecUpdateTsS)
	req.Header.Set("X-mssvc-access-id",userinfo.AccessId)

	//fmt.Println(req.Header)

	rsp,err:= client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	return rsp
}

