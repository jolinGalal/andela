package csvfile

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"github.com/jolinGalal/andela/internal/model"
)

type CSV struct{}

func NewCSV() CSVI {
	return &CSV{}
}

type CSVI interface {
	CreateCSV(res *map[int]model.Result)
}

// CreateCSV ...
func (c *CSV) CreateCSV(res *map[int]model.Result) {
	csvFile, err := os.Create("comments.csv")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer csvFile.Close()
	csvwriter := csv.NewWriter(csvFile)
	csvwriter.Write(header())
	for _, resRow := range *res {
		var row []string
		row = append(row, strconv.Itoa(resRow.UserID))
		row = append(row, strconv.Itoa(resRow.ID))
		row = append(row, resRow.Title)
		row = append(row, resRow.Body)
		row = append(row, resRow.Comments)
		csvwriter.Write(row)
	}
	csvwriter.Flush()

}

// header func to retrun header to CSV
func header() []string {
	var row = []string{
		"userID", "ID", "Title", "Body", "Comments"}
	return row

}
