package main

import (
	"fmt"
	"github.com/imroc/req"
	"github.com/tidwall/gjson"
	"os"
	"regexp"
	"time"
)

func main()  {
	userId := os.Getenv("userId")
	password := os.Getenv("pass")
	fmt.Printf("当前账号:%s ,密码:%s\n",userId,password)
	for{
		timeStr:=time.Now().Format("2006-01-02 15:04:05")
		e,_:=req.Get("http://www.baidu.com")
		match, _ := regexp.MatchString("eportal", e.String())
		if match{
			r_href, _ := regexp.Compile("href=.*?\\?(.*?)'")
			queryString:=r_href.FindStringSubmatch( e.String())
			url:="http://192.168.50.3:8080/eportal/InterFace.do?method=login"
			param := req.Param{
				"userId": userId,
				"password":password,
				"queryString":queryString,
				"operatorPwd":"",
				"validcode": "",
			}
			responce,_:=req.Post(url,param)
			if gjson.Get(responce.String(),"result").String() =="fail"{
				fmt.Println( gjson.Get(responce.String(),"message").String())
			}else {
				fmt.Println(timeStr,": 认证成功")
			}
		}else {
			fmt.Println(timeStr,": 已在线")
		}
		time.Sleep(60*10*time.Second)
	}
}
