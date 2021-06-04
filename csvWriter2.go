package main

import (
"encoding/csv"
"fmt"
"os"
)


func main() {
	var rows [][]string
	numberOfVisitors := 100


	for i := 1; i <= numberOfVisitors; i++ {
		rows = append(rows, []string{fmt.Sprintf("visitor%d", i)})
	}

	if csvfile, err := os.Create(fmt.Sprintf("%dvisitor.csv", numberOfVisitors)); err != nil {
		println(fmt.Sprintf("failed creating file: %s", err))
	} else {

		csvwriter := csv.NewWriter(csvfile)

		for _, row := range rows {
			_ = csvwriter.Write(row)
		}

		csvwriter.Flush()

		if err := csvfile.Close(); err != nil{
			println("womp")
		}
	}
}

