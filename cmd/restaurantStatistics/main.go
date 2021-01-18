package main

import (
	wrapper "TweetExtractor/internal/APIwrapper"
	"TweetExtractor/pkg/ESpkg"
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

type Result struct {
	Name  string
	Count map[string]int
	Total int
	VN    float32
	Exp   float64
	Log   float64
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
		exp := 0.0
		// log := 0.0
		log := 1 //test
		query := "{\"query\":{\"bool\":{\"should\":[{\"match_phrase\":{\"text\":\"at " + e + " in 高松市 \"}},{\"match_phrase\":{\"text\":\"at " + e + " in Takamatsu\"}}]}}}"
		m, _ := wrapper.GetRestaurantStatistics(es, "restaurant", query)
		for _, v := range m {
			total += v //totalだけ先に計算
		}
		for _, v := range m {
			exp += math.Pow(3.85, float64(v-1)) - 1
			// log += math.Log10(float64(v))
			log *= v
		}

		if len(m) >= 1 { //FIXME 一時的に足切りしてみているが考察が必要
			result = append(result, Result{Name: e, Count: m, Total: total, VN: (float32)(len(m)) / (float32)(total), Exp: math.Log10(exp/float64(len(m)) + 1), Log: float64(log) / float64(total)})
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Log > result[j].Log
	})
	uke := 0 //FIXME tmp
	for _, r := range result {
		uke += r.Total
		fmt.Printf("%s : %+v, total:%d, %d users ,来店者新規度:%f, 指数穴場度:%f，対数穴場度:%f\n", r.Name, r.Count, r.Total, len(r.Count), r.VN, r.Exp, r.Log)
	}
	fmt.Printf("平均値%f", float64(uke)/float64(len(result)))
}
