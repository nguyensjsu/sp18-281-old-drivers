package main

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/satori/go.uuid"
	"time"
)

type ReviewManager struct {
	redisClient *redis.Client
}

type Review struct {
	ReviewId string
	UserId  string
	Item    string
	Date    string
}

func NewReviewManager(ipAddr string, port int) *ReviewManager {
	addr := ipAddr + ":" + string(port)
	client := redis.NewClient(&redis.Options{
		Addr: addr})

	om := &ReviewManager{
		redisClient: client}
	return om
}

func (om *ReviewManager) isOpen() bool {
	return om.redisClient != nil
}

// return ReviewJsonString, success
func (om *ReviewManager) GetReview(reviewId string) (string, bool) {
	if !om.isOpen() {
		return "", false
	}

	val, err := om.redisClient.Get(reviewId).Result()
	if err != nil {
		return "", false
	}
	return val, true
}

func (om *ReviewManager) UpdateReview(newReview *Review) bool {
	if !om.isOpen() {
		return false
	}
	buf, err := json.Marshal(newReview)
	if err != nil {
		return false
	}

	val := string(buf[:])
	if err := om.redisClient.Set(newReview.ReviewId, val, 0).Err(); err != nil {
		return false
	}
	return true
}

// Return true on success
func (om *ReviewManager) DeleteReview(reviewId string) bool {
	if !om.isOpen() {
		return false
	}

	err := om.redisClient.Del(reviewId)
	if err != nil {
		return false
	}
	return true
}

func (om *ReviewManager) CreateReview(userId string, item string) (string, bool) {
	if !om.isOpen() {
		return "", false
	}
	uuid, _ := uuid.NewV4()
	review := &Review{
		ReviewId: uuid.String(),
		UserId:   userId,
		Item:     item,
		Date:     time.Now().String()}

	buf, err := json.Marshal(review)
	if err != nil {
		return "", false
	}
	val := string(buf[:])
	// Use redis Pipeline
	pipe := om.redisClient.Pipeline()
	// Add to review set
	if err := pipe.Set(review.UserId, val, 0); err != nil {
		return "", false
	}
	// Add to user list
	if err := pipe.RPush(userId, review.UserId).Err(); err != nil {
		return "", false
	}

	if _, err := pipe.Exec(); err != nil {
		return "", false
	}
	return val, true
}

func (om *ReviewManager) GetReviewByInventory(inventoryId string) ([]string, bool) {
	if !om.isOpen() {
		return nil, false
	}
	// Get list length
	len, err := om.redisClient.LLen(inventoryId).Result()
	if err != nil {
		return nil, false
	}
	// Get inventory ids
	ids, err := om.redisClient.LRange(inventoryId, 0, len-1).Result()
	if err != nil {
		return nil, false
	}

	var reviews []string
	for _, id := range ids {
		review, ok := om.GetReview(id)
		if !ok {
			return nil, false
		}
		reviews = append(reviews, review)
	}
	return reviews, false
}
