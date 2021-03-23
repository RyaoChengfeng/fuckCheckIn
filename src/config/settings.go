package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

func init() {
	configFile := "config.yml"

	// 如果有设置 ENV ，则使用ENV中的环境
	if v, ok := os.LookupEnv("ENV"); ok {
		configFile = v + ".yml"
	}

	data, err := ioutil.ReadFile(fmt.Sprintf("../env/%s", configFile))
	if err != nil {
		panic(err)
		return
	}

	config := &Config{}

	err = yaml.Unmarshal(data, config)
	if err != nil {
		fmt.Println("Unmarshal config error!")
		panic(err)
		return
	}
	C=config
	fmt.Println("Config " + configFile + " loaded.")
}