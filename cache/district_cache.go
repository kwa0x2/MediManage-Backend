package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/kwa0x2/MediManage-Backend/models"
	"strings"
)

type DistrictCache struct {
	RedisClient *redis.Client
}

func (c *DistrictCache) SetAllDistrictsByProvince(districts []*models.District, provinceName string) error {
	ctx := context.Background()

	data, err := json.Marshal(districts)
	if err != nil {
		return err
	}
	key := fmt.Sprintf("districts_all_%s", strings.ToLower(provinceName))
	err = c.RedisClient.Set(ctx, key, data, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (c *DistrictCache) GetAllDistrictsByProvince(provinceName string) ([]*models.District, error) {
	ctx := context.Background()
	key := fmt.Sprintf("districts_all_%s", strings.ToLower(provinceName))
	val, err := c.RedisClient.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var districts []*models.District
	err = json.Unmarshal([]byte(val), &districts)
	if err != nil {
		return nil, err
	}

	return districts, nil
}
