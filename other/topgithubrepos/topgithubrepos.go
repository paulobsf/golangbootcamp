package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Result struct {
	Items []*Repository
}

type Repository struct {
	FullName      string `json:"full_name"`
	Description   string
	Homepage      string
	WatchersCount int `json:"watchers_count"`
	Language      string
}

func (r *Repository) String() string {
	return fmt.Sprintf("(%d)\t%s - %s", r.WatchersCount, r.FullName, r.Description)
}

func GetTopRepos(language string, limit int) (string, int) {
	var (
		result     string
		totalStars int
		data       Result
	)

	url := fmt.Sprintf("https://api.github.com/search/repositories?q=+language:%s&sort=stars&order=desc", language)

	res, _ := http.Get(url)
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	json.Unmarshal(body, &data)

	for i, item := range data.Items {
		result += fmt.Sprintf("%s\n", item)
		totalStars += item.WatchersCount
		if i == limit-1 {
			break //	fugly... should I just use a regular for loop??
		}
	}

	return result, totalStars
}

func main() {
	languages := []string{
		"go",
		"java",
		"scala",
		"clojure",
		"javascript",
		"python",
		"ruby",
	}

	for _, language := range languages {
		repos, totalStars := GetTopRepos(language, 10)
		fmt.Printf("%s (%d)\n--\n", language, totalStars)
		fmt.Println(repos)
	}
}
