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

1. Validate user input
2. Read line from csv file
3. Validate image paths
4. Calculate mean squared error for the 2 images
5. Write results to output csv
6. Repeat 2-5 until you get to the end of the input csv

Note: Program exits early if there are any errors that would prohibit the calculation of a pair of images.

 

