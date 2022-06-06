# simple-demo

## 抖音项目服务端简单示例

具体功能内容参考飞书说明文档

工程无其他依赖，直接编译运行即可

```shell
go build && ./simple-demo
```

### 配置文件

```json
{
  "listen_port": 8080, //http监听端口
  "base_url": "http://10.0.2.2:8080/static/", //静态资源访问地址
  "use_cos": true, //是否使用cos
  "cos":{
    "secret_id":"AKIDDPIwPpesUoqRwxlTkkeBkmsqTwH2rAn3", //cos secret_id
    "secret_key":"8Dk6jFZP1nWL1B7NbZkpb5KNcP9LQDfq", //cos secret_key
    "address":"https://douyin-1258365609.cos.ap-shanghai.myqcloud.com" //cos地址
  },
  "mysql":{
    "host":"localhost", //mysql地址
    "port":3306, //mysql端口
    "user":"dy", //mysql用户名
    "password":"douyin", //mysql密码
    "database":"simple_demo" //mysql数据库名
  }
}
```

### 功能说明
基本完成所有功能，使用cos时自动获取视频封面