package csvutil

import (
	"github.com/greg-rychlewski/image-compare/imageutil"
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

// Add new columns to csv and write to output file

func Process(inputFile *os.File, outputFile *os.File) error {
        csvReader := csv.NewReader(inputFile)
        csvWriter := csv.NewWriter(outputFile)

        isHeader := true

        for {
                row, err := csvReader.Read()

                if err == io.EOF {
                        break
                }

                if !isHeader {
			mse, elapsedTime, err := imageutil.MeanSquaredError(row[0], row[1])

			if err != nil {
				return err
			}


                        row = []string{row[0], row[1], strconv.FormatFloat(mse, 'f', -1, 64), strconv.FormatFloat(elapsedTime, 'f', 4, 64)}
                } else {
                        row = []string{row[0], row[1], "similar", "elapsed (seconds)"}
                }

                csvWriter.Write(row)
                csvWriter.Flush()

                if isHeader {
                        isHeader = false
                }
        }

	return nil
}
