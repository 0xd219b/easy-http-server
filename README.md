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

### Mock模式

./easy-http-server mock -y [yaml_file_path]

`yaml_file_path`: yaml文件的路径

```yaml
# yaml文件
- url: /test
  method: POST
  resp: "{\"test\":\"created\"}"
- url: /test/file
  method: GET
  resp: "{\"test\":\"ok\"}"
```

## features
- [x] 支持单一模式
- [x] 支持mock模式
- [x] 支持yaml文件配置
- [ ] 支持多种请求类型（目前仅支持POST、GET）
- [ ] 支持多种响应类型（目前仅支持json）
- [ ] 支持响应Header