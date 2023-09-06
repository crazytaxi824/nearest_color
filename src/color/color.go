package color

import (
	"encoding/hex"
	"strconv"
)

type Color struct {
	rgb [3]byte
}

// 6 digit string
func ParseHEX(h string) (*Color, error) {
	byts, err := hex.DecodeString(h)
	if err != nil {
		return nil, err
	}

	return &Color{
		rgb: [3]byte{byts[0], byts[1], byts[2]},
	}, nil
}

func (c *Color) ToHEX() string {
	return "#" + hex.EncodeToString([]byte{c.rgb[0], c.rgb[1], c.rgb[2]})
}

func (c *Color) ToRGB() [3]byte {
	return c.rgb
}

func (c *Color) FindNearest8bitColor() (color8bit string, nearestColor *Color) {
	r, ri := findNearestColorValue(c.rgb[0])
	g, gi := findNearestColorValue(c.rgb[1])
	b, bi := findNearestColorValue(c.rgb[2])

	// 通过 rgb 计算 8bit color 的位置.
	color8bit = strconv.Itoa(16 + ri*36 + gi*6 + bi)
	nearestColor = &Color{rgb: [3]byte{r, g, b}}
	return color8bit, nearestColor
}

// (hex):     0, 5F,  87,  AF,  D7,  FF
// (decimal): 0, 95, 135, 175, 215, 255
var colorCubeStep = [6]byte{0, 95, 135, 175, 215, 255}

// 返回值应该在 {0, 95, 135, 175, 215, 255} 之中.
// 不要将 byte 直接加减, 会因为超出 255 而出现错误.
func findNearestColorValue(code byte) (nearestRGB byte, index int) {
	for i := 0; i < 5; i++ {
		// 在某个区间内
		if code >= colorCubeStep[i] && code < colorCubeStep[i+1] {
			if code <= colorCubeStep[i]/2+colorCubeStep[i+1]/2 {
				return colorCubeStep[i], i
			}
			return colorCubeStep[i+1], i + 1
		}
	}

	// 如果上述情况都不存在, 则只剩下一种情况.
	return 255, 5
}
