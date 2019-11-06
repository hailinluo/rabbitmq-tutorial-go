## 发布/订阅模式

为了说明这个模式，我们将会构建一个简单的日志系统。它包含两个程序，一个产生日志消息，另一个则接收并打印这些日志。

消息可以被多次消费，比如一个消费者将日志写入磁盘，而另一个则打印日志到屏幕。

#### RabbitMQ中完整的消息传递模型

在此模型中，producer并不知道消息会被发送到哪一个队列中，而是将其发送到exchange中。

Exchange一方面从Producer接收消息，另一方面将消息推送到一个或多个队列。

- direct
- topic
- headers
- fanout：发送消息到所有已知的队列


#### 相关命令

```
// 列出所有的Exchanges
sudo rabbitmqctl list_exchanges
```
