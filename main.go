package main

import (
	"os"

	"github.com/jamesrwhite/godeletemytweets/client"
)

func main() {
	consumerKey := os.Getenv("CONSUMER_KEY")
	consumerSecret := os.Getenv("CONSUMER_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN")
	accessSecret := os.Getenv("ACCESS_SECRET")

	client.New(
		consumerKey,
		consumerSecret,
		accessToken,
		accessSecret,
	)
}
