package imageutil

import "testing"

func TestDecodeImage(t *testing.T) {
    // Test #1: Non-existing image should produce an error
    if _, err := DecodeImage("this_image_file_does_not_exist"); err == nil {
        t.Error("Non-existing image doesn't produce an error")
    }

    // Test#2: File that is not an image should produce an error
    if _, err := DecodeImage("decode.go"); err == nil {
        t.Error("Non-image file doesn't produce an error")
    }

    // Test3: File that is a valid image should not produce an error
    if _, err := DecodeImage("../_testdata/images/federer1.png"); err != nil {
        t.Error("Valid image produces an error")
    }
}
