# 售票服务重构
## 前置任务
1. 演出计划以及对应的演出时间，放入redis中，方便现在买票时判断是否在售票时间内。
   2. 实现演出计划的实时上映，表结构添加上映字段，表示是否上映。
## 流程调整
   - redis做只读缓存，redis票默认不进行过期。
   - redis中票过期时机
     - 演出计划要上映了，此时让票过期/删除。
   - 买票场景下，多用户抢票处理流程
       1. 查看redis票状态，判断是否能进行购票
       2. 使用redis分布式锁限制到一个用户请求，不更新redis，直接修改DB内容。
       3. 用MySQL事务保证DB执行正确。
   - 通过canal将DB票记录的变更同步到redis中，满足最终一致性。

# 订单服务重构
- 原消息队列
  - 订单服务与售票服务之间通信，用持久化保证订单服务一定能收到消息
  - 订单服务收到消息之后，将消息放入重试队列，并回复ack
- 重试队列
  - 收到消息之后，业务读取消息进行消费，如果消费成功，则回复ack
  - 如果消费失败，消息继续保留在消息队列中，等待下次消费