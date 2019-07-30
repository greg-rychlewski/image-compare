package imageutil

import "testing"

func TestMSE(t *testing.T) {
    blackTransparent, _ := DecodeImage("../_testdata/images/0000.png")
    blackOpaque, _ := DecodeImage("../_testdata/images/black.png")
    whiteOpaque, _ := DecodeImage("../_testdata/images/white.png")

    // Test #1: MSE between an image with all (0,0,0,0) rgba and an image with all (255,255,255,255) rgba should be 1
    if mse, _, _ := MeanSquaredError(blackTransparent, whiteOpaque); mse != 1.0 {
        t.Error("MSE between image with all (0,0,0,0) rgba and image with all (255,255,255,255) rgba is not 1")
    }

    // Test #2: MSE between an image and itself should be 0
    if mse, _, _ := MeanSquaredError(whiteOpaque, whiteOpaque); mse != 0.0 {
        t.Error("MSE between an image and itself is not 0")
    }

    // Test #3: R, G, B, A should contribute equally to MSE. The MSE between an opaque white image and opaque black image should be 0.75.
    if mse, _, _ := MeanSquaredError(whiteOpaque, blackOpaque); mse != 0.75 {
        t.Error("R, G, B, A do not contribute equally to MSE")
    }
}

func TestSSE(t *testing.T) {
    image1, _ := DecodeImage("../_testdata/images/federer1.png")
    image2, _ := DecodeImage("../_testdata/images/federer2.png")

    // Test #1: SSE between two images with different sizes should produce an error
    if _, err := sumSquaredErrors(image1, image2); err == nil {
        t.Error("SSE betwen two images with different sizes does not produce an error")
    }
}

func TestNormalizeRGBA(t *testing.T) {
    maxRGBA := uint32(65535)

    // Test #1: Normalizing largest RGBA value should return 1
    if normalizeRGBA(maxRGBA) != 1 {
        t.Error("Largest RGBA value is not normalized to 1")
    }

    // Test #2: Normalizing 0 should return 0
    if normalizeRGBA(uint32(0)) != 0 {
        t.Error("Normalizing 0 does not return 0")
    }
}
