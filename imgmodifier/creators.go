package imgmodifier

import (
	"fmt"

	"github.com/davidbyttow/govips/v2/vips"
)

// TODO: Optimize
func createBrightnessMatrix(img *vips.ImageRef) [][]uint8 {
	// Might be solved by this?
	// However, we still have to traverse through
	// each pixel, to convert to the ASCII (it takes a lot of time)
	// err := vips.Pixelate(img, 255.0)
	// if err != nil {
	//     panic(err)
	// }
    bytes, _, err := img.ExportPng(&vips.PngExportParams{})
    // Go through the bytes and count the average and try to convert to AScII
    fmt.Println(bytes)
	goImg, err := img.ToImage(&vips.ExportParams{})
	if err != nil {
		panic(err)
	}

	bounds := goImg.Bounds()

	var bm = make([][]uint8, bounds.Max.Y)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		bm[y] = make([]uint8, bounds.Max.X)

		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			bm[y][x] = coloursAverageNative(goImg.At(x, y))
		}
	}

	return bm
}

func createAsciiMatrix(bm [][]uint8) [][]byte {
	chars := make([][]byte, len(bm))

	for i := 0; i < len(bm); i++ {
		chars[i] = make([]byte, len(bm[i]))

		for j := 0; j < len(bm[i]); j++ {
			chars[i][j] = convertToAscii(bm[i][j])
		}
	}

	return chars
}
