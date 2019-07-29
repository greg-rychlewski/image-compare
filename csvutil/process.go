package csvutil

import (
	"github.com/greg-rychlewski/image-compare/imageutil"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Add new columns to csv and write to output file
func Process(inputFile *os.File, outputFile *os.File, headerIncluded bool) error {
	// Create csv reader and writer
        csvReader := csv.NewReader(inputFile)
        csvWriter := csv.NewWriter(outputFile)

	// Create output csv header
	header := [4]string{"image1", "image2", "similar", "elapsed(seconds)"}
	csvWriter.Write(header)
	csvWriter.Flush()

	// Process input csv one line at a time
	rowCount := 0

        for {
		rowCount++

		// Read row from csv
                row, err := csvReader.Read()

                if err == io.EOF {
                        break
                }

		if err != nil {
			return err
		}

		// Skip row if it's the header
		if rowCount == 1 && headerIncluded {
			continue
		}

		// Create new slice instead of appending to current row 
		// Do this in case there is unexpected data in columns 3+
		mse, elapsedTime, err := imageutil.MeanSquaredError(row[0], row[1])

		if err != nil {
			return err
		}

                row = []string{row[0], row[1], strconv.FormatFloat(mse, 'f', -1, 64), strconv.FormatFloat(elapsedTime, 'f', 4, 64)}


		// Write new row to csv
                csvWriter.Write(row)
                csvWriter.Flush()
        }

	fmt.Printf("%d rows successfully processed", rowCount)

	return nil
}
