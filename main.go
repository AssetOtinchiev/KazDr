package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strings"
	"sync"
	"time"
)

func get_resp_time(url string, corePosition int) {
	time_start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		//log.Printf("Error fetching: %v", err)
		//log.Print(http.StatusInternalServerError,"; ",url, "; 0 ms")
		fmt.Println(http.StatusInternalServerError,"; ","0",url, "; 0 ms")
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error fetching: %v", err)
	}

	fmt.Print(resp.Status)
	fmt.Print("; ",len(body),"; ")
	fmt.Print(time.Since(time_start),"ms; ", url)
	fmt.Println(" ядра ",corePosition)

	defer resp.Body.Close()
	fmt.Println()
}

func main() {
	time_start := time.Now()
	content, _ := ioutil.ReadFile("urls.txt")
	urls := strings.Split(string(content), "\n")

	cores := runtime.NumCPU()
	runtime.GOMAXPROCS(cores)

	waitGrp := new(sync.WaitGroup)

	inputAsync := make(chan string, cores)

	for i := 0; i < cores; i++ {
		waitGrp.Add(1)
		go func() {
			defer waitGrp.Done()
			for url := range inputAsync {
				get_resp_time(url, i)
			}
		}()
	}

	for _, url := range urls {
		if url != "" {
			inputAsync <- url
		}
	}

	close(inputAsync)
	waitGrp.Wait()

	fmt.Println(time.Since(time_start), "End")
	var input string
	fmt.Scanln(&input)

}










//input := os.Args[1:]
//
//if len(input) == 0{
//	fmt.Println("end")
//	return
//}
//
//cores := runtime.NumCPU()
//
//fmt.Printf("This machine has %d CPU cores. \n", cores)
//
//// maximize CPU usage for maximum performance
//runtime.GOMAXPROCS(cores)
//
//for _,link := range input{
//	resp, err := http.Get(link)
//	if err!= nil{
//		fmt.Println("Fail")
//		log.Fatalln(err)
//		return
//	}
//
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		log.Fatalln(err)
//	}
//	log.Println(string(body))
//}
//
//fmt.Println(input)

