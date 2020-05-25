
rabbitMQ是一个在AMQP协议标准基础上完整的，可服用的企业消息系统。它遵循Mozilla Public License开源协议，采用 Erlang 实现的工业级的消息队列(MQ)服务器，Rabbit MQ 是建立在Erlang OTP平台上。

1、安装Erlang
下载地址：https://www.erlang.org/downloads，
设置环境变量
修改环境变量path，增加Erlang变量至path，%ERLANG_HOME%\bin;
打开cmd命令框，输入erl

2、安装rabbitmq
下载地址：http://www.rabbitmq.com/download.html
修改环境变量path，增加rabbitmq变量至path，%RABBITMQ_SERVER%\sbin;
打开cmd命令框，切换\sbin目录下，输入rabbitmqctl status

    rabbitmq-plugins.bat enable rabbitmq_management

安装完成浏览
```html
http://localhost:15672/
```


Go加载
```html
go get github.com/streadway/amqp
```



基本使用打开连接请求

=======================================
```html


conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
failOnError(err, "Failed to connect to RabbitMQ")
defer conn.Close()

ch, err := conn.Channel()
failOnError(err,"Failed to open a channel")
defer ch.Close()

q,err := ch.QueueDeclare(
   "queue",
   true,
   false,
   false,
   false,
   nil,
   )
failOnError(err,"Failed To declare a queue")
============================================
//1，存值方法
body := `HELLOW WORLDASdfes`
err = ch.Publish("",q.Name,false,false,amqp.Publishing{
   ContentType:     "texp/plain",
   Body:            []byte(body),
})
 
//2 ，取值
msgs, err := ch.Consume(
   q.Name, // queue
   "",     // consumer
   true,   // auto-ack
   false,  // exclusive
   false,  // no-local
   false,  // no-wait
   nil,    // args
)
 
forever := make(chan bool)
go func() {
   for d := range msgs {
      log.Printf("Received a message: %s", d.Body)
   }
}()
log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
<-forever

*******************************************
task_work设置
//1，存值、
err = ch.Publish(
   "",
   q.Name,
   false,
   false,
   amqp.Publishing{
      ContentType:     "text/plain",
      DeliveryMode:    amqp.Persistent,
      Body:            []byte(body),
   },
   )

//2，取值设置qos 调用Consume 中函数 Ack(false)
err = ch.Qos(
   1,
   0,
   false,
)
msgs, err := ch.Consume(
   q.Name, // queue
   "",     // consumer
   false,  // auto-ack
   false,  // exclusive
   false,  // no-local
   false,  // no-wait
   nil,    // args
)
 
forever:= make(chan bool)
go func() {
   for d := range msgs {
      log.Printf("Received a message: %s", d.Body)
   
      d.Ack(false)
   }
}()

 
```