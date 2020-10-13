## go-zero 练习

#### bookstore/api目录下通过goctl生成api/bookstore.api
```shell
    mkdir api
    goctl api -o bookstore.api
```

###### 使用goctl生成API Gateway代码

```
goctl api go -api bookstore.api -dir .
```

#### 启动gateway 服务

```
go run bookstore.go -f etc/bookstore-api.yaml
```

##### 测试服务是否可用
```
    curl -i "http://localhost:8888/check?book=go-zero"
```
##### 返回的结果 

```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sat, 10 Oct 2020 09:22:09 GMT
Content-Length: 25

{"found":false,"price":0}%
```

##### 可以通过命令生成proto文件模板
```
goctl rpc template -o add.proto
```

##### 用goctl生成rpc代码，在rpc/add目录下执行命令
```
goctl rpc proto -src add.proto
```

### 编写check rpc服务
```
goctl rpc template -o check.proto
```

### 在rpc/check目录下执行命令

```
goctl rpc proto -src check.proto
```

### 启动服务

```
go run check.go -f etc/check.yaml
```

### 定义数据库表结构，并生成CRUD+cache代码
###### 在rpc/model目录下执行如下命令生成CRUD+cache代码，-c表示使用redis cache
```
goctl model mysql ddl -c -src book.sql -dir .
```


###  测试
###### 添加书箱
```
curl -i "http://localhost:8888/add?book=go-zero&price=10"
```
##### 查询书箱
```
curl -i "http://localhost:8888/check?book=go-zero"
```