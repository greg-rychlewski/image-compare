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

You should see a program named `image_compare` in your current directory.

Developers wishing to build from source must install Go on their computer and use the _go_ _get_ command
  
    go get github.com/greg-rychlewski/image-compare
  
## Usage

The program expects the user to supply a csv file with two columns. Each column should be populated with paths to image files. 

The paths must either be absolute or relative to the directory you are running the program under. Using absolute paths is preferable in case you ever decide to move the location of the program.

To start the program, move to the directory containing the program and use the following command

MacOS/Git Bash for Windows

    ./image_compare -in path_to_csv_file
    
 Windows CMD
 
    .\image_compare -in path_to_csv_file
 
If you add the location of the program to your `PATH` environment variable, you can run the program from anywhere using just `image_compare`. 
