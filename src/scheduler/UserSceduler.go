package scheduler

import (
	"github.com/go-resty/resty/v2"
	"github.com/jasonlvhit/gocron"
	"log"
)

func SyncUser() {
	scheduler := gocron.NewScheduler()
	log.Println("start sync user")
	scheduler.Every(5).Seconds().Do(fetchUserInfo)
	<-scheduler.Start()
}

type DemoResult struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

const URL = "https://jsonplaceholder.typicode.com/todos/1"

func fetchUserInfo() {
	client := resty.New()

	res := new(DemoResult)

	response, err := client.R().SetHeader("Content-Type", "application/json").SetResult(res).Get(URL)
	if err != nil {
		log.Println("fetch user info error")
	}
	if response.IsSuccess() {
		log.Println("UserInfo=%v", res)
	}
}
