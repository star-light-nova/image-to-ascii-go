package imgmodifier

import (
	"github.com/davidbyttow/govips/v2/vips"
	"os"
)

func ImgToAscii(imgBytes []byte, resizeScale float64) [][]byte {
	vips.Startup(&vips.Config{})
	// Find out how to gracefully shutdown after finishing the server.
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
	defer f.Close()
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(chars); i++ {
		f.Write(chars[i])
		f.Write([]byte{10}) // \n
	}

	return chars
}
