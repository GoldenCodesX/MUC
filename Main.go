package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

var name string

func main() {
	fmt.Println("How many characters would you like your username to have?")

	var characters int

	fmt.Scanln(&characters)

	fmt.Println("Username: ", name)

	fmt.Println(getHTTP(characters))
}

func getHTTP(amount int) (response string) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.mojang.com/users/profiles/minecraft/" + usernameCheck(amount), nil)

	req.Header.Set("User-Agent", "test")
	if err != nil {
		return string(err.Error())
	}

	resp, err := client.Do(req)
	if err != nil {
		return string(err.Error())
	}

	defer resp.Body.Close()

	body, e := ioutil.ReadAll(resp.Body)

	if err != nil {
		return string(e.Error())
	}

	if body == nil {
		fmt.Println("This username is available!")
	}

	return string(body)
}

// Make the request to Mojang to check if the username is available, and check when it will be available!
func usernameCheck(username int) string {
	if username >= 3 {
		name = String(username)
		return String(username)
	} else {
		return "Error"
	}
}

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}
