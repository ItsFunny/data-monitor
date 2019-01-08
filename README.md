## 预期的技术栈
* Go做底层的抓包,分析,数据库采用influxdb
* Java rest请求Go服务器获取数据
* Java Freemarker 页面展示
* Java 做微服务化(SpringCloud),Go也做微服务化(后期)

## 流量监控
## 预期的功能
* 异常流量统计
* 正常流量统计
* 流量峰值
## 服务拆分
* 流量抓取服务
* 流量分析服务
* 流量存储服务
## 服务间通讯
* protobuf
* zmq