package spider

import (
	"testing"

	"github.com/Agzdjy/proxy-pool/storage"
	"github.com/go-redis/redis"
)

var spider Spider = &Ip181{}

func TestDo(t *testing.T) {
	store := storage.NewRedisClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	err := spider.Do("http://www.ip181.com/", store)

	if err != nil {
		t.Error("spider do error", err)
		return
	}
}
