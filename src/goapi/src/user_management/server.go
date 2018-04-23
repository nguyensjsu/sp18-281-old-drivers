package main

import (
	"github.com/go-redis/redis"
	"github.com/satori/go.uuid"
	"github.com/tinylib/msgp/_generated"
	"encoding/json"
)

type UserManager struct {
	redisClient *redis.Client
}

type User struct {
	UserId string
	UserName string
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

// return UserJsonString success
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
func (um *UserManager) CreateUser() (string, bool){
	if !um.isOpen() {
		return "", false
	}
	uuid, _ := uuid.NewV4()
	user := &User{
		UserId: uuid.String(),
		UserName:
	}
	buf, err := json.Marshal(user)
	if err != nil {
		return "", false
	}
	val := string(buf[:])
	err := um.redisClient.Set(user.UserId, val,0)
	if err != nil {
		return "", false
	}
	return val, true
}