# GoGoGo


## 第三方库管理－godep使用

```
go get github.com/tools/godep # 下载工具

godep save # 每次添加并使用第三方库后使用

godep save ./... # 抓取应用程序中所有的第三方库依赖

godep update foo/bar # 更新某个库

godep restore # 恢复到Godeps中指定的版本

visit https://github.com/tools/godep # more
```