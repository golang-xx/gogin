## 1.初始化
```
go mod init ginweb01
```

## 2.依赖 golang
```
docker pull golang
```

## 3.编译docker image

```
docker build -t ginweb01:v1.0.0 .
```

## 4.运行
```
docker run -d -p 8091:8091 ginweb01:v1.0.0
```




