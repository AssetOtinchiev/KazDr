package Handler

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func SetupCloseHandler(totalThreads map[int]int) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c

		for key, value := range totalThreads{
			fmt.Println("Порядковый номер параллельного потока запроса: ",key," Число запросов: ",value)
		}
		os.Exit(0)
	}()
}
