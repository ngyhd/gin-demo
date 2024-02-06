package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gin-demo/internal/config"
	"gin-demo/internal/model"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"strconv"
)

//Redis key 规范 s:gin-demo:xxx  s代表key的类型为string gin-demo为服务名 xxx为自定义值
// s代表string
// hs代表hashmap
// se代表set
// zs代表zset
// l代表list
// bf代表布隆过滤器
// hy代表hyperloglog
// b代表bitmap

const serviceName = "gin-demo"

func userInfoKey(userId string) string {
	return fmt.Sprintf("s:%v:userinfo:%v", serviceName, userId)
}

func GetUserInfo(ctx context.Context, userId string) (*model.User, error) {
	result, err := config.GetRedis().Get(ctx, userInfoKey(userId)).Result()
	if err != nil {
		return nil, err
	}
	if result == "" {
		return nil, redis.Nil
	}
	var user model.User
	err = json.Unmarshal([]byte(result), &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func SetUserInfo(ctx context.Context, user model.User) error {
	marshal, err := json.Marshal(user)
	if err != nil {
		return err
	}
	_, err = config.GetRedis().Set(ctx, userInfoKey(strconv.Itoa(user.Id)), marshal, 0).Result()
	if err != nil {
		return err
	}
	return nil
}

func DelUserInfo(ctx context.Context, userId string) error {
	_, err := config.GetRedis().Del(ctx, userInfoKey(userId)).Result()
	if err != nil {
		return err
	}
	return nil
}

// RefreshUserInfo 刷新用户信息缓存
func RefreshUserInfo(ctx context.Context, userId string) (*model.User, error) {
	resp := model.User{}
	tx := config.GetDB().Where("id = ?", userId).First(&resp)
	if tx.Error != nil && !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return &resp, tx.Error
	}
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return &resp, nil
	}
	err := SetUserInfo(ctx, resp)
	if err != nil {
		return &resp, err
	}
	return &resp, nil
}
