package imgmodifier

// Sorted by how each rune fills it's own box.
const ASCIICHARS = " `.-'\":_,^=;><+!rc*/z?sLTv)J7(|Fi{C}fI31tlu[neoZ5Yxjya]2ESwqkP6h9d4VpOGbUAKXHm8RD#$Bg0MNWQ%&@"

func convertToAscii(brightness byte) byte {
	asciiIndex := (brightness % (byte(len(ASCIICHARS) - 1)))

	return ASCIICHARS[asciiIndex]
}
