# RediStructHash
A Go File Which Can Automatically Convert Struct to Redis Hash

## About

The traditional way to create a Redis hash from Go struct is to hardcode the fields one by one, both tedious and prone to error. So I wrote this mini ORM to automate this process.

Since the code is really short, you can make any customization to meet your own request.

### Redis Package

[go-redis/redis](https://github.com/go-redis/redis)

## Usage

### Tag

`redistructhash:"no"`: tell function not to create hash for this field

### Naming

The field name in struct should be either CamelCase or camelCase as a convention, and the corresponding name of Redis hash field will be converted to camel_case. However, any form of name is supported.

## Example

```go
type Video struct {
	Id         int64
	PlayUrl    string
	Title      string `redistructhash:"no"` // use "no" tag to prevent creating field
	CreateTime int64
}

func main() {
	var rdb *redis.Client
	var ctx = context.Background()
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	video := Video{
		Id:         1,
		PlayUrl:    "youtube.com",
		Title:      "my video",
		CreateTime: time.Now().Unix(),
	}
	if err := RedisStructHash(rdb, ctx, video, "key"); err != nil {
		log.Println(err)
	}
	fmt.Println(rdb.HGetAll(ctx, "key").Result())
}
```

output:

```
map[create_time:1655903430 id:1 play_url:youtube.com] <nil>
```