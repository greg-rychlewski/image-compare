# Image Comparer

Calculates the mean squared error for pairs of images.

## Installation

You can download the latest release from GitHub

    https://github.com/greg-rychlewski/image-compare/releases/latest

To be notified of new releases, log into your GitHub account and click the `watch` button at the top of this page.

Extract the contents of the file as follows

32/64-bit Windows

    tar xf image_compare-windows-386.tar / tar xf image_compare-windows-amd64.tar

32/64-bit MacOS (Darwin)

    tar xf image_compare-darwin-386.tar / tar xf image_compare-darwin-amd64.tar

The program will be extracted as `image_compare`.

Developers wishing to build from source must install Go on their computer and use the _go_ _get_ command
  
    go get github.com/greg-rychlewski/image-compare
  
## Usage

You are expected to supply a csv file with two columns. Each column should be populated with paths to image files. 

The paths must either be absolute or relative to your current directory. Absolute paths are preferable in case you run the program from different locations.

Start the program by moving to the directory containing `image_compare` and run the following command

MacOS/Git Bash for Windows

    ./image_compare -in input_csv_file
    
Windows CMD
 
    .\image_compare -in input_csv_file
 
The results will be saved to a file in your current directory. 

You can specify your own output file by adding `-out output_csv_file` to the above commands.

Run `image_compare -h` to see a list of other options you can specify.

## Design 

### Flow Chart

<p align="center"> 
<img src="https://github.com/greg-rychlewski/image-compare/blob/master/_testdata/images/flowchart.png">
</p>

1. Validate user input. The user can specify input/output csv locations.
2. Read line from csv file. This will contain 2 image paths.
3. Validate image paths.
4. Calculate mean squared error for the 2 images.
5. Write results to output csv. This will include the original image paths as well as the mse and the elapsed time for the calculation.
6. Repeat 2-5 until you get to the end of the input csv

Note: Program exits early if there are any errors that would prohibit the calculation of a pair of images.

### Additional Details
 
#### Command-line flags
 
I allow the user to specify the following flags:
- Input csv
  - Required because the program can't do anything without it.
- Output csv
  - Optional for user convenience. Default output is timestamped down to the second to ensure uniqueness.
- No header in csv file
  - Included for user convenience. Their csv file might be missing a header sometimes.
- Version information
  - Included so user can see if they have the latest release. Added other information for developers (git hash, go version, build time) in case they are asked to help the user when something is wrong.

#### CSV Processing

- I chose to read the csv file one line at a time instead of reading all of it into memory at once. This avoids issues when the file is larger than the amount of available RAM.
- When creating an output row, I don't append the results to the entire input row. It's possible the input csv could have data in columns 3+ even though they're supposed to be empty. 
- I exit the entire program if a single line has an unrecoverable error. I could have skipped over problematic lines and continued to process the good ones. I chose this way because the requirements state that the output should be in the same order as the input. This means the user is going to have to rerun the entire file again and there's no point in continuing.

#### Image Processing

- I chose mean squared error for the similarity calculation. My reasons for choosing it are because it's quick, simple and guaranteed to give 0 when 2 images are equal.
