package config

import "time"

//CheckInApi     = "https://checkin.mosoteach.cn:19528/checkin"
//CheckInHost    = "checkin.mosoteach.cn:19528"

const (
	LoginApi       = "https://www.mosoteach.cn/web/index.php?c=passport&m=account_login"
	CourseApi      = "http://api.mosoteach.cn/mssvc/index.php/cc/list_joined"
	CheckInListApi = "https://api.mosoteach.cn/mssvc/index.php/checkin/index"
	CheckInApi     = "https://api.mosoteach.cn/mssvc/index.php/cc_clockin/clockin"
	MainHost       = "api.mosoteach.cn"
	IsCheckInOPen  = "https://api.mosoteach.cn/mssvc/index.php/checkin/current_open"
	TimeSleep      = time.Minute * 3
)

var C *Config

type Config struct {
	User     user     `yaml:"user"`
	Location location `yaml:"location"`
	LogConf  logConf  `yaml:"logconf"`
	TimeConf timeConf `yaml:"timeconf"`
	Debug    bool     `yaml:"debug"`
}

type user struct {
	Username string `yaml:"username"`
	Passwd   string `yaml:"password"`
}

type location struct {
	Lat string `yaml:"lat"`
	Lng string `yaml:"lng"`
}

type logConf struct {
	LogPath     string `yaml:"logpath"`
	LogFileName string `yaml:"logfilename"`
}

type timeConf struct {
	StartHour    int `yaml:"start_hour"`
	StartMin     int `yaml:"start_min"`
	StartNextDay int `yaml:"start_next_day"`
	Duration     int `yaml:"duration"`
}

type Class struct {
	ClassName string
	ClassId   string
}

type UserInfo struct {
	UserId            string
	AccessSecret      string
	AccessId          string
	LastSecUpdateTime string
	LastSecUpdateTsS  string
}
