package middleware



import (
	"github.com/go-redis/redis/v8"
	
)

var (
	rdb *redis.Client
)

func ConnectRedis() {
	// fungsi connect redis
	rdb = redis.NewClient(&redis.Options{
        Addr:     "containers-us-west-157.railway.app:6868",
        Password: "qoRSgaJ6WFdUDBVSv5NC",
        DB:       0,
    })

	return
}
