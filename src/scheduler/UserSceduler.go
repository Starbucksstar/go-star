package scheduler

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
	"io"
	"log"
	"net/http"
	"star/src/inject"
)

func SyncUser() {
	scheduler := gocron.NewScheduler()
	log.Println("start sync user")
	_ = scheduler.Every(5).Minutes().Do(fetchUserInfo())
	<-scheduler.Start()
}

func fetchUserInfo() func() {
	return func() {
		apiURL := "https://jsonplaceholder.typicode.com/todos/1"
		response, err := http.Get(apiURL)
		if err != nil {
			fmt.Println("HTTP请求发生错误:", err)
			return
		}
		defer response.Body.Close()

		body, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println("读取响应内容发生错误:", err)
			return
		}
		log.Println("响应内容:", string(body))

		userService := inject.InitUserService()
		users, err := userService.FindAllUser(1, 10)
		if err != nil {
			log.Println(err)
		}
		log.Println(users)
	}
}
