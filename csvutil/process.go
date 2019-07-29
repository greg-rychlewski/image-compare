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
			mse, err := imageutil.MeanSquaredError(row[0], row[1])

			if err != nil {
				return err
			}


                        row = append(row, strconv.FormatFloat(mse, 'f', -1, 64), "0.001")
                } else {
                        row = append(row, "similar", "elapsed")
                }

                csvWriter.Write(row)
                csvWriter.Flush()

                if isHeader {
                        isHeader = false
                }
        }

	return nil
}
