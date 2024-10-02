# 进行扫描redis key 和模糊匹配key

# 使用操作：
```shell
.\redisscankey.exe -h
{"level":"info","msg":"初始化脚本","time":"2024-10-02 08:05:56"}
Usage of D:\Go\src\redisscankey\redisscankey.exe:
  -ip string
        当前的redis 地址（如果是集群，用逗号分隔多个地址） (default "127.0.0.1:6307")
  -key string
        当前的redis key (default "key*")
  -password string
        当前的redis 密码 (default "password")
---
redisscankey.exe  --ip="62.234.1.167:31634" --key="xi*" --password="lush222123"
{"level":"info","msg":"初始化脚本","time":"2024-10-02 08:09:01"}
{"level":"info","msg":"初始完成","time":"2024-10-02 08:09:01"}
{"level":"info","msg":"开始进行连接redis","time":"2024-10-02 08:09:01"}
{"level":"info","msg":"redis连接成功","time":"2024-10-02 08:09:01"}
{"level":"info","msg":"开始执行","time":"2024-10-02 08:09:01"}
{"level":"info","msg":"当前节点：62.234.1.167:31634 的key数：2","time":"2024-10-02 08:09:01"}      
{"level":"info","msg":"当前节点：62.234.1.167:31634 模糊匹配key数：2","time":"2024-10-02 08:09:01"}
{"level":"info","msg":"当前redis一共有2","time":"2024-10-02 08:09:01"}
{"level":"info","msg":"当前redis一共有2","time":"2024-10-02 08:09:01"}
{"level":"info","msg":"执行完成","time":"2024-10-02 08:09:01"}
```
