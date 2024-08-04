package cache

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/kwa0x2/MediManage-Backend/models"
	"time"
)

type ProvinceCache struct {
	RedisClient *redis.Client
}

func (c *ProvinceCache) SetAllProvince(provinces []*models.Province) error {
	ctx := context.Background()

	data, err := json.Marshal(provinces)
	if err != nil {
		return err
	}

	expiration := 24 * time.Hour

	err = c.RedisClient.Set(ctx, "province_all", data, expiration).Err()
	if err != nil {
		return err
	}

	return nil
}

func (c *ProvinceCache) GetAllProvince() ([]*models.Province, error) {
	ctx := context.Background()

	val, err := c.RedisClient.Get(ctx, "province_all").Result()
	if err != nil {
		return nil, err
	}

	var provinces []*models.Province
	err = json.Unmarshal([]byte(val), &provinces)
	if err != nil {
		return nil, err
	}

	return provinces, nil
}
