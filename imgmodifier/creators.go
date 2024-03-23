package imgmodifier

import (
	"github.com/davidbyttow/govips/v2/vips"
)

func createCharMatrix(img *vips.ImageRef) [][]byte {
	bytes, err := img.ToBytes()
	if err != nil {
		panic(err)
	}

	var chars = make([][]byte, img.Height())
	hCnt, wCnt := 0, 0

	chars[hCnt] = make([]byte, img.Width())

	for i := 0; i <= len(bytes)-4; i += 4 {
		if wCnt == img.Width() {
			wCnt = 0

			hCnt++
			if hCnt == img.Height() {
				break
			}

			chars[hCnt] = make([]byte, img.Width())
		}

		//      red        green        blue
		avg := (bytes[i] + bytes[i+1] + bytes[i+2]) / 3

		chars[hCnt][wCnt] = convertToAscii(avg)

		wCnt++
	}

	return chars
}
