package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jolinGalal/andela/internal/model"
)

// GetComments function to call comments api
func (a *API) GetComments() (*[]model.Comment, error) {

	response, err := http.Get("https://jsonplaceholder.typicode.com/comments")

	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	var commentLst []model.Comment
	err = json.Unmarshal(responseData, &commentLst)
	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}
	return &commentLst, nil
}
