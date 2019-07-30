package imageutil

import (
    _ "image/png"
    _ "image/gif"
    _ "image/jpeg"
    "image"
    "os"
    "strings"
)

// Return decoded image data from file path
func DecodeImage(imagePath string) (image.Image, error) {
    imageFile, err := os.Open(strings.TrimSpace(imagePath))

    if err != nil {
        return nil, err
    }

    defer imageFile.Close()

    image, _, err := image.Decode(imageFile)

    if err != nil {
        return nil, err
    }

    return image, nil
}
