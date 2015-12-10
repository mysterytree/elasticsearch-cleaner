# About escleaner

[![build status](https://gitlab.baozou.com/ci/projects/17/status.png?ref=master)](https://gitlab.baozou.com/ci/projects/17?ref=master)

## 简介
* 自动删除elasticsearch旧日志的工具

## 主要功能
* 每晚自动删除x天前日志
* 删除某个日期到x天前的所有日志



## 安装环境

###  docker:

```
Client:
 Version:      1.9.1
 API version:  1.21
 Go version:   go1.4.2
 Git commit:   a34a1d5
 Built:        Fri Nov 20 13:12:04 UTC 2015
 OS/Arch:      linux/amd64

Server:
 Version:      1.9.1
 API version:  1.21
 Go version:   go1.4.2
 Git commit:   a34a1d5
 Built:        Fri Nov 20 13:12:04 UTC 2015
 OS/Arch:      linux/amd64

```
 
## 安装方法

### 1.编辑Dockerfile，根据实际情况修改CMD
* ``-i``  后跟需删除的index前缀，如``logstash-momoapi-`` 多个index以``，``分割
* ``-n``后跟保留多少天的日志
* ``-t``  后跟需删除的日期，如2015.11.02 ，将删除2015.11.02 到n天前的所有日志
* ``-h``es的domain
* ``－p``es端口

如,删除前缀``logstash-momoapi-``和``logstash-ribao-``的2014.11.12到5天前的所有日志,并且每天定时删除5天前的日志


```
CMD /go/bin/es-cleaner －i logstash-momoapi-,logstash-ribao-  -n 5   -t 2014.11.12
```

只执行每天定时删除,不添加``－t``即可

```
CMD /go/bin/es-cleaner -i logstash-momoapi-,logstash-ribao-  -n 10  -h 192.168.1.2 -p 9200
```
### 2.打包镜像

执行：

```
docker build -t dockernj.baozou.com/nanjing/escleaner .
```
### 3.上传
```
docker login

docker push dockernj.baozou.com/nanjing/escleaner
```
### 4.运行
```
docker run －d --rm dockernj.baozou.com/nanjing/escleaner
```