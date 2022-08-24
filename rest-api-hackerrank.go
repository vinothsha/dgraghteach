package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func getTotalGoals(team string, year string) string {

	url := "https://jsonmock.hackerrank.com/api/football_matches?year=" + year + "&team1=" + team + "&page=1"

	resp, _ := http.Get(url)

	// defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var s map[string][]map[string]string
	json.Unmarshal(body, &s)
	var count int
	// var arr []int

	for i := 0; i <= len(s); i++ {
		num, _ := strconv.Atoi(s["data"][i]["team1goals"])
		count += num
	}
	url = "https://jsonmock.hackerrank.com/api/football_matches?year=" + year + "&team2=" + team + "&page=1"

	resp, _ = http.Get(url)

	// defer resp.Body.Close()
	body, _ = ioutil.ReadAll(resp.Body)
	// var s map[string][]map[string]string
	json.Unmarshal(body, &s)

	for i := 0; i <= len(s); i++ {
		num, _ := strconv.Atoi(s["data"][i]["team2goals"])
		count += num
	}

	fmt.Println(count)
	return ""
}

func main() {
	// var team, year string
	// fmt.Scanln(&team)
	// fmt.Scanln(&year)
	getTotalGoals("Barcelona", "2011")
	// fmt.Println(result)
}
