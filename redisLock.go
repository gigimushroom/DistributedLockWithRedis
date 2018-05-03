package redlock

import(
	"github.com/go-redis/redis"
	"fmt"
	"time"
)


type RedLock struct {
	C *redis.Client
	Password int64
	Key string
}

// NewRedisHandler creates new redis handler
func NewRedisHandler(addr string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (r *RedLock) SetKey(key string) {
	r.Key = key
}

func (r *RedLock) genRandomValue() int64 {
	r.Password = time.Now().UnixNano()
	return r.Password
}

func (r *RedLock) AcquireLock(block bool) bool {

	if r.Password == 0 {
		r.genRandomValue()
	}

	px := time.Duration(time.Second * 3)
	ok, err := r.C.SetNX(r.Key, r.Password, px).Result()
	if err != nil {
		panic(err)
	}
	if ok {
		fmt.Println("get the lock")
	} 
	return ok
}

func (r *RedLock) ReleaseLock() bool {
	val, err := r.C.Get(r.Key).Int64()
	if err != nil {
		return false
	} 

	if val == r.Password {
		fmt.Println("password match, delete the key")
		_, err = r.C.Del(r.Key).Result()
		if err != nil {
			fmt.Println("Del failed")
			return false
		}
		return true
	}

	return false
}
