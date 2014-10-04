package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Repository struct {
	Type         string
	Username     string
	Name         string
	Owner        string
	Homepage     string
	Description  string
	Language     string
	Watchers     int
	Followers    int
	Forks        int
	Size         int
	OpenIssues   int
	Score        float32
	HasDownloads bool
	HasIssues    bool
	HasWiki      bool
	Fork         bool
	Private      bool
	Url          string
	Created      string
	CreatedAt    string
	PushedAt     string
	Pushed       string
}

func (r *Repository) String() string {
	return fmt.Sprintf("(%d)\t%s/%s - %s\n", r.Followers, r.Owner, r.Name, r.Description)
}

func main() {
	url := "https://api.github.com/legacy/repos/search/go?language=Go&sort=stars"
	res, _ := http.Get(url)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	var data map[string][]*Repository

	json.Unmarshal(body, &data)

	for i, v := range data["repositories"] {
		fmt.Printf("%s", v)
		if i == 9 {
			break
		}
	}
}
