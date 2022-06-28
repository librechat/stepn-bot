package handler

import (
	"testing"
)

func TestRichPrice(t *testing.T) {
	println(richPrice("gst", "100"))
	println(richPrice("gmt", "10"))
	println(richPrice("sol"))
	println(richPrice("usd", "10"))
}
