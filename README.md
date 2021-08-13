## 收图服务



### 简介

---

一个极其简单的收图服务



### 编译版部署

---

#### 下载编译好的程序

```shell
wget https://github.com/PantsuDango/receive_img/releases/download/Ver1.0/receive_img_server
```

#### 赋予权限

```shell
chmod +x receive_img_server
```

#### 运行

```shell
./receive_img_server [端口号] > log.INFO 2>&1 &
```

#### 例如

```shell
./receive_img_server 23333 > log.INFO 2>&1 &
```



### 源码部署

---

```shell
git clone https://github.com/PantsuDango/receive_img
```

#### 下载依赖

```go
go mod tidy
```

#### 编译

```go
go build app.go
```

#### 运行

```shell
nohup go run app.go [端口号] > log.INFO 2>&1 &
```

#### 例如

```
nohup go run app.go 23333 > log.INFO 2>&1 &
```



### 图片保存目录

---

├── images
│   ├── ENG
│   ├── JAP
│   └── KOR



### 接口说明

---

##### 请求URL
- ` http://你的IP地址:端口号/receive_img `
##### 请求方式
- POST 

##### 请求示例

```json
{
    "Language": "JAP",
    "ImageBase64": "/9j/4AAQSkZJRgABAQEAYABgAAD/2wBDAAgGBgcGBQgHBwcJCQgKDBQNDAsLDBkSEw8UHRofHh0aHBwgJC4nICIsIxwcKDcpLDAxNDQ0Hyc5PTgyPC4zNDL/2wBDAQkJCQwLDBgNDRgyIRwhMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjL/wAARCAAbACgDASIAAhEBAxEB/8QAHwAAAQUBAQEBAQEAAAAAAAAAAAECAwQFBgcICQoL/8QAtRAAAgEDAwIEAwUFBAQAAAF9AQIDAAQRBRIhMUEGE1FhByJxFDKBkaEII0KxwRVS0fAkM2JyggkKFhcYGRolJicoKSo0NTY3ODk6Q0RFRkdISUpTVFVWV1hZWmNkZWZnaGlqc3R1dnd4eXqDhIWGh4iJipKTlJWWl5iZmqKjpKWmp6ipqrKztLW2t7i5usLDxMXGx8jJytLT1NXW19jZ2uHi4+Tl5ufo6erx8vP09fb3+Pn6/8QAHwEAAwEBAQEBAQEBAQAAAAAAAAECAwQFBgcICQoL/8QAtREAAgECBAQDBAcFBAQAAQJ3AAECAxEEBSExBhJBUQdhcRMiMoEIFEKRobHBCSMzUvAVYnLRChYkNOEl8RcYGRomJygpKjU2Nzg5OkNERUZHSElKU1RVVldYWVpjZGVmZ2hpanN0dXZ3eHl6goOEhYaHiImKkpOUlZaXmJmaoqOkpaanqKmqsrO0tba3uLm6wsPExcbHyMnK0tPU1dbX2Nna4uPk5ebn6Onq8vP09fb3+Pn6/9oADAMBAAIRAxEAPwDwal2k9BmkzXR+CLXxJf8AiD7J4Xup7W9liIlnimMQjiBBYu46KCB+lQal7xR8PdS8PWlnfwR3N7p02nwXkt2LYpHC0hI8stkgkfLzx94cCsvVvB2v6HrFrpeoac8N1dsq2wLqUm3EAbXB2nkjvxnnFep/FG40IXOh2PiDX9cnb+yLdimnhZreQ5Yedl3G4tg84yQBzXF+NbXVNN8J+HreLXm1XwvdB7jTy8QRonHDowySCu7GNxHXFNolMn0z4SavfaJqVzcTRWeo2wVra0lkiK3I53AOH+VhjuMdOeeCvVtF8N6ZLqXg+L+wYXsrjRo5brHh+GWCSQxMd0lyfmRsgHGDk4/vUU7CufM1dR4K8T2vhyfVIr+1mnstUsZLGc27hJUV8ZZCQRnjoeK5eipL6Hod78RtLvNfaW48L217oy6ZHpcVrcuBMkUedrrKF+R8k/dHp6ZrF8V+L4vEGnaTpNhpSaXpOlo4gtxMZnLOcuzOQM5I9O5rlqKLisj1ux+J2g23ijw/rEtvqrLpGiR2CwqECvMFdWP3vu4YYPXjpRXklFFw5Uf/2Q=="
}
```

##### 请求参数说明 

| 参数名      | 是否必选 | 类型   | 说明                                                         |
| ----------- | -------- | ------ | ------------------------------------------------------------ |
| Language    | 是       | String | 语种，可选值包括：<br/>- ENG：英文<br/>\- JAP：日语<br/>- KOR：韩语 |
| ImageBase64 | 是       | String | 图像数据，base64编码                                         |


##### 返回示例 

``` json
{
    "Code": 0,
    "RequestID": "b0171c0c-029f-4ad4-bd04-1412db56b4d0",
    "Result": "Success",
    "Status": "Success"
}
```
