package main

import (
	"src/config"
	"src/controller"
	"src/util"
	"src/util/log"
	"time"
)

func main() {
	startTimer(MainCheckInFunc)
}

func startTimer(f func()) {
	func() {
		for {
			f()
			now := time.Now()
			// 计算下一个时间点
			next := now.Add(time.Duration(config.C.TimeConf.StartNextDay) * time.Hour * 24)
			hour := config.C.TimeConf.StartHour
			min := config.C.TimeConf.StartMin
			next = time.Date(next.Year(), next.Month(), next.Day(), hour, min, 0, 0, next.Location())
			t := time.NewTimer(next.Sub(now))
			log.Logger.Infof("下一次开始执行程序为%v", next.Format("2006-01-02 15:04:05"))
			<-t.C
		}
	}()
}

func MainCheckInFunc() {
	//return func() {
	username := config.C.User.Username
	passwd := config.C.User.Passwd

	rsp := util.Login(username, passwd)
	userinfo := controller.HandleLogin(rsp)

	rsp = util.GetUserCourses(userinfo)
	AllClass := controller.HandleGetUserCourses(rsp)

	try := 1
	duration:=config.C.TimeConf.Duration
	log.Logger.Infof("执行程序,持续%v分钟...", duration)
	for i := time.Duration(0); i <= time.Duration(duration); i += config.TimeSleep {
		log.Logger.Infof("开始自动签到，时间间隔为%v\n", config.TimeSleep)
		log.Logger.Infof("第%v次尝试\n", try)
		for _, class := range AllClass {
			rsp = util.IsCheckInOpen(userinfo, class.ClassId)
			rst, msg := controller.HandleIsCheckInOpen(rsp)
			log.Logger.Infof("%v：result_msg:%v\n", class.ClassName, msg)
			if rst {
				count := 1
				for {
					//log.Logger.Info("%v：获取签到列表\n",class.ClassName)
					//rsp = util.GetCheckInList(userinfo, class.ClassId)
					//checkInId = controller.HandleGetCheckInList(rsp)
					log.Logger.Infof("%v:开始签到\n", class.ClassName)
					rsp = util.CheckIn(userinfo, class.ClassId)
					rst, msg := controller.HandleCheckIn(rsp)
					log.Logger.Infof("%v:result_msg:%v\n", class.ClassName, msg)
					if rst {
						log.Logger.Infof("%v:签到成功！\n", class.ClassName)
						break
					} else if count > 10 {
						log.Logger.Infof("Attention:%v签到失败，请检查网络，重新尝试，否则你可能失去生命！\n", class.ClassName)
						break
					} else {
						log.Logger.Infof("%v签到失败，第%v次重试\n", class.ClassName, count)
						count++
					}
					time.Sleep(time.Second * 5)
				}
			} else {
				log.Logger.Infof("%v：签到未开始\n", class.ClassName)
			}
		}
		try++
		time.Sleep(config.TimeSleep)
	}
	//}
}
