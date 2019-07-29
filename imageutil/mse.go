package imageutil

import (
	"github.com/nfnt/resize"
	"errors"
	"math"
	"os"
	"strings"
	"time"
	"image"
	_ "image/gif"
	_ "image/png"
	_ "image/jpeg"
)


func MeanSquaredError(imagePath1 string, imagePath2 string) (float64, float64, error) {
	// Get start time so that the elapsed time for the function run can be returned
	start := time.Now()

	// Open images and decode them
	image1, err := os.Open(strings.TrimSpace(imagePath1))

	if err != nil {
		return 0.0, 0.0, err
	}

	defer image1.Close()

	imageData1, _, err := image.Decode(image1)

	if err != nil {
		return 0.0, 0.0, err
	}

	image2, err := os.Open(strings.TrimSpace(imagePath2))

	if err != nil {
		return 0.0, 0.0, err
	}

	defer image2.Close()

	imageData2, _, err := image.Decode(image2)

	if err != nil {
		return 0.0, 0.0, err
	}


	// Resize images so that they are equal
	imageData1, imageData2 = makeImageSizesEqual(imageData1, imageData2)

        // Calculate sum of squared errors
	sse, err := sumSquaredErrors(imageData1, imageData2)

	if err != nil {
		return 0.0, 0.0, err
	}

	mse := sse / (4.0 * float64(imageData1.Bounds().Max.X * imageData2.Bounds().Max.Y))


	return mse, time.Since(start).Seconds(), nil
}

func makeImageSizesEqual(image1 image.Image, image2 image.Image) (image.Image, image.Image) {
	// If sizes are equal then return original images
	if image1.Bounds().Max == image2.Bounds().Max {
		return image1, image2
	}

	// If sizes are not equal then shrink
	minWidth := uint(math.Min(float64(image1.Bounds().Max.X), float64(image2.Bounds().Max.X)))
	minHeight := uint(math.Min(float64(image1.Bounds().Max.Y), float64(image2.Bounds().Max.Y)))

	image1 = resize.Resize(minWidth, minHeight, image1, resize.Lanczos3)
	image2 = resize.Resize(minWidth, minHeight, image2, resize.Lanczos3)

	return image1, image2
}

func sumSquaredErrors(image1 image.Image, image2 image.Image) (float64, error) {
	// Throw error if image sizes are not equal
	if image1.Bounds().Max != image2.Bounds().Max {
		return 0.0, errors.New("Image sizes are unequal. Cannot calculate sum of squared errors")
	}

	// Loop through pixels and calculate squared distance between RGBA values
	result := 0.0

	for y := 0; y < image1.Bounds().Max.Y; y++ {
                 for x := 0; x < image1.Bounds().Max.X; x++ {
                        r1, g1, b1, a1 := image1.At(x, y).RGBA()
			r2, g2, b2, a2 := image2.At(x, y).RGBA()

			result += squaredDistanceNormalizeRGBA([4]uint32{r1, g1, b1, a1}, [4]uint32{r2, g2, b2, a2})
                 }
        }

	return result, nil
}

func squaredDistanceNormalizeRGBA(point1 [4]uint32, point2 [4]uint32) float64 {
	// Calculate squared Euclidean distance between RGBA points
	// Normalize the points between 0 and 1 before calculating distance to avoid overflow
	maxValueRGBA := 65535.0

	result := 0.0

	for i := 0; i < 4; i++ {
		if point1[i] < point2[i] {
			result += math.Pow(float64(point2[i] - point1[i]) / maxValueRGBA, 2)
		} else {
			result += math.Pow(float64(point1[i] - point2[i]) / maxValueRGBA, 2)
		}
	}

	return result
}
