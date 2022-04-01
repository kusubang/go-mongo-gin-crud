# Go + Gin + MongoDB

Creating a CRUD app using Go and Gin and MongoDB

## 몽고DB 클라이언트 using docker
몽고DB 가 `mongodb`로 실행 중인경우 다음과 같이 한다.
```
# docker exec -it mongodb bash

```
컨테이너로 접속하고 나면 다음과 같이 한다.
```
# mongosh "mongodb://1234:1234@127.0.0.1:27017
```

## reference
- [Creating a CRUD application using GO and MongoDB](https://medium.com/@kumar16.pawan/creating-a-crud-application-using-go-and-mongodb-cc077ce2d0e)

- [How to use MongoDB with Go](https://blog.logrocket.com/how-to-use-mongodb-with-go/)
- [Using the MongoDB Shell](https://www.mongodb.com/basics/examples)
- [MongoDB Examples With Golang](https://blog.ruanbekker.com/blog/2019/04/17/mongodb-examples-with-golang/)
