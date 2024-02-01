package store

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// EmailVerificationRepository is a struct that defines the email verification repository
type EmailVerificationRepository struct {
	db         *redis.Client
	expiration time.Duration
}

// NewEmailVerificationRepository is a function that returns a new email verification repository
func NewEmailVerificationRepository(db *redis.Client, expiration time.Duration) *EmailVerificationRepository {
	return &EmailVerificationRepository{db, expiration}
}

// Create is a method that creates a new email verification token
func (r *EmailVerificationRepository) Create(ctx context.Context, email string, token string) error {

	key := fmt.Sprintf("email-verification:%s", email)

	return r.db.Set(ctx, key, token, r.expiration).Err()
}

// Get is a method that gets an email verification token
func (r *EmailVerificationRepository) Get(ctx context.Context, email string) (string, error) {

	key := fmt.Sprintf("email-verification:%s", email)

	return r.db.Get(ctx, key).Result()
}

// Delete is a method that deletes an email verification token
func (r *EmailVerificationRepository) Delete(ctx context.Context, email string) error {

	key := fmt.Sprintf("email-verification:%s", email)

	return r.db.Del(ctx, key).Err()
}
