package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/go-redis/redis"
	"github.com/urfave/negroni"
	"github.com/gorilla/mux"
	"net/url"
	"github.com/satori/go.uuid"
	"github.com/unrolled/render"
)


type InventoryManager struct {
	redisClient *Client
}

type Inventory struct {
	InventoryId		string
	InventoryName	string
	InventoryPrice	float32
	Amount  		uint32
	Reviews			[]string
}

// redis sentinel for automatic failover
func NewInventoryManager(master string, sentinels []string) *InventoryManager {
	client := redis.NewFailoverClient(&redis.FailoverOption{
		MasterName:    master,
		SentinelAddrs: sentinels})

	im := &InventoryManager{
		redisClient: client}
	return im
}

// verify connection
func (im *InventoryManager) isOpen() bool {
	return im.redisClient != nil
}

// get all inventory
func (im *InventoryManager) GetAllInventory() ([]string, bool) {
	if !im.isOpen() {
		return "", false
	}

	// get list length
	len, err := im.redisClient.LLen(*).Result()
	if err != nil {
		return "", false
	}

	val, err := im.redisClient.Get(*).Result()
	if err != nil {
		return "", false
	}

	return val, true
}

// get inventory by inventoryId
func (im *InventoryManager) GetInventory(inventoryId string) (string, bool) {
	if !im.isOpen() {
		return "", false
	}

	val, err := im.redisClient.Get(inventoryId).Result()
	if err != nil {
		return "", false
	}

	return val, true
}

// create new inventory
func (im *InventoryManager) CreateInventory(name string, price float32, amount uint32) (string, bool) {
	if !im.isOpen() {
		return "", false
	}

	uuid, _ := uuid.NewV4()
	inventory := &Inventory {
		InventoryId:	uuid.string(),
		InventoryName:	name,
		InventoryPrice: price,
		Amount:			amount
	}

	buf, err := json.Marshal(inventory)
	if err != nil {
		return "", false
	}

	val := string(buf[:])

	result, err := im.redisClient.Set(inventory.InventoryId, val, 0)

	if err != nil {
		return "", false
	}

	return result, true
		
}

// update inventory
func (im *InventoryManager) UpdateInventory(inventory *Inventory) bool {
	if !im.isOpen() {
		return false
	}

	buf, err := json.Marshal(inventory)
	if err != nil {
		return false
	}

	val := string(buf[:])
	result, err := im.redisClient.Set(inventory.InventoryId, val, 0)

	if err != nil {
		return false		
	}

	return true
}

// delete inventory
func (im *InventoryManager) DeleteInventory(inventoryId string) bool {
	if !im.isOpen() {
		return false
	}

	err := im.redisClient.Del(inventoryId)

	if err != nil {
		return false
	}

	return true
}

