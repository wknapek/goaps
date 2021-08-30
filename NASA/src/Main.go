package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type outputJson struct {
	Error string   `json:"error"`
	Urls  []string `json:"urls"`
}

func main() {
	key := flag.String("api_key", "DEMO_KEY", "api_key for NASA API")
	port := flag.String("port", "8080", "port for server")
	que := flag.Int("req_limit", 5, "limit of concurent request")
	flag.Parse()
	http.HandleFunc("/pictures", func(w http.ResponseWriter, r *http.Request) {
		httpHandler(w, r, *key, *que)
	})
	log.Println("Start listen on localHost:", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}

func isEndDate(date time.Time, endDate time.Time) bool {
	if date.Year() == endDate.Year() && date.Month() == endDate.Month() && date.Day() == endDate.Day() {
		return false
	}
	return true
}

func httpHandler(w http.ResponseWriter, r *http.Request, key string, qu int) {
	output := outputJson{}
	startDate, err := time.Parse("2006-01-02", r.URL.Query().Get("start_date"))
	if err != nil {
		log.Fatalln(err)
	}
	endDate, err := time.Parse("2006-01-02", r.URL.Query().Get("end_date"))
	if checkDate(startDate, endDate) {
		output.Error = "end date error"
	}
	if err != nil {
		log.Fatalln(err)
	}
	days := make([]time.Time, 1)
	for day := startDate; isEndDate(day, endDate); day = day.AddDate(0, 0, 1) {
		days = append(days, day)
	}
	taskSplitter(qu, key, &output, days)
	json.NewEncoder(w).Encode(output)
}

func checkDate(startDate time.Time, endDate time.Time) bool {
	if startDate.Before(endDate) {
		return false
	}
	if endDate.Before(time.Now()) {
		return false
	}
	return true
}

func getFromNasa(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return body, nil
}

func preparePictures(ch chan<- string, key string, days []time.Time) {
	for _, day := range days {
		body, err := getFromNasa("https://api.nasa.gov/planetary/apod?api_key=" + key + "&date=" + day.Format("2006-01-02"))
		if err != nil {
			log.Fatalln(err)
			ch <- "error:" + err.Error()
		}
		var pictures map[string]interface{}
		if err := json.Unmarshal(body, &pictures); err != nil {
			log.Fatalln(err)
			ch <- "error:" + err.Error()
		}
		if pictures["hdurl"] == nil {
			if pictures["url"] == nil {
				ch <- "error: no picture"
			}
			ch <- string(fmt.Sprintf("%s", pictures["url"]))
		} else {
			ch <- string(fmt.Sprintf("%s", pictures["hdurl"]))
		}
	}
}

func taskSplitter(tasks int, key string, output *outputJson, days []time.Time) {
	ch := make(chan string, tasks)
	setURLs := len(days) / tasks
	beginOfSet := 0
	for endOfSet := setURLs; endOfSet < len(days); endOfSet += setURLs {
		go preparePictures(ch, key, days[beginOfSet:endOfSet])
		beginOfSet = endOfSet
	}
	chtext := <-ch
	fmt.Println(chtext)
	if !strings.Contains(chtext, "https://apod.nasa.gov/") {
		output.Error += chtext
	} else {
		output.Urls = append(output.Urls, chtext)
	}
}
