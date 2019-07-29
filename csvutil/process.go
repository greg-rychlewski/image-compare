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
	// Create csv reader and writer
        csvReader := csv.NewReader(inputFile)
        csvWriter := csv.NewWriter(outputFile)

	// Loop through csv file one line at a time
	// Assume there is a header row and process it differently
        isHeader := true

        for {
		// Read row from csv
                row, err := csvReader.Read()

                if err == io.EOF {
                        break
                }

		if err != nil {
			return err
		}

		// If first row, add new columns names
		// Otherwise, calculate mse and time taken for mse calculation
		// Create new slice instead of appending to current row
		// Do this in case there is unexpected data in columns 3+
                if !isHeader {
			mse, elapsedTime, err := imageutil.MeanSquaredError(row[0], row[1])

			if err != nil {
				return err
			}

                        row = []string{row[0], row[1], strconv.FormatFloat(mse, 'f', -1, 64), strconv.FormatFloat(elapsedTime, 'f', 4, 64)}
                } else {
                        row = []string{row[0], row[1], "similar", "elapsed (seconds)"}
                }

		// Write new row to csv
                csvWriter.Write(row)
                csvWriter.Flush()

                if isHeader {
                        isHeader = false
                }
        }

	return nil
}
