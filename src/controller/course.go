package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"src/config"
	"src/util/log"
)

func HandleGetUserCourses(rsp *http.Response) []config.Class {
	content, _ := ioutil.ReadAll(rsp.Body)
	//rspBody := string(content)
	log.Logger.Debug(string(content))
	r := make(map[string]interface{})
	jsonErr := json.Unmarshal(content, &r)
	if jsonErr != nil {
		log.Logger.Fatalln(jsonErr)
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
