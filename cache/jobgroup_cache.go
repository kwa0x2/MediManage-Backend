package cache

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/kwa0x2/MediManage-Backend/models"
	"time"
)

type JobGroupCache struct {
	RedisClient *redis.Client
}

func (c *JobGroupCache) SetAllJobGroups(jobGroups []*models.JobGroup) error {
	ctx := context.Background()

	data, err := json.Marshal(jobGroups)
	if err != nil {
		return err
	}

	expiration := 24 * time.Hour

	err = c.RedisClient.Set(ctx, "job_groups", data, expiration).Err()
	if err != nil {
		return err
	}

	return nil
}

func (c *JobGroupCache) GetAllJobGroups() ([]*models.JobGroup, error) {
	ctx := context.Background()

	val, err := c.RedisClient.Get(ctx, "job_groups").Result()
	if err != nil {
		return nil, err
	}

	var jobGroups []*models.JobGroup
	err = json.Unmarshal([]byte(val), &jobGroups)
	if err != nil {
		return nil, err
	}

	return jobGroups, nil
}
