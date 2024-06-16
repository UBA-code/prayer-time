package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) == 3 {
		country := os.Args[1]
		city := os.Args[2]
		fmt.Println("getting data ...")
		fetchData(fmt.Sprintf("http://api.aladhan.com/v1/timingsByCity?city=%v&country=%v", city, country))
	} else {
		fmt.Println("usage: prayTime \"country\" \"city\"")
	}
}

func fetchData(apiUrl string) {
	response, err := http.Get(apiUrl)
	checkError(err)
	defer response.Body.Close()
	if response.StatusCode != 200 {
		panic("request failed: " + response.Status)
	}
	contentBytes, err := io.ReadAll(response.Body)
	checkError(err)
	var data apiResponse
	err = json.Unmarshal(contentBytes, &data)
	checkError(err)

	template := ""

	template += fmt.Sprintln("_______________________")
	template += fmt.Sprintf("|\t%v    |\n|---------------------|\n", data.Data.Date.Gregorian.Date)
	template += fmt.Sprintf("| %-10v:\t%v |\n", "Fajr", data.Data.Timings.Fajr)
	template += fmt.Sprintf("| %-10v:\t%v |\n", "Sunrise", data.Data.Timings.Sunrise)
	template += fmt.Sprintf("| %-10v:\t%v |\n", "Dhuhr", data.Data.Timings.Dhuhr)
	template += fmt.Sprintf("| %-10v:\t%v |\n", "Asr", data.Data.Timings.Asr)
	template += fmt.Sprintf("| %-10v:\t%v |\n", "Maghrib", data.Data.Timings.Maghrib)
	template += fmt.Sprintf("| %-10v:\t%v |\n", "Isha", data.Data.Timings.Isha)
	template += fmt.Sprintln("-----------------------")

	fmt.Println(template)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
