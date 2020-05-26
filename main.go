package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
	help "./Helpers"
	mdl "./Models"
	hand "./Handler"
)

var cities =[]*mdl.CvcOutput{}
var totalThreads = make(map[int]int)

func main() {
	time_start := time.Now()
	content, _ := ioutil.ReadFile("urls.txt")
	urls := strings.Split(string(content), "\n")

	cores := runtime.NumCPU()
	runtime.GOMAXPROCS(cores)

	waitGrp := new(sync.WaitGroup)

	inputAsync := make(chan string, cores)

	for i := 1; i <= cores; i++ {
		waitGrp.Add(1)
		go func(i2 int) {
			defer waitGrp.Done()

			for url := range inputAsync {
				get_resp_time(url, i2)
			}
		}(i)
	}

	for _, url := range urls {
		if url != "" {
			inputAsync <- url
		}
	}

	close(inputAsync)
	waitGrp.Wait()

	help.SaveToCSV(cities)

	fmt.Println(time.Since(time_start), "End")

	hand.SetupCloseHandler(totalThreads)
	var input string
	fmt.Scanln(&input)
}


func get_resp_time(url string, corePosition int) {
	time_start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(http.StatusInternalServerError,"; ","0",url, "; 0 ms")
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error fetching: %v", err)
	}
		totalThreads[corePosition] += 1

	timeStr := time.Since(time_start).String()

	city:=new(mdl.CvcOutput)
	city.Path = url
	city.StatusCode = strconv.Itoa(resp.StatusCode)
	city.Weight = strconv.Itoa(len(body))
	city.ResponseWait = timeStr[:len(timeStr)-1] + "ms"

	cities=append(cities,city)
	defer resp.Body.Close()
}
