package helpers

import (
	"encoding/csv"
	"log"
	"os"
	mdl "../Models"
)

func SaveToCSV(cities []*mdl.CvcOutput) bool{
	file, err := os.Create("result.csv")
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range  cities {
		result := []string{value.Path,value.StatusCode,value.Weight,value.ResponseWait}
		err := writer.Write(result)
		if err  != nil{
			return false
		}
		checkError("Cannot write to file", err)
	}
	return true
}



func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
