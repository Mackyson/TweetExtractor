package main

import (
	wrapper "TweetExtractor/internal/APIwrapper"
	"TweetExtractor/pkg/ESpkg"
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Result struct {
	Name  string
	Count map[string]int
	Total int
}

var result = make([]Result, 0)

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	tmp := stdin.Text()
	list := strings.Split(tmp, ",")
	es, _ := ESpkg.GetDBClient()

	for _, e := range list {
		total := 0
		// query := "{\"query\":{\"bool\":{\"should\":[{\"match_phrase\":{\"text\":\"at " + spotName + " in 高松市 \"}},{\"match_phrase\":{\"text\":\"at " + spotName + " in Takamatsu\"}}]}}}"
		query := "{\"query\": {\"term\": {\"user.id_str\":\"" + e + "\"}}}"
		m, _ := wrapper.GetAllStatistics(es, "restaurant", query)
		for _, v := range m {
			total += v
		}
		result = append(result, Result{Name: e, Count: m, Total: total})
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Total > result[j].Total //降順ソート
	})
	for _, r := range result {
		fmt.Printf("%s : %+v, total:%d\n", r.Name, r.Count, r.Total)
	}
}
