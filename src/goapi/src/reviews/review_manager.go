package main

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/satori/go.uuid"
	"log"
	"time"
	"util"
)

type ReviewManager struct {
	redisClient *redis.ClusterClient
}

type Review struct {
	ReviewId string
	UserId   string
	Item     string
	Content  string
	Date     string
}

func NewReviewManager(addrs []string) *ReviewManager {
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: addrs})

	rm := &ReviewManager{
		redisClient: client}
	return rm
}

func (rm *ReviewManager) isOpen() bool {
	return rm.redisClient != nil
}

// return ReviewJsonString, success
func (rm *ReviewManager) GetReview(userId string, reviewId string) (string, bool) {
	if !rm.isOpen() {
		log.Println("Bad Connection")
		return "", false
	}

	key := util.GenKey(userId, reviewId)
	val, err := rm.redisClient.Get(key).Result()
	if err != nil {
		log.Printf("Read frrm redis failed %v\n", err)
		return "", false
	}
	return val, true
}

func (rm *ReviewManager) UpdateReview(review *Review, newContent string) bool {
	if !rm.isOpen() {
		return false
	}
	if len(newContent) != 0 {
		review.Content = newContent
	}
	buf, err := json.Marshal(review)
	if err != nil {
		return false
	}

	val := string(buf[:])
	key := util.GenKey(review.UserId, review.ReviewId)
	if err := rm.redisClient.Set(key, val, 0).Err(); err != nil {
		return false
	}
	return true
}

// Return true on success
func (rm *ReviewManager) DeleteReview(userId string, reviewId string) bool {
	if !rm.isOpen() {
		return false
	}

	key := util.GenKey(userId, reviewId)
	if err := rm.redisClient.Del(key).Err(); err != nil {
		log.Printf("Redis delete failed %v\n", err)
		return false
	}
	return true
}

func (rm *ReviewManager) CreateReview(userId string, item string, content string) (string, bool) {
	if !rm.isOpen() {
		return "", false
	}
	uuid, _ := uuid.NewV4()
	review := &Review{
		ReviewId: uuid.String(),
		UserId:   userId,
		Item:     item,
		Content:  content,
		Date:     time.Now().String()}

	buf, err := json.Marshal(review)
	if err != nil {
		log.Printf("Marchal error.")
		return "", false
	}
	val := string(buf[:])
	// Use redis Pipeline
	pipe := rm.redisClient.Pipeline()
	// Add to review set
	key := util.GenKey(userId, review.ReviewId)
	if err := pipe.Set(key, val, 0).Err(); err != nil {
		log.Printf("Set error.")
		return "", false
	}
	// Add to user list
	if err := pipe.RPush(userId, review.ReviewId).Err(); err != nil {
		log.Printf("Push error.")
		return "", false
	}

	if _, err := pipe.Exec(); err != nil {
		log.Printf("Exec error %v\n.", err)
		return "", false
	}
	return val, true
}

func (rm *ReviewManager) GetReviewByReviewId(userId string) ([]string, bool) {
	if !rm.isOpen() {
		return nil, false
	}
	// Get list length
	len, err := rm.redisClient.LLen(userId).Result()

	if err != nil {
		log.Printf("Get list failed!")
		return nil, false
	}
	// Get inventory ids
	ids, err := rm.redisClient.LRange(userId, 0, len-1).Result()
	if err != nil {
		log.Printf("Get IDs failed!")
		return nil, false
	}

	var reviews []string
	for _, id := range ids {
		review, ok := rm.GetReview(userId, id)
		if !ok {
			log.Printf("Get reviews failed!")
			return nil, false
		}
		reviews = append(reviews, review)
	}
	return reviews, true
}
