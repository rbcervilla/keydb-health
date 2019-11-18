package main

import (
	"flag"
	"fmt"
	"github.com/go-redis/redis/v7"
	"net/http"
	"strings"
	"time"
)

func main() {

	h := flag.String("h", "127.0.0.1", "KeyDB host")
	p := flag.String("p", "6379", "KeyDB port")
	sp := flag.String("sp", "8080", "Server listener port")
	flag.Parse()

	client := redis.NewClient(&redis.Options{
		Addr:        fmt.Sprintf("%s:%s", *h, *p),
		DialTimeout: time.Second,
		ReadTimeout: time.Second,
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {

		result, err := client.Info().Result()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		info := parseInfo(result)
		if info["loading"] != "0" || info["master_sync_in_progress"] != "0" {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})

	if err := http.ListenAndServe(":"+*sp, nil); err != nil {
		panic(err)
	}
}

func parseInfo(str string) map[string]string {

	info := make(map[string]string)
	lines := strings.Split(str, "\r\n")

	for _, line := range lines {

		if strings.HasPrefix(line, "#") {
			continue
		}

		pair := strings.Split(line, ":")
		if len(pair) != 2 {
			continue
		}

		info[pair[0]] = pair[1]
	}

	return info
}
