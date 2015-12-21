package main

import (
	"flag"
	"fmt"
	"github.com/jasonlvhit/gocron"
	"github.com/mattbaird/elastigo/lib"
	"strings"
	"time"
)

const (
	//日期模板
	layout = "2006.01.02"
)

var (
	//创建es连接
	esConn *elastigo.Conn
	//需要保留的天数，默认10
	remainNumber = flag.Int("n", 10, "Remain Number")
	//开始删除的时间
	startDay = flag.String("t", "", "Start Deleting Date")
	//需要删除的indexname前半段
	indexPtr = flag.String("i", "logstash-momoapi-", "index")
	hostPtr  = flag.String("h", "localhost", "host")
	portPtr  = flag.String("p", "9200", "port")

	//接收indexName数组
	indexNames []string
)

func main() {
	//控制台接收参数
	flag.Parse()
	esConn = myConn()
	indexNames = strings.Split(*indexPtr, ",")
	for i := 0; i < len(indexNames); i++ {
		//删除参数日期到十天前的所有日志
		if *startDay != "" {
			deletePreviousIndex(indexNames[i])
		}
	}
	//计划任务
	gocron.Every(1).Day().At("02:00").Do(autoDelete)
	<-gocron.Start()
}

//删除index天前的所有日志
func deletePreviousIndex(indexName string) {
	expiredate, err := time.Parse(layout, *startDay)
	if err != nil {
		fmt.Println(err)
	} else {
		limitdate := time.Now().AddDate(0, 0, (1 - *remainNumber))
		expireduration := (int)(expiredate.Sub(limitdate).Minutes() / 60 / 24)
		for expireduration < 0 {
			deleteIndex(indexName + limitdate.AddDate(0, 0, expireduration).Format(layout))
			expireduration++
		}
	}
}

//计划任务，自动删除
func autoDelete() {
	for i := 0; i < len(indexNames); i++ {
		//遍历删除
		deleteIndex(indexNames[i] + time.Now().AddDate(0, 0, 0-*remainNumber).Format(layout))
	}
}

//删除某天日志
func deleteIndex(index string) {
	_, err := esConn.DeleteIndex(index)

	if err != nil {
		fmt.Printf("del index error: %v %s \n", err, index)
	} else {
		fmt.Println("**" + index + "**delete**")
	}
}

//新建conn
func myConn() *elastigo.Conn {
	return &elastigo.Conn{
		// Maintain these for backwards compatibility
		Protocol:       "http",
		Domain:         *hostPtr,
		ClusterDomains: []string{"localhost"},
		Port:           *portPtr,
		DecayDuration:  time.Duration(0 * time.Second),
	}
}
