# go-latency


## Introduction
Just a simple latency test for Go against a remote Database.

It's a simple test to measure the latency of a remote database. I wanted it to be simple and easy to use and to be able to put the queries in a file verses hardcoding them in the code.


```BASH
go get github.com/go-latency/latency
```

```BASH
go-latency -h                       
  -e string
        endpoint to test latency to and from
  -f string
        sql file to run
exit status 1

go-latency -e 10.x.x.x -f test.sql
ID :  30
USER :  root
HOST :  172.17.0.1:59542
DB :  NULL
COMMAND :  Query
TIME :  0
STATE :  executing
INFO :  SELECT ID,USER,HOST,DB,COMMAND,TIME,STATE,INFO FROM information_schema.processlist WHERE command !='Sleep'
--------------------------------------------------------------------------------------------------------------------
Latency:  1.782614ms
```