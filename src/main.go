package main

import (
	"fmt"
	"src/config"
	"src/controller"
	"src/util"
	"time"
)

func main() {
	username := config.C.User.Username
	passwd := config.C.User.Passwd

	rsp := util.Login(username, passwd)
	userinfo := controller.HandleLogin(rsp)

	rsp = util.GetUserCourses(userinfo)
	AllClass := controller.HandleGetUserCourses(rsp)

	try := 1
	for {
		fmt.Printf("开始自动签到，时间间隔为%v秒\n", config.TimeSleep)
		fmt.Printf("第%v次尝试\n", try)
		for _, class := range AllClass {
			rsp = util.IsCheckInOpen(userinfo, class.ClassId)
			rst, msg := controller.HandleIsCheckInOpen(rsp)
			fmt.Printf("%v：result_msg:%v\n", class.ClassName, msg)
			if rst {
				count:=1
				for {
					//fmt.Printf("%v：获取签到列表\n",class.ClassName)
					//rsp = util.GetCheckInList(userinfo, class.ClassId)
					//checkInId = controller.HandleGetCheckInList(rsp)
					fmt.Printf("%v:开始签到\n", class.ClassName)
					rsp = util.CheckIn(userinfo, class.ClassId)
					rst, msg := controller.HandleCheckIn(rsp)
					fmt.Printf("%v:result_msg:%v\n",class.ClassName,msg)
					if rst {
						fmt.Printf("%v:签到成功！\n",class.ClassName)
						break
					} else if count>10 {
						fmt.Printf("Attention:%v签到失败，请检查网络，重新尝试，否则你可能失去生命！\n",class.ClassName)
						break
					} else {
						fmt.Printf("%v签到失败，第%v次重试\n",class.ClassName,count)
						count++
					}
					time.Sleep(time.Second*5)
				}
			} else {
				fmt.Printf("%v：签到未开始\n",class.ClassName)
			}
		}
		try++
		time.Sleep(config.TimeSleep)
	}
}
