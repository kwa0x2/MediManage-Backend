package cache

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/kwa0x2/MediManage-Backend/models"
	"time"
)

type ClinicCache struct {
	RedisClient *redis.Client
}

func (c *ClinicCache) SetAllClinics(clinics []*models.Clinic) error {
	ctx := context.Background()

	data, err := json.Marshal(clinics)
	if err != nil {
		return err
	}

	expiration := 24 * time.Hour

	err = c.RedisClient.Set(ctx, "clinics", data, expiration).Err()
	if err != nil {
		return err
	}

	return nil
}

func (c *ClinicCache) GetAllClinics() ([]*models.Clinic, error) {
	ctx := context.Background()

	val, err := c.RedisClient.Get(ctx, "clinics").Result()
	if err != nil {
		return nil, err
	}

	var clinics []*models.Clinic
	err = json.Unmarshal([]byte(val), &clinics)
	if err != nil {
		return nil, err
	}

	return clinics, nil
}
