package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//func HandleGetCheckInList(rsp *http.Response) string {
//	content, _ := ioutil.ReadAll(rsp.Body)
//	rspBody := string(content)
//	fmt.Println(rspBody)
//	r := make(map[string]interface{})
//	jsonErr := json.Unmarshal(content, &r)
//	if jsonErr != nil {
//		log.Fatalln(jsonErr)
//	}
//	checkInId:=r["data"].([]interface{})[0].(map[string]interface{})["checkin_id"].(string)
//	return checkInId
//}

func HandleIsCheckInOpen(rsp *http.Response) (bool,string) {
	content, _ := ioutil.ReadAll(rsp.Body)
	rspBody := string(content)
	fmt.Println(rspBody)
	r := make(map[string]interface{})
	jsonErr := json.Unmarshal(content, &r)
	if jsonErr != nil {
		log.Fatalln(jsonErr)
	}
	rst:= int(r["result_code"].(float64))
	msg := r["result_msg"].(string)
	if rst==0 {
		return true,msg
	} else {
		return false,msg
	}
}

func HandleCheckIn(rsp *http.Response) (bool,string) {
	content, _ := ioutil.ReadAll(rsp.Body)
	rspBody := string(content)
	fmt.Println(rspBody)
	r := make(map[string]interface{})
	jsonErr := json.Unmarshal(content, &r)
	if jsonErr != nil {
		log.Fatalln(jsonErr)
	}
	rst:=int(r["result_code"].(float64))
	msg:=r["result_msg"].(string)
	if rst==0 {
		return true,msg
	} else {
		return false,msg
	}
}