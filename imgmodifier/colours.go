package imgmodifier

import "image/color"

// For Vips
func coloursAverage(rgba []float64) uint8 {
	r, g, b := uint8(rgba[0]), uint8(rgba[1]), uint8(rgba[2])

	r8 := uint8((r >> 16) & 0xff)
	g8 := uint8((g >> 8) & 0xff)
	b8 := uint8(b & 0xff)

	return (r8 + g8 + b8) / 3
}

// Native go implementation
func coloursAverageNative(color color.Color) uint8 {
    r, g, b, _ := color.RGBA()

    r8 := (uint8(r) >> 16) & 0xff
    g8 := (uint8(g) >> 8) & 0xff
    b8 := (uint8(b) & 0xff)

	return (r8 + g8 + b8) / 3
}
