package csvutil

import (
    "os"
    "testing"
)

func TestProcess(t *testing.T) {
    realInputCSV := "../_testdata/unit_test.csv"
    realOutputCSV := "../out.csv"

    // Test #1: Empty input path should produce an error
    _, err := Process("", realOutputCSV, true)
    os.Remove(realOutputCSV)

    if err == nil {
        t.Error("Empty input path did not produce an error.")
    }

    // Test #2: Empty output path should produce an error
    _, err = Process(realInputCSV, "", true)

    if err == nil {
        t.Error("Empty output path did not produce an error.")
    }

    // Test #3: Non-existing input file should produce an error
    _, err = Process("../_testdata/test_small.csvs", "realOutputCSV", true)
    os.Remove(realOutputCSV)

    if err == nil {
        t.Error("Non-existing input file did not produce an error.", true)
    }

    // Test #4: Non-existing output directory should produce an error
    _, err = Process(realInputCSV, "/not/a/real/directory", true)

    if err == nil {
        t.Error("Non-existing output directory did not produce an error.")
    }   

    // Test #5: Non-csv input file should produce an error
    _, err = Process("../main.go", realOutputCSV, true)
    os.Remove(realOutputCSV)

    if err == nil {
        t.Error("Non-csv input file did not produce an error")
    }

    // Test #6: Valid parameters shouldn't produce an error + the number of rows in the input csv should be returned
    numRows, err := Process(realInputCSV, realOutputCSV, true)
    os.Remove(realOutputCSV)

    if numRows != 6 {
        t.Error("Wrong number of csv rows are produced")
    }

    if err != nil {
        t.Error("Error was produced even though file was processed without issues")
    }
}