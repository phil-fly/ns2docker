

### 获取docker容器列表

### 通过namespace查询容器信息
```go
    dockerContainer := ns2docker.Search("4026532229")
    log.Infof("使用namespace: %v 查询到容器: %v","4026532229",dockerContainer)
```

### 通过容器ID查询namespace
```go
    if namespace,err := ns2docker.QueryNs("bf5d97f08b5a");err == nil {
        log.Infof("使用容器ID：%s 查询到namespace: %v","bf5d97f08b5a",namespace)
    }else{
        log.Error(err.Error())
    }
```