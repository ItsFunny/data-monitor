## 流量监控服务

- 通过zmq 从capture中接收消息
    -   为什么:因为流量通常都是很大的,为了防止崩溃,所以通过zmq削锋