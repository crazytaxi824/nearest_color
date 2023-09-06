package color

import "testing"

func TestParseHex(t *testing.T) {
	c, err := ParseHEX("FF0000")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(c.ToHEX())
	t.Log(c.ToRGB())
}

func TestNearestVal(t *testing.T) {
	t.Log(findNearestColorValue(200))
}

func TestFindNear8BitColor(t *testing.T) {
	c, err := ParseHEX("CA7482")
	if err != nil {
		t.Error(err)
		return
	}

	i, nc := c.FindNearest8bitColor()
	t.Log(i)
	t.Log(nc.ToHEX())
}
