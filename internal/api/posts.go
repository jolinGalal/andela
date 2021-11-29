package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jolinGalal/andela/internal/model"
)

// GetPosts funcition to get posts from the api
func (a *API) GetPosts() (*[]model.Post, error) {

	response, err := http.Get("https://jsonplaceholder.typicode.com/posts")

	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	var postLst []model.Post
	err = json.Unmarshal(responseData, &postLst)
	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}
	return &postLst, nil
}
