package main

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
	"github.com/mattbaird/elastigo/lib"
	"time"
)

const (
	//需要保留的天数
	remainNumber = 1
	//日期模板
	layout = "2006.01.02"
	//index库的前半段
	indexname = "logstash-momoapi-"
)

//创建es链接
var esconn = elastigo.NewConn()

func main() {
	//删除参数日期到十天前的所有日志

	startDay = flag.String("s", "10", "days before")

	// deleteOldIndex(startDay)

	//计划任务
	gocron.Every(1).Day().At("02:00").Do(autoDelete)
	gocron.Start()
}

//删除10天前的所有日志
func deleteoldindex(t string) {
	expiredate, err := time.Parse(layout, t)
	if err != nil {
		fmt.Println(err)
	} else {
		limitdate := time.Now().AddDate(0, 0, (1 - remainNumber))
		expireduration := (int)(expiredate.Sub(limitdate).Minutes() / 60 / 24)
		for expireduration < 0 {
			deleteIndex(limitdate.AddDate(0, 0, expireduration).Format(layout))
			expireduration++
		}
	}
}

//计划任务，自动删除
func autoDelete() {
	deleteIndex(time.Now().AddDate(0, 0, -10).Format(layout))
}

//删除某天日志
func deleteIndex(index string) {
	_, err := esconn.DeleteIndex(indexname + index)

	if err != nil {
		fmt.Printf("del index error: %v \n", err)
	} else {
		fmt.Println("**" + index + "**delete**")
	}
}
