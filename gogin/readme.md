## 1.初始化
```
go mod init gogin
```

## 2.依赖 golang
```
docker pull golang
```

## 3.编译docker image

```
docker build -t gogin:v1.0.0 .
```

## 4.运行
```
docker run -d -p 8080:8080 gogin:v1.0.0
```




