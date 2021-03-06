package storage

import (
	"encoding/json"
	"strings"

	"github.com/Agzdjy/proxy-pool/model"
	"github.com/go-redis/redis"
)

type Redis struct {
	client *redis.Client
}

var _ Storage = &Redis{}

func genKey(protocol string) string {
	return "ip_proxy:" + strings.ToLower(protocol)
}

func decodeValue(ipStr string) (ip *model.IP, err error) {
	ip = &model.IP{}
	err = json.Unmarshal([]byte(ipStr), ip)
	if err != nil {
		return nil, err
	}
	return ip, nil
}

func encodeValue(ip *model.IP) string {
	b, _ := json.Marshal(ip)
	return string(b)
}

func (r *Redis) Save(ip *model.IP) error {
	key := genKey(ip.Protocol)
	err := r.client.SAdd(key, encodeValue(ip)).Err()
	return err
}

func (r *Redis) Del(ip *model.IP) bool {
	key := genKey(ip.Protocol)
	value := encodeValue(ip)
	err := r.client.SRem(key, value).Err()
	if err != nil {
		return false
	}

	return true
}

func (r *Redis) RangeOne(protocol string) (ip *model.IP, err error) {
	key := genKey(protocol)
	val, err := r.client.SRandMember(key).Result()
	if err != nil {
		return nil, err
	}

	ip, err = decodeValue(val)
	return
}

func (r *Redis) Close() error {
	return r.client.Close()
}
