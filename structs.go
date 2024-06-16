package main

type date struct {
	Today     string `json:"readable"`
	Gregorian struct {
		Date string `json:"date"`
	} `json:"gregorian"`
}

type timings struct {
	Fajr    string `json:"Fajr"`
	Sunrise string `json:"Sunrise"`
	Dhuhr   string `json:"Dhuhr"`
	Asr     string `json:"Asr"`
	Maghrib string `json:"Maghrib"`
	Isha    string `json:"Isha"`
}

type data struct {
	Timings timings `json:"timings"`
	Date    date    `json:"date"`
}

type apiResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   data   `json:"data"`
}
