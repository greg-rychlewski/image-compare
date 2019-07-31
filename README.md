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

Developers wishing to build from source must install Go on their computer and use the `go get` command
  
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
  - Included so user can see if they have the latest release. Added other information for developers (git hash, go version, build time) in case they are asked to help the user when something goes wrong.

#### CSV Processing

- I chose to read the csv file one line at a time instead of reading all of it into memory at once. This avoids issues when the file is larger than the amount of available RAM.
- I created a custom error type for csv-related errors. This is so I can add the csv line number to the error message. It makes it easier for the user to investigate issues.
- When creating an output row, I don't append the results to the entire input row. It's possible the input csv could have data in columns 3+ even though they're supposed to be empty. I take only the first 2 columns and append the results to those.
- I exit the entire program if a single line has an unrecoverable error. I could have skipped over problematic lines and continued to process the good ones. I chose this way because the requirements state that the output should be in the same order as the input. This means the user is going to have to rerun the entire file again and there's no point in continuing.

#### Image Processing

- I chose mean squared error for the similarity calculation. My reasons for choosing it are because it's quick, simple and guaranteed to give 0 when the images are equal. From the requirements, it sounds like we are more concerned with getting decent scores quickly than taking a long time to get perfect scores.
- If the images have different sizes then I shrink the bigger one so that it is equal to the smaller one. I do this because mse is undefined when one of the images is missing a pixel. I chose to scale rather than crop so that we are not ignoring large sections of an image. I chose to scale down rather than up so that we are looping over less pixels when calculating mse. 
- When calculating mse, I normalize all the rgba values between 0 and 1 to avoid potential overflow. RGBA values in Go are between 0 and 65535 instead of 0 and 255.

#### Build Script

- The build script contains important metadata such as app version, go version, git hash and build time. This is passed into the app using ldflags so that users can access it through a command-line option.
- Unit tests are run before building the app to help ensure code quality.
- The app is built for 32/64-bit Windows and MacOS. These are saved as .tar files for 2 reasons. The first is so I don't have to include os/architecture information in the executable names. That would make it annoying for the users to call them. The second is so the MacOS files retain their executable permissions. If the bare file is downloaded, it will have permissions 644. Archiving it with tar allows the permission metadata to be saved.

#### Dependencies

- I took care to use as little dependencies outside of the Go standard libary as possible. In general, it's harder to make sure those kinds of libraries will be updated and remain compatible as Go evolves. I only used one external library for image resizing.
- I use Go modules to make sure the information for the one external libary I used is available to other developers.

#### Continuous Integration

- I use Travis CI to run my build script whenever I push to GitHub. I also use it to automatically deploy a release whenever I push a tagged commit.

#### Code Structure

- My philosophy for the code structure was to be as modular as possible and have the main function use those modules. This makes the code more organized and less coupled. This is a pretty small project, but I feel like it's a good idea to start this way from the beginning to avoid a lot of refactoring when the project gets bigger.

#### Opportunities for Improvement

- MSE is calculated by looping through the pixels one at a time. I would like to see if Go has a way to vectorize these operations, like Python's numpy. I've read about the gonum package and would be interested in seeing if it can be used with image data.
- This might be overkill, but if the number of images being processed gets really big it might be fun to see if multi-threading speeds things up. If I was going to try this, I would set it up in the following way:
  - One worker reading the csv lines and putting them into a buffer
  - Several workers reading from that buffer, processing the images and putting the output into another buffer
  - One worker reading from the output buffer and writing the results to the output csv
