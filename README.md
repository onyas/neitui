# 内推我

后端用golang抓取多个网站的内推信息，可能对找工作的你有所帮助

#### 本地运行如何运行
```
go build -o bin/neitui -v .
heroku local
```

#### 查看heroku已有的变量配置
```shell script
heroku addons:create heroku-postgresql:hobby-dev
heroku config
```
#### 发布
```shell script
git push heroku master
```
