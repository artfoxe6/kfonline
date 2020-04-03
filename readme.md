
##### 如何使用

1. 下载本项目，编辑config.ini文件，至少需要配置 MySQL,Redis,Centrifugo,Jwt,
2. 本项目依赖Centrifugo服务，请下载并运行Centrifugo服务，服务配置文件请参考
项目中的centrifugo_server.json
3. go run main.go

##### 功能

- 基本客服接待
- 客户备注
- 聊天内容持久化
- 可扩展

##### 待优化的部分

- 将centrifugo集成到项目中

##### 本项目使用了

- [gin](https://github.com/gin-gonic/gin)
- [quick_gin](https://github.com/codeAB/quick_gin)
- [centrifugo](https://github.com/centrifugal/centrifugo)
    

##### 架构梳理

>客服登录系统后，开启一个自己的频道
>
>用户进入网页后，调用接口查找在线的客服，如果有可用客服，就调用接口获取授权凭证jwt
>
>用户订阅客服的频道并自己创建一个频道，用户发送消息到客服频道，客服回消息到用户自己的频道
>
>用户和客服只允许通过接口代理发送消息到频道中，系统做持久化处理

    
    
    
    