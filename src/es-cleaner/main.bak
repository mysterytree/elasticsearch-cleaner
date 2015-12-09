package main

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
	"github.com/mattbaird/elastigo/lib"
	"time"
	"flag"
)

const (
	//日期模板
	layout = "2006.01.02"
	//index库的前半段
	indexName = "logstash-momoapi-"
)

var(
    //创建es连接
    esConn = elastigo.NewConn()
    //需要保留的天数，默认10
    remainNumber = flag.Int("n",10,"Remain Number")
    //开始删除的时间
	startDay = flag.String("t", "", "Start Deleting Date")
)

func main() {
	//控制台接收参数
    flag.Parse()
	//删除参数日期到十天前的所有日志
    if *startDay!=""{
    	deletePreviousIndex()
    }
	//计划任务
	gocron.Every(1).Day().At("02:00").Do(autoDelete)
	gocron.Start()
}

//删除10天前的所有日志
func deletePreviousIndex() {
	expiredate, err := time.Parse(layout, *startDay)
	if err != nil {
		fmt.Println(err)
	} else {
		limitdate := time.Now().AddDate(0, 0, (1 - *remainNumber))
		expireduration := (int)(expiredate.Sub(limitdate).Minutes() / 60 / 24)
		for expireduration < 0 {
			deleteIndex(indexName+limitdate.AddDate(0, 0, expireduration).Format(layout))
			expireduration++
		}
	}
}

//计划任务，自动删除
func autoDelete() {
	deleteIndex(indexName+time.Now().AddDate(0, 0, 1 - *remainNumber).Format(layout))
}

//删除某天日志
func deleteIndex(index string) {
	_, err := esConn.DeleteIndex(indexName + index)

	if err != nil {
		fmt.Printf("del index error: %v \n", err)
	} else {
		fmt.Println("**" + index + "**delete**")
	}
}
