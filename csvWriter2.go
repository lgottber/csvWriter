package main

import (
"encoding/csv"
"fmt"
"os"
)


func main() {
	var rows [][]string


	for i := 1; i <= 10000000; i++ {
		rows = append(rows, []string{fmt.Sprintf("visitor%d", i)})
	}

	if csvfile, err := os.Create("10000000visitor.csv"); err != nil {
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

