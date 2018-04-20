import main
import "github.com/go-redis/redis"
import "fmt"

func NewClient() {
	client := redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
// Output: PONG <nil>
}
