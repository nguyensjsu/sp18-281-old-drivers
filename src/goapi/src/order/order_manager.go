package main

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/satori/go.uuid"
	"log"
	"strconv"
	"time"
)

type OrderManager struct {
	redisClient *redis.Client
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

func NewOrderManager(ipAddr string, port int) *OrderManager {
	addr := ipAddr + ":" + strconv.Itoa(port)
	log.Printf("Redis Connection address %s\n", addr)
	client := redis.NewClient(&redis.Options{
		Addr: addr})

	om := &OrderManager{
		redisClient: client}
	return om
}

func (om *OrderManager) isOpen() bool {
	return om.redisClient != nil
}

// return OrderJsonString, success
func (om *OrderManager) GetOrder(orderId string) (string, bool) {
	if !om.isOpen() {
		log.Println("Bad Connection")
		return "", false
	}

	val, err := om.redisClient.Get(orderId).Result()
	if err != nil {
		log.Printf("Read from redis failed %v\n", err)
		return "", false
	}
	return val, true
}

func (om *OrderManager) UpdateOrder(newOrder *Order) bool {
	if !om.isOpen() {
		log.Printf("Bad Connection")
		return false
	}
	buf, err := json.Marshal(newOrder)
	if err != nil {
		log.Printf("Jsonize failed %v\n", err)
		return false
	}

	val := string(buf[:])
	if err := om.redisClient.Set(newOrder.OrderId, val, 0).Err(); err != nil {
		log.Printf("Set redis failed key: %s err: %v\n", newOrder.OrderId, err)
		return false
	}
	return true
}

// Return true on success
func (om *OrderManager) DeleteOrder(orderId string) bool {
	if !om.isOpen() {
		log.Println("Bad Connection")
		return false
	}

	if err := om.redisClient.Del(orderId).Err(); err != nil {
		log.Printf("Redis delete failed %v\n", err)
		return false
	}
	return true
}

func (om *OrderManager) CreateOrder(userId string, items []string) (string, bool) {
	if !om.isOpen() {
		log.Println("Bad Connection")
		return "", false
	}
	uuid, _ := uuid.NewV4()
	order := &Order{
		OrderId: uuid.String(),
		UserId:  userId,
		Items:   items,
		Date:    time.Now().String(),
		Status:  CREATED}

	var val string
	if buf, err := json.Marshal(order); err != nil {
		log.Println("Jsonize failed")
		return "", false
	} else {
		val = string(buf)
	}
	// Use redis Pipeline
	pipe := om.redisClient.Pipeline()
	// Add to order set
	if err := pipe.Set(order.OrderId, val, 0).Err(); err != nil {
		log.Printf("Set failed %v\n", err)
		return "", false
	}
	// Add to user list
	if err := pipe.RPush(userId, order.OrderId).Err(); err != nil {
		log.Printf("RPush failed %v\n", err)
		return "", false
	}

	if _, err := pipe.Exec(); err != nil {
		log.Printf("Exec failed %v\n", err)
		return "", false
	}
	return val, true
}

func (om *OrderManager) GetOrderByUser(userId string) ([]string, bool) {
	if !om.isOpen() {
		log.Println("Not open")
		return nil, false
	}
	// Get list length
	len, err := om.redisClient.LLen(userId).Result()
	if err != nil {
		log.Printf("Redis LLen failed key: %s, err: %v\n", userId, err)
		return nil, false
	}
	// Get order ids
	ids, err := om.redisClient.LRange(userId, 0, len-1).Result()
	if err != nil {
		log.Printf("Redis LRange failed key: %s, err: %v\n", userId, err)
		return nil, false
	}

	var orders []string
	for _, id := range ids {
		orderVal, ok := om.GetOrder(id)
		if !ok {
			log.Printf("Order %s not existed\n", id)
			continue
		}
		orders = append(orders, orderVal)
	}
	return orders, true
}
