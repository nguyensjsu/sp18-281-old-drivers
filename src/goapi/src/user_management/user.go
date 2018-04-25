package main

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/satori/go.uuid"
	"strconv"
)

var redis_server_ip = "127.0.0.1"
var redis_server_port = 6379

type UserManager struct {
	redisClient *redis.Client
}

type User struct {
	UserId string
	UserName string
	Phone string
	Balance int
}

func NewUserManager(ipAddr string, port int) *UserManager {
	addr := ipAddr + ":" + string(port)
	client := redis.NewClient(&redis.Options{
		Addr: addr})

	um := &UserManager{
		redisClient: client}
	return um
}

func (um *UserManager) isOpen() bool {
	return um.redisClient != nil
}

// get user
func (um *UserManager) GetUser(userId string) (string, bool) {
	if !um.isOpen() {
		return "", false
	}
	val, err := um.redisClient.Get(userId).Result()
	if err != nil {
		return "", false
	}
	return val, true
}

// update user
func (um *UserManager) UpdateUser(newUser *User) bool {
	if !um.isOpen() {
		return false
	}
	buf, err := json.Marshal(newUser)
	if err != nil {
		return false
	}
	val := string(buf[:])
	if err := um.redisClient.Set(newUser.UserId, val, 0).Err(); err != nil {
		return false
	}
	return true
}

// delete user
func (um *UserManager) DeleteUser(userId string) bool {
	if !um.isOpen() {
		return false
	}
	err := um.redisClient.Del(userId)
	if err != nil {
		return false
	}
	return true
}

// create User
func (um *UserManager) CreateUser(username string, phone string, balance string) (string, bool){
	if !um.isOpen() {
		return "", false
	}
	bal, err := strconv.Atoi(balance)
	if err != nil {
		return "", false
	}
	uuid, _ := uuid.NewV4()
	user := &User{
		UserId: uuid.String(),
		UserName: username,
		Phone: phone,
		Balance: bal}
	buf, err := json.Marshal(user)
	if err != nil {
		return "", false
	}
	val := string(buf[:])
	err = um.redisClient.Set(user.UserId, val,0).Err()
	if err != nil {
		return "", false
	}
	return val, true
}