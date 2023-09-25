





## 数据库设计
TABLE `user` 
`id` bigint(20) NOT NULL AUTO_INCREMENT,
`user_id` bigint(20) NOT NULL,
使用两个id,防止泄露网站用户数量,分库分表时user_id可能重复，....

分布式ID生成器
唯一性 不重复
递增   ID递增
高可用  任何时候都能生成正确的ID
高性能 高并发环境下可用

