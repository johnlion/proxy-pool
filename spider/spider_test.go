package spider

import (
	"testing"

	"github.com/Agzdjy/go-proxy/storage"
	"github.com/go-redis/redis"
)

var spider Spider = &ip181{}

func TestDo(t *testing.T) {
	store := storage.NewRedisClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	count, err := spider.Do("http://www.ip181.com/", store)

	if err != nil {
		t.Errorf("spider do error", err)
		return
	}

	if count < 1 {
		t.Error("spider run but got none")
		return
	}
}