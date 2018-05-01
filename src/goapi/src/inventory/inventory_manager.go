package main

import (
	"encoding/json"
	// "fmt"
	// "net/http"
	"github.com/go-redis/redis"
	// "github.com/codegangsta/negroni"
	// "github.com/gorilla/mux"
	// "net/url"
	"github.com/satori/go.uuid"
	// "github.com/unrolled/render"
	"log"
	"strconv"
)

type InventoryManager struct {
	redisClient *redis.ClusterClient
}

type Inventory struct {
	InventoryId    string
	InventoryName  string
	InventoryPrice float64
	Amount         int64
	Reviews        []string
}

// redis sentinel for automatic failover
func NewInventoryManager(addrs []string) *InventoryManager {
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: addrs})

	im := &InventoryManager{
		redisClient: client}
	return im
}

// verify connection
func (im *InventoryManager) isOpen() bool {
	return im.redisClient != nil
}

// get all inventory
func (im *InventoryManager) GetAllInventory() (map[string]string, bool) {
	if !im.isOpen() {
		return nil, false
	}

	// get list length
	var cursor uint64
	var n int
	for {
		var keys []string
		var err error
		keys, cursor, err = im.redisClient.Scan(cursor, "", 10).Result()
		if err != nil {
			panic(err)
		}
		n += len(keys)
		if cursor == 0 {
			break
		}
	} // n represents the sum of keys

	// create map store all the key-value pair
	inventoryMap := make(map[string]string)
	for i := 0; i < n; i++ {

		keyv, _ := im.redisClient.Get("key" + strconv.Itoa(i)).Result()
		kvalue, _ := im.redisClient.Get(keyv).Result()
		inventoryMap[keyv] = kvalue

	}

	return inventoryMap, true
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
func (im *InventoryManager) CreateInventory(name string, price float64, amount int64) (string, bool) {
	if !im.isOpen() {
		return "", false
	}

	uuid, _ := uuid.NewV4()
	inventory := &Inventory{
		InventoryId:    uuid.String(),
		InventoryName:  name,
		InventoryPrice: price,
		Amount:         amount}

	buf, err := json.Marshal(inventory)
	if err != nil {
		log.Printf("inventory marshal failed %v\n")
		return "", false
	}

	val := string(buf[:])

	err = im.redisClient.Set(inventory.InventoryId, val, 0).Err()

	if err != nil {
		log.Printf("inventroy set failed")
		return "", false
	}

	return val, true

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
	im.redisClient.Set(inventory.InventoryId, val, 0)

	/*
		if err != nil {
			return false
		}
	*/

	return true
}

// delete inventory
func (im *InventoryManager) DeleteInventory(inventoryId string) bool {
	if !im.isOpen() {
		return false
	}

	err := im.redisClient.Del(inventoryId).Err()

	if err != nil {
		log.Printf("delete failed")
		return false
	}

	return true
}
