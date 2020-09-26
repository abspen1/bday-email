package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/my/repo/GitHub/bday-email/go-implementation/emailer"
	"github.com/my/repo/go/src/github.com/joho/godotenv"
)

// {"id":147,"setup":"How much does a hipster weigh?","punchline":"An instagram.","type":"general"}
type joke struct {
	Identity  int    `json:"id"`
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
}

var types = make([]string, 2)

var jokeURL = "https://us-central1-dadsofunny.cloudfunctions.net/DadJokes/random/type/"

func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func getJoke(info emailer.Birthday) emailer.Birthday {
	// Set seed so we don't get same random numbers each time
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// Get either a 1 or 0 at random
	index := (r1.Intn(2))
	types[0] = "knock-knock"
	types[1] = "general"
	url := jokeURL + types[index]

	jokeClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := jokeClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var jokes []joke
	jsonErr := json.Unmarshal(body, &jokes)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	for _, jokes1 := range jokes {
		info.JokePunchLine = jokes1.Punchline
		info.JokeSetup = jokes1.Setup
	}
	return info

}

func getDate() string {
	// Get date in format Mon 01
	dt := time.Now()
	date := dt.Format("2006-Jan-02")

	mon := date[5:8]
	day := date[9:11]
	concatenated := mon + " " + day

	return concatenated
}

func checkBirthday() {
	date := getDate()
	secret := goDotEnvVariable("REDIS")

	client, err := redis.Dial("tcp", "10.10.10.1:6379", redis.DialDatabase(6), redis.DialPassword(secret))
	if err != nil {
		log.Fatal(err)
	}
	hash, _ := redis.StringMap(client.Do("HGETALL", "bday"))
	for key, value := range hash {
		if key[len(key)-6:] == date {
			sendEmail(key[:len(key)-7], value)
		}
	}
}

func sendEmail(name, email string) {
	var info emailer.Birthday
	info = getJoke(info)
	info.Name = name
	info.Email = email
	emailer.SendBdayEmail(info)
}

func main() {
	checkBirthday()
	// Do a job at a specific time - 'hour:min:sec' - seconds optional
	// gocron.Every(1).Day().At("01:15").Do(birthday)

	// NextRun gets the next running time
	// _, time := gocron.NextRun()
	// fmt.Println(time)

	// Start all the pending jobs
	// <-gocron.Start()
}
