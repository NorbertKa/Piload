package main

import (
	"bytes"
	"image"
	"image/jpeg"
)

func fromPngToJpg(inFile *bytes.Buffer) {
	t, _, _ := image.Decode(inFile)
	jpeg.Encode(inFile, t, nil)
}
