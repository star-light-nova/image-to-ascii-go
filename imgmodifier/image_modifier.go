package imgmodifier

import (
	"github.com/davidbyttow/govips/v2/vips"
	"os"
)

func ImgToAscii(imgBytes []byte, resizeScale float64) [][]byte {
	vips.Startup(&vips.Config{})
	// defer vips.Shutdown()

	img, err := vips.NewImageFromBuffer(imgBytes)
	if err != nil {
		panic(err)
	}

	err = img.Resize(resizeScale, vips.KernelLanczos3)
	if err != nil {
		panic(err)
	}

	chars := createCharMatrix(img)

	f, err := os.Create("asciioutput.txt")
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(chars); i++ {
		f.Write(chars[i])
		f.Write([]byte{10})
	}

	return chars
}
