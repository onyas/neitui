# neitui
this is the API for neitui

```
go run main.go

or

go build -o bin/neitui -v .
heroku local
```

heroku config:set DSNAME=postgres://postgres:root@127.0.0.1:5432/chat?sslmode=disable

list all the config
```shell script
heroku addons:create heroku-postgresql:hobby-dev
heroku config
```
