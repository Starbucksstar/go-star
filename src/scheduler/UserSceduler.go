package scheduler

import (
	"github.com/jasonlvhit/gocron"
	"log"
	"star/src/client"
)

func SyncUser() {
	scheduler := gocron.NewScheduler()
	log.Println("start sync user")
	err := scheduler.Every(5).Seconds().Do(fetchUserInfo)
	if err != nil {
		log.Println("sync user error info=", err.Error())
		return
	}
	<-scheduler.Start()
}

const URL = "https://jsonplaceholder.typicode.com/todos/1"

func fetchUserInfo() {
	httpClient := client.NewHttpClient()
	result := httpClient.GET(nil, URL)
	log.Println(result)
}
