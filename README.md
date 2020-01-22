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

#### 新增变量
```shell script
heroku config:set HTTP_AUTH=123
```

#### 发布
```shell script
git push heroku master
```

#### 定时任务
在这里配置并触发 https://cron-job.org/
