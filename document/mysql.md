### 数据库

- 关联关系
  - 数据库是有状态的服务，横向扩展非常不好，应尽可能减轻数据库的压力
  - 按照DDD的理论，关联关系应该在 Repository 层面处理，而不是 dao 层
  
- 开启 sql 语句的日志

  > db = db.Debug()

- 软删除

  > 业界通用做法
  >
  > 同时，软删除并没有什么用

- 使用 Docker Compose 启动数据库

- 

### GORM

- 声明模型：https://gorm.io/zh_CN/docs/models.html

### Docker

- 使用 Docker Compose 启动数据库

- Docker Compose 语法

  > https://hub.docker.com/_/mysql

- 

