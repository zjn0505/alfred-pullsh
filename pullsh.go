package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"

	aw "github.com/deanishe/awgo"
)

var (
	wf     *aw.Workflow
	query  string
	isPush bool
	cache  *aw.Cache
)

func init() {
	wf = aw.New()
	dir, err := ioutil.TempDir("", "alfred-pullsh")
	if err != nil {
		log.Panic("Failed to create cache dir")
		panic(err)
	}
	defer os.RemoveAll(dir)
	cache = aw.NewCache(dir)
}

func run() {
	wf.Args() // call to handle any magic actions
	flag.Parse()

	if args := flag.Args(); len(args) > 1 {
		query = args[1]
		isPush = args[0] == "push"
		log.Printf(args[0] + " " + args[1])
	}
	if !isPush {
		// Pull
		if query != "" && len(query) >= 4 && len(query) <= 5 {

			data := pullMemo(query).Memo

			createItem(data)
			wf.SendFeedback()
		}
	} else {
		// Push
		if query != "" && len(query) >= 1 {

			data := pushMemo(query)
			av := aw.NewArgVars()
			av.Arg(data.Memo.ID)
			av.Var("memoId", data.Memo.ID).
				Var("link", "https://pullsh.me/"+data.Memo.ID)
			err := av.Send()

			if err != nil {
				log.Panic("Error", err)
			}
		}
	}
}

func main() {
	wf.Run(run)
}
