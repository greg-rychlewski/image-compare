package csvutil

import (
	"github.com/greg-rychlewski/image-compare/imageutil"
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

// Add new columns to csv and write to output file
func Process(inputPath string, outputPath string, headerIncluded bool) (int, error) {
	// Open input file
        inputFile, err := os.Open(inputPath)

        if err != nil {
                return 0, err
        }

        defer inputFile.Close()

        // Create output file
        outputFile, err := os.OpenFile(outputPath, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0644)

        if err != nil {
                return 0, err
        }

        defer outputFile.Close()

	// Create csv reader and writer
        csvReader := csv.NewReader(inputFile)
        csvWriter := csv.NewWriter(outputFile)

	// Writer header to output csv
	csvWriter.Write([]string{"image1", "image2", "similarity", "elapsed"})
	csvWriter.Flush()

	// Process input csv one line at a time
	isFirstRow := true
	numProcessedPairs := 0

        for {
		// Read row from csv
                row, err := csvReader.Read()

                if err == io.EOF {
                        break
                }

		if err != nil {
			return 0, err
		}

		// If current row is header, skip it
		if isFirstRow && headerIncluded {
			isFirstRow = false
			continue
		}

		// Decode images
		image1, err := imageutil.DecodeImage(row[0])

		if err != nil {
			return 0, err
		}

		image2, err := imageutil.DecodeImage(row[1])

		if err != nil {
			return 0, err
		}

		// Calculate mse along with the time the computation took
		mse, elapsedTime, err := imageutil.MeanSquaredError(image1, image2)

		if err != nil {
			return 0, err
		}

                row = []string{row[0], row[1], strconv.FormatFloat(mse, 'f', -1, 64), strconv.FormatFloat(elapsedTime, 'f', 4, 64)}


		// Write new row to csv
                csvWriter.Write(row)
                csvWriter.Flush()

		numProcessedPairs++

		if isFirstRow {
			isFirstRow = false
		}
        }

	return numProcessedPairs, nil
}
