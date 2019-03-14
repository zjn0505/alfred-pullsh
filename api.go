package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const (
	urlPull string = "https://api.jienan.xyz/memo/?memoId=%v"
	urlPush string = "https://api.jienan.xyz/memo"
)

// PullMemo - pull memo by memoId from API
func PullMemo(memoID string) Resp {
	url := fmt.Sprintf(urlPull, memoID)
	log.Println(url)
	resp, err := http.Get(url)

	if err != nil {
		log.Panic("Error", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Panic("Error", err)
	}

	return ParseResponse(string(body))
}

// PushMemo - push memo to API
func PushMemo(content string) Resp {

	resp, err := http.PostForm(urlPush, url.Values{"msg": {query}})
	if err != nil {
		log.Panicln("Error", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Panic("Error", err)
	}

	return ParseResponse(string(body))
}
