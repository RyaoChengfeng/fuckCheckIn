package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"src/config"
)

func HandleGetUserCourses(rsp *http.Response) []config.Class {
	content, _ := ioutil.ReadAll(rsp.Body)
	//rspBody := string(content)
	//fmt.Println(rspBody)
	r := make(map[string]interface{})
	jsonErr := json.Unmarshal(content, &r)
	if jsonErr != nil {
		log.Fatalln(jsonErr)
	}
	AllClass := r["rows"].([]interface{})
	var Classes []config.Class
	for _, value := range AllClass {
		classId := value.(map[string]interface{})["id"].(string)
		className := value.(map[string]interface{})["course"].(map[string]interface{})["name"].(string)
		class := config.Class{
			ClassName: className,
			ClassId:   classId,
		}
		Classes = append(Classes, class)
	}
	return Classes
}
