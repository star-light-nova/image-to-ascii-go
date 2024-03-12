package imgmodifier

import (
	"github.com/davidbyttow/govips/v2/vips"
)

func createBrightnessMatrix(img *vips.ImageRef) [][]byte {
	bytes, err := img.ToBytes()
	if err != nil {
		panic(err)
	}

	var bm = make([][]byte, img.Height())
	hCnt, wCnt := 0, 0
	bm[hCnt] = make([]byte, img.Width())

	for i := 0; i <= len(bytes)-4; i += 4 {
		if wCnt == img.Width() {
			wCnt = 0

			hCnt++
			if hCnt == img.Height() {
				break
			}

			bm[hCnt] = make([]byte, img.Width())
		}

		r := bytes[i]
		g := bytes[i+1]
		b := bytes[i+2]

		avg := (r + g + b) / 3

		bm[hCnt][wCnt] = avg

		wCnt++
	}

	return bm
}

func createAsciiMatrix(bm [][]byte) [][]byte {
	chars := make([][]byte, len(bm))

	for i := 0; i < len(bm); i++ {
		chars[i] = make([]byte, len(bm[i]))

		for j := 0; j < len(bm[i]); j++ {
			chars[i][j] = convertToAscii(bm[i][j])
		}
	}

	return chars
}
