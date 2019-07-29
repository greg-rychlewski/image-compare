package csvutil

import (
	"encoding/csv"
	"io"
	"os"
)

// Add new columns to csv and write to output file

func Process(inputFile *os.File, outputFile *os.File) {
        csvReader := csv.NewReader(inputFile)
        csvWriter := csv.NewWriter(outputFile)

        isHeader := true

        for {
                row, err := csvReader.Read()

                if err == io.EOF {
                        break
                }

                if !isHeader {
                        row = append(row, "0.1", "0.001")
                } else {
                        row = append(row, "similar", "elapsed")
                }

                csvWriter.Write(row)
                csvWriter.Flush()

                if isHeader {
                        isHeader = false
                }
        }
}
