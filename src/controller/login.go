package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"src/config"
	"src/util/log"
)

func HandleLogin(rsp *http.Response) config.UserInfo {
	content, _ := ioutil.ReadAll(rsp.Body)
	//rspBody := string(content)
	log.Logger.Debug(string(content))
	r := make(map[string]interface{})
	jsonErr := json.Unmarshal(content, &r)
	if jsonErr != nil {
		log.Logger.Fatalln(jsonErr)
	}

	userinfo := config.UserInfo{
		UserId:            r["user"].(map[string]interface{})["user_id"].(string),
		AccessSecret:      r["user"].(map[string]interface{})["access_secret"].(string),
		AccessId:          r["user"].(map[string]interface{})["access_id"].(string),
		LastSecUpdateTime: r["user"].(map[string]interface{})["last_sec_update_time"].(string),
		LastSecUpdateTsS:  r["user"].(map[string]interface{})["last_sec_update_ts_s"].(string),
	}
	return userinfo
}
