package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Post struct {
	UserId, Id  int
	Title, Body string
}

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var posts []Post
	err = json.Unmarshal(body, &posts)
	if err != nil {
		panic(err)
	}

	for _, post := range posts {
		fmt.Printf("UserId: %v\nId: %v\nTitle: %v\nBody: %v\n\n", post.UserId, post.Id, post.Title, post.Body)
	}
}
