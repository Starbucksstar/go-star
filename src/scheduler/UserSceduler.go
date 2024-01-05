package scheduler

import (
	"github.com/jasonlvhit/gocron"
	"log"
	"star/src/client"
)

func SyncUser() {
	scheduler := gocron.NewScheduler()
	log.Println("start sync user")
	scheduler.Every(5).Seconds().Do(fetchUserInfo)
	<-scheduler.Start()
}

const URL = "https://jsonplaceholder.typicode.com/todos/1"

func fetchUserInfo() {
	httpClient := client.NewHttpClient()
	result := httpClient.GET(nil, URL)
	log.Println(result)
}
