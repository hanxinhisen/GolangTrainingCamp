
###1.使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。

####环境准备

```
docker run -d --name redis-server -p6388:6379  redis:alpine
```

####分别测试
```
redis-benchmark -h 192.168.1.186 -p 6388 -d 10 -t get,set
redis-benchmark -h 192.168.1.186 -p 6388 -d 20 -t get,set
redis-benchmark -h 192.168.1.186 -p 6388 -d 50 -t get,set
redis-benchmark -h 192.168.1.186 -p 6388 -d 100 -t get,set
redis-benchmark -h 192.168.1.186 -p 6388 -d 200 -t get,set
redis-benchmark -h 192.168.1.186 -p 6388 -d 1000 -t get,set
redis-benchmark -h 192.168.1.186 -p 6388 -d 5000 -t get,set
redis-benchmark -h 192.168.1.186 -p 6388 -d 10000 -t get,set
```


####结果汇总

| size  | get(QPS) | set(QPS) |
|:-----:|:--------:|:--------:|
|  10   | 60753.34 | 55834.73 |
|  20   | 59665.87 | 61087.36 |
|  50   | 61349.70 | 58892.82 |
|  100  | 55897.15 | 57703.40 |
|  200  | 61462.82 | 61124.69 |
| 1000  | 58685.45 | 65274.15 |
| 5000  | 54288.82 | 53304.90 |
| 10000 | 51813.47 | 45330.91 |

### 2.写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息 , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。

####执行脚本
```go run main.go```
####结果汇总
| 大小    | 插入前    | 插入后        | 变化         | 平均          | 取整    |
|-------|--------|------------|------------|-------------|-------|
| 10    | 68384  | 8047872    | 7979488    | 39.89744    | 40    |
| 10    | 68384  | 8047872    | 7979488    | 39.89744    | 40    |
| 20    | 48064  | 9648048    | 9599984    | 47.99992    | 48    |
| 50    | 48216  | 16048200   | 15999984   | 79.99992    | 80    |
| 100   | 48368  | 27248352   | 27199984   | 135.99992   | 136   |
| 200   | 48520  | 49648504   | 49599984   | 247.99992   | 248   |
| 1000  | 48672  | 209648656  | 209599984  | 1047.99992  | 1048  |
| 5000  | 48824  | 1028848808 | 1028799984 | 5143.99992  | 5144  |
| 10000 | 48976  | 2052848960 | 2052799984 | 10263.99992 | 10264 |