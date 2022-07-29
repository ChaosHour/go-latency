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

The latency is the time it takes to execute the query.

What inspired me to create this tool and it's still a work in progress. The way you can use BASH and the mysql client to test the latency.

```BASH
date +"%T" && mysql -vv -h 10.x.x.x -e "SELECT 1" mysql && date +"%T"
11:27:09
--------------
SELECT 1
--------------

+---+
| 1 |
+---+
| 1 |
+---+
1 row in set (0.00 sec)

Bye
11:27:09
```
