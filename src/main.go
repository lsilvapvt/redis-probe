package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

func homepage(w http.ResponseWriter, r *http.Request) {

	redisServerAddress := getEnv("REDIS_SERVER_ADDRESS", "localhost")
	redisServerPort := getEnv("REDIS_SERVER_PORT", "6379")
	redisDatabase := getEnv("REDIS_DATABASE", "0")

	fmt.Fprintln(w, "Probe for Redis server "+redisServerAddress+":"+redisServerPort+", database "+redisDatabase)
	fmt.Fprintln(w, "Endpoints:")
	fmt.Fprintln(w, "  /ping    - sends a PING to Redis server")
	fmt.Fprintln(w, "  /set     - saves the current time as the value for a MYTIME on the server")
	fmt.Fprintln(w, "  /get     - retrieves the value of the MYTIME variable stored on the server")
	fmt.Fprintln(w, "  /delete  - deletes the MYTIME variable stored on the server")
}

func pingRedisServer(w http.ResponseWriter, r *http.Request) {

	pong, err := redisClient.Ping().Result()

	if err != nil {
		fmt.Fprintln(w, err)
	} else {
		fmt.Fprintln(w, pong)
	}

}

func setMyTimeVariable(w http.ResponseWriter, r *http.Request) {

	mytime := time.Now().Format("2006-01-02 15:04:05")
	err := redisClient.Set("MYTIME", mytime, 0).Err()

	if err != nil {
		fmt.Fprintln(w, err)
	} else {
		fmt.Fprintln(w, "MYTIME variable set to "+mytime)
	}

}

func getMyTimeVariable(w http.ResponseWriter, r *http.Request) {

	val, err := redisClient.Get("MYTIME").Result()

	if err != nil {
		fmt.Fprint(w, "Variable not found or error while retrieving it: ")
		fmt.Fprintln(w, err)
	} else {
		fmt.Fprintln(w, "MYTIME variable retrieved: "+val)
	}
}

func deleteMyTimeVariable(w http.ResponseWriter, r *http.Request) {

	result, err := redisClient.Del("MYTIME").Result()

	if err != nil {
		fmt.Fprintln(w, err)
	} else {
		fmt.Fprint(w, "Number of entries deleted: ")
		fmt.Fprintln(w, result)
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func handleRequests() {
	proveServicePort := getEnv("REDIS_PROBE_SERVICE_PORT", "8888")

	http.HandleFunc("/", homepage)
	http.HandleFunc("/ping", pingRedisServer)
	http.HandleFunc("/set", setMyTimeVariable)
	http.HandleFunc("/get", getMyTimeVariable)
	http.HandleFunc("/delete", deleteMyTimeVariable)
	log.Fatal(http.ListenAndServe(":"+proveServicePort, nil))
}

func initRedisServerConnection() {
	redisServerAddress := getEnv("REDIS_SERVER_ADDRESS", "localhost")
	redisServerPort := getEnv("REDIS_SERVER_PORT", "6379")
	redisPassword := getEnv("REDIS_PASSWORD", "password")
	redisDatabase, _ := strconv.Atoi(getEnv("REDIS_DATABASE", "0"))

	redisClient = redis.NewClient(&redis.Options{
		Addr:     redisServerAddress + ":" + redisServerPort,
		Password: redisPassword,
		DB:       redisDatabase,
	})
}

func main() {
	initRedisServerConnection()
	handleRequests()
}

// export REDIS_SERVER_ADDRESS=
// export REDIS_SERVER_PORT=6379
// export REDIS_PASSWORD=
// export REDIS_DATABASE=0
// export REDIS_PROBE_SERVICE_PORT=80
