# Quick Image Server

## 简介
Quick Image Server是一个简单易用的图像服务器，目前支持JPEG格式。

## 配置文件

```JSON
{
	"ListenAddr":":8086",      <-监听地址
	"Storage":"/var/www/html/image/storage/"      <-存储位置
}
```

## 使用方法
### 上传图片
POST /
表单参数:
uploadfile file类型,要上传的图片
返回值:
{图片ID}

### 下载图片
GET /{图片ID}
返回值:
{图片文件}