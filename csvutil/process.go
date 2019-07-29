package csvutil

import (
	"github.com/greg-rychlewski/image-compare/imageutil"
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"fmt"
)

// Add new columns to csv and write to output file
func Process(inputFile *os.File, outputFile *os.File, headerIncluded bool) error {
	// Create csv reader and writer
        csvReader := csv.NewReader(inputFile)
        csvWriter := csv.NewWriter(outputFile)

	// Loop through csv file one line at a time
        isFirstRow := true

        for {
		// Read row from csv
                row, err := csvReader.Read()

                if err == io.EOF {
                        break
                }

		if err != nil {
			return err
		}

		// Create new slice instead of appending to current row 
		// Do this in case unexpected data is in columns 3+
		if isFirstRow && headerIncluded {
			row = []string{row[0], row[1], "similar", "elapsed (seconds)"}
		} else if isFirstRow {
			row = []string{"image1", "image2", "similar", "elapsed (seconds"}
		} else{
			mse, elapsedTime, err := imageutil.MeanSquaredError(row[0], row[1])

			if err != nil {
				return err
			}

                        row = []string{row[0], row[1], strconv.FormatFloat(mse, 'f', -1, 64), strconv.FormatFloat(elapsedTime, 'f', 4, 64)}
		}

		// Write new row to csv
                csvWriter.Write(row)
                csvWriter.Flush()

                if isFirstRow {
                        isFirstRow = false
                }
        }

	return nil
}
