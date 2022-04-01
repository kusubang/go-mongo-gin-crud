# Go + Gin + MongoDB

Creating a CRUD app using Go and Gin and MongoDB

## 몽고DB 클라이언트 using docker
몽고DB 가 `mongodb`로 실행 중인경우 다음과 같이 한다.
```
host$ docker exec -it mongodb bash
```
컨테이너로 접속하고 나면 다음과 같이 한다.
```
# mongosh "mongodb://1234:1234@127.0.0.1:27017
```

## 고찰
Controller 구현할 때 몽고DB 클라이언트를 사용하는데 있어서
1. controller 내부에 global config 를 두는게 좋을지
2. 외부 주입을 통해서 사용하는 방법이 좋을지 

```
var collection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var validate = validator.New()

func CreateUser() gin.HandlerFunc {
    return func(c *gin.Context) {
      collection.Insert()
```
1. configs 패키지로부터 collection 을 얻어오는 코드가 controller 에 있다.
2. CreateUser 함수는 gin.HandlerFunc 를 반환하고 있다.

```
type UserHandler struct {
	collection *mongo.Collection
}

func NewUserHandler(collection *mongo.Collection) *UserHandler {
	return &UserHandler{collection}
}

func (h *UserHandler) CreateAUser(c *gin.Context) {
  h.collection.Insert()
```
UserHandler 구조체를 정의하고 팩토리 함수를 제공하여 외부에서 collection 을 주입하고 있다.

## Reference
- [Creating a CRUD application using GO and MongoDB](https://medium.com/@kumar16.pawan/creating-a-crud-application-using-go-and-mongodb-cc077ce2d0e)
- [How to use MongoDB with Go](https://blog.logrocket.com/how-to-use-mongodb-with-go/)
- [Using the MongoDB Shell](https://www.mongodb.com/basics/examples)
- [MongoDB Examples With Golang](https://blog.ruanbekker.com/blog/2019/04/17/mongodb-examples-with-golang/)
- [Build a REST API with Golang and MongoDB - Gin-gonic Version](https://dev.to/hackmamba/build-a-rest-api-with-golang-and-mongodb-gin-gonic-version-269m)
