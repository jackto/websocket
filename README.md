#概要
用golang尝试做一个websocket的服务端。保持长链接，适合于ios/android的消息推送，聊天室/游戏的实时消息活动


## 主要设计思路

- 维持websocket连接池
- 维持一个消息队列，消息队列的消息都会推送到所有的连接客户端
- 用协程模拟客户端请求，从而调戏服务端