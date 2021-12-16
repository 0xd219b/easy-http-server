# easy-http-server
快速创建一个http server服务

## Usage
### 单一模式

easy-http-server -p [server_port] single -u [url_path] -m [request_method]

`server_port`: 服务的端口号

`url_path`: 服务的path

`request_method`: 请求类型

本地启动一个监听在`9080`端口，请求类型为`POST`，`path`是`/info`的服务
```
./easy-http-server -p 9080 single -u info -m POST
```