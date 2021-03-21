## 1.初始化
```
go mod init ginweb
```

## 2.依赖 golang
```
docker pull golang
```

## 3.编译docker image

```
docker build -t ginweb:v1.0.0 .
```

## 4.运行
```
docker run -d -p 8090:8090 ginweb:v1.0.0
```




