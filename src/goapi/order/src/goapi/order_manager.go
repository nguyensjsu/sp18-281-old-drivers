package main

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/satori/go.uuid"
	"time"
)

type OrderManager struct {
	redisClient *Client
}

type OrderStatus int

const (
	CREATED   OrderStatus = 0
	COMPLETED             = 1
	FAILED                = 2
)

type Order struct {
	OrderId string
	UserId  string
	Items   []string
	Date    string
	Status  OrderStatus
}

func NewOrderManager(master string, sentinels []string) *OrderManager {
	client := redis.NewFailoverClient(&redis.FailoverOption{
		MasterName:    master,
		SentinelAddrs: sentinels})

	om := &OrderManager{
		redisClient: client}
	return om
}

func (om *OrderManager) isOpen() bool {
	return om.redisClient != nil
}

// return OrderJsonString, success
func (om *OrderManager) GetOrder(orderId string) (string, bool) {
	if !isOpen() {
		return "", false
	}

	val, err := om.redisClient.Get(order).Result()
	if err != nil {
		return "", false
	}
	return val, true
}

func (om *OrderManager) UpdateOrder(newOrder *Order) bool {
	if !isOpen() {
		return false
	}
	buf, err := json.Marshal(newOrder)
	if err != nil {
		return false
	}

	val := string(buf[:])
	err := om.redisClient.Set(newOrder.OrderId, val).Err()
	if err != nil {
		return false
	}
	return true
}

// Return true on success
func (om *OrderManager) DeleteOrder(orderId) bool {
	if !isOpen() {
		return false
	}

	_, err := om.redisClient.Del(orderid)
	if err != nil {
		return false
	}
	return true
}

func (om *OrderManager) CreateOrder(userId string, items []string) (string, bool) {
	if !isOpen() {
		return false
	}
	uuid, _ := uuid.NewV4()
	order := &Order{
		OrderId: uuid.String(),
		UserId:  userId,
		Items:   items,
		Date:    time.Now().String(),
		Status:  CREATED}

	buf, err := json.Marshal(order)
	if err != nil {
		return "", false
	}
	val := string(buf[:])
	// Use redis Pipeline
	pipe := om.redisClient.Pipeline()
	// Add to order set
	err := pipe.Set(order.OrderId, val)
	if err != nil {
		return "", false
	}
	// Add to user list
	_, err := pipe.RPush(userId, order.OrderId)
	if err != nil {
		return "", false
	}

	_, err := pipe.Exec()
	if err != nil {
		return "", false
	}
	return val, true
}

func (om *OrderManager) GetOrderByUser(userId string) ([]string, bool) {
	if !isOpen() {
		return nil, false
	}
	om.redisClient
	// Get list length
	len, err := om.redisClient.LLen(userId).Result()
	if err != nil {
		return "", false
	}
	// Get order ids
	ids, err := om.redisClient.LRange(userId, 0, len-1).Result()
	if err != nil {
		return "", false
	}

	var orders []string
	for _, id := range ids {
		orderVal, ok := om.GetOrder(id)
		if !ok {
			return nil, false
		}
		orders = append(orders, orderVal)
	}
	return orders, false
}
