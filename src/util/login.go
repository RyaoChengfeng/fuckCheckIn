package util

import (
	"log"
	"net/http"
	"net/url"
	"src/config"
	"strings"
	"time"
)

func Login(username string,passwd string) *http.Response {
	LoginUrl:= config.LoginApi
	userAgent := config.UserAgent
	formData := url.Values{
		"account_name": {username},
		"user_pwd": {passwd},
		"remember_me": {"N"},
	}
	body:= strings.NewReader(formData.Encode())
	client := &http.Client{
		Timeout: 5*time.Second,
	}
	req ,err := http.NewRequest(http.MethodPost,LoginUrl,body)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("User-Agent",userAgent)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rsp,err:= client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	return rsp
}
