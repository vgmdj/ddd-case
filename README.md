# ddd-case

ddd domo，遵循ddd设计，代码结构严格分层，不允许跨层引用

## 结构

文件结构如下

- app 应用层
  - service 应用服务
  - repository 文件缓存相关仓储定义
- domain 领域层
  - entity 领域对象
    - vo 值对象
  - event 领域事件
  - repository db仓储接口定义
    - facade 数据库相关接口定义
  - service 领域服务
- infrastructure 基础层
  - config 配置相关
  - repository 仓储实现
    - ent 数据库操作组件及po
    - persistence 实现domain里定义的仓储接口
- interfaces 接口层
  - assembler 组装器
  - dto 数据传输对象
  - facade 调用接口
- config 配置文件
- Makefile 编译运行脚本
- Dockerfile 容器打包指令
- main.go main函数

## 举例

一个新接口运行的顺序应该为：

- interfaces.httpfacade 接收到的消息 interfaces.dto
- interfaces.assembler 把消息结构处理成 domain.entity 即 do
- app.service 拿到 do 后，调用一个或多个 domain.service 处理
- domain.repository.facade 定义处理接口
- infrasturcture.repository.persistence实现了domain.repository.facade里定义的接口
- infrastructure.repository.persistence 用 infrasturcture.repository.ent 实现的 dao 将 do 转化成 infrastructure.repository.ent里定义的po ，并持久化
- domain.service 将 infrastructure.repository.persistence 实现的 domain.repository.facade 接口结果返回

![图示](https://static001.geekbang.org/resource/image/2d/1d/2d6a328a9fd8b4b3906bb9f59435ca1d.png)

## 注意

- 领域内不可以调用外部其他服务
- 如果需要调用外部服务，需要移到应用层中实现
- 应用层不应包含任何逻辑处理
- 所有逻辑处理应该移到领域服务中实现
