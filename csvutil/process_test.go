package csvutil

import (
    "os"
    "testing"
)

func TestProcess(t *testing.T) {
    validInputCSV := "../_testdata/unit_test.csv"
    validOutputCSV := "../out.csv"

    // Test #1: Empty input path should produce an error
    _, err := Process("", validOutputCSV, true)
    os.Remove(validOutputCSV)

    if err == nil {
        t.Error("Empty input path did not produce an error.")
    }

    // Test #2: Empty output path should produce an error
    _, err = Process(validInputCSV, "", true)

    if err == nil {
        t.Error("Empty output path did not produce an error.")
    }

    // Test #3: Non-existing input file should produce an error
    _, err = Process("../_testdata/test_small.csvs", "validOutputCSV", true)
    os.Remove(validOutputCSV)

    if err == nil {
        t.Error("Non-existing input file did not produce an error.", true)
    }

    // Test #4: Non-existing output directory should produce an error
    _, err = Process(validInputCSV, "/not/a/real/directory", true)

    if err == nil {
        t.Error("Non-existing output directory did not produce an error.")
    }   

    // Test #5: Non-csv input file should produce an error
    _, err = Process("../main.go", validOutputCSV, true)
    os.Remove(validOutputCSV)

    if err == nil {
        t.Error("Non-csv input file did not produce an error")
    }

    // Test #6: Valid parameters shouldn't produce an error + the number of rows in the input csv should be returned
    numRows, err := Process(validInputCSV, validOutputCSV, true)
    os.Remove(validOutputCSV)

    if numRows != 6 {
        t.Error("Wrong number of csv rows are produced")
    }

    if err != nil {
        t.Error("Error was produced even though file was processed without issues")
    }
}