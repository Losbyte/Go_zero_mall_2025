# Go_zero_mall_2025
Detailed Description about how to deploy this classic project with the newest version of Go-zero

## Environment Configuration

此项目中使用了Etcd\Mysql\Redis\Mysql-manage\redis-manage\prometheus\grafana\jaeger\DTM

```bash
docker pull xxx 
```

| DTM          | yedf/dtm                      | HTTP端口号为36789，gRPC端口号为36790 |
| ------------ | ----------------------------- | ------------------------------------ |
| ETCD         | quay.io/coreos/etcd:v3.5.12   | 2379                                 |
| Mysql        | mysql:5.7                     | 3306,admin,123456,123456             |
| Redis        | redis:5.0                     | 6379                                 |
| Mysql Manage | phpmyadmin/phpmyadmin         | admin,123456,123456,mysql,3306,1000  |
| Redis Manage | erikdubbelboer/phpredisadmin  | admin,123456,redis,6379,2000         |
| Prometheus   | bitnami/prometheus            | 3000                                 |
| Grafana      | grafana/grafana               | 4000                                 |
| Jaeger       | jaegertracing/all-in-one:1.28 | 5000                                 |

上述镜像拉取到docker desktop后，可以直接利用给出的docker-compose.yml 进行Golang环境部署，后续此文件也是启动容器的起点。注意DTM的布置，在使用前需要人为在Mysql中添加一张表。

此处使用的Golang的版本为1.24.7，Go-zero的版本为1.9





