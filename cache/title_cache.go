package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/kwa0x2/MediManage-Backend/models"
	"strings"
	"time"
)

type TitleCache struct {
	RedisClient *redis.Client
}

func (c *TitleCache) SetAllTitleByJobGroupName(titles []*models.Title, jobGroupName string) error {
	ctx := context.Background()

	data, err := json.Marshal(titles)
	if err != nil {
		return err
	}

	expiration := 24 * time.Hour

	key := fmt.Sprintf("title_%s", strings.ToLower(jobGroupName))
	err = c.RedisClient.Set(ctx, key, data, expiration).Err()
	if err != nil {
		return err
	}

	return nil
}

func (c *TitleCache) GetAllTitleByJobGroupName(jobGroupName string) ([]*models.Title, error) {
	ctx := context.Background()
	key := fmt.Sprintf("titles_%s", strings.ToLower(jobGroupName))
	val, err := c.RedisClient.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var titles []*models.Title
	err = json.Unmarshal([]byte(val), &titles)
	if err != nil {
		return nil, err
	}

	return titles, nil
}
