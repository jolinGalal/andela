package main

import (
	"fmt"
	"log"

	"github.com/jolinGalal/andela/internal/api"
	"github.com/jolinGalal/andela/internal/csvfile"
	"github.com/jolinGalal/andela/internal/model"
)

// main ...
func main() {
	var a = api.NewAPI()
	postsLst, err := a.GetPosts()
	if err != nil {
		log.Fatal(err)
	}
	var ResltMap = make(map[int]model.Result)

	for _, val := range *postsLst {
		res := model.Result{
			UserID:   val.UserID,
			Title:    val.Title,
			ID:       val.ID,
			Body:     val.Body,
			Comments: "",
		}
		ResltMap[val.ID] = res
	}

	commentsLst, err := a.GetComments()
	if err != nil {
		log.Fatal(err)
	}

	for _, val := range *commentsLst {
		res, ok := ResltMap[val.PostID]
		if !ok {
			log.Fatal(fmt.Errorf("can't find this post ID %d", val.PostID))
		}
		if res.Comments == "" {
			res.Comments = val.Body
		} else {
			res.Comments = fmt.Sprintf("%s|%s", res.Comments, val.Body)
		}
		ResltMap[val.PostID] = res
	}

	var c = csvfile.NewCSV()
	c.CreateCSV(&ResltMap)
}
