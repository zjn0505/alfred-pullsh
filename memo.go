package main

import (
	"encoding/json"
	"log"
)

// CreateItem - Create Workflow Item from Memo
func CreateItem(memo Memo) {
	title := memo.ID
	subtitle := memo.Msg
	wf.NewItem(title).
		UID(title).
		Autocomplete(title).
		Subtitle(subtitle).
		Arg(subtitle).
		Copytext(subtitle).
		Quicklook(subtitle).
		Valid(true).
		IsFile(false)
}

// ParseResponse - parse response of api into Resp
func ParseResponse(jsonString string) Resp {
	var result Resp

	err := json.Unmarshal([]byte(jsonString), &result)

	if err != nil {
		log.Panic("Error", err)
	}

	return result
}

// Resp - struct for api response
type Resp struct {
	Status int    `json:"result"`
	Msg    string `json:"msg"`
	Memo   Memo   `json:"memo,omitempty"`
}

// Memo - struct for memo
type Memo struct {
	ID  string `json:"_id"`
	Msg string `json:"msg"`
}
