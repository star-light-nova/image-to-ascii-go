package imgmodifier

import (
	"github.com/davidbyttow/govips/v2/vips"
	"os"
	// "fmt"
	// "encoding/json"
)

func ImgToAscii(imgBytes []byte, resizeScale float64) [][]byte {
	vips.Startup(
		&vips.Config{
			ConcurrencyLevel: 2,
		})
	// defer vips.Shutdown()

	img, err := vips.NewImageFromBuffer(imgBytes)
	if err != nil {
		panic(err)
	}

	err = img.Resize(resizeScale, vips.KernelLanczos2)
	if err != nil {
		panic(err)
	}

	brightnessMatrix := createBrightnessMatrix(img)
	chars := createAsciiMatrix(brightnessMatrix)

	f, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(chars); i++ {
		f.Write(chars[i])
		f.Write([]byte{10})
	}

	return chars
}
