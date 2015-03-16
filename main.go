package main

import (
	"bytes"
	"fmt"
	"image/color"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"

	"github.com/daddye/vips"
	"github.com/disintegration/imaging"
)

type ImageSet interface {
	Set(x, y int, c color.Color)
}

func main() {
	var dirname string = "/home/kmulvey/Downloads/banff_files"
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal("cant read dir")
	}
	for _, file := range files {
		options := vips.Options{
			Width:        8,
			Height:       8,
			Crop:         false,
			Extend:       vips.EXTEND_WHITE,
			Interpolator: vips.BILINEAR,
			Gravity:      vips.CENTRE,
			Quality:      95,
		}
		f, _ := os.Open(dirname + "/" + file.Name())
		inBuf, _ := ioutil.ReadAll(f)
		buf, err := vips.Resize(inBuf, options)
		if err != nil {
			log.Fatal(err)
		}
		outImg, err := jpeg.Decode(bytes.NewReader(buf))
		if err != nil {
			log.Fatal(err)
		}
		gray := imaging.Grayscale(outImg)
		var total uint8 = 0
		for _, value := range gray.Pix {
			total += value
		}
		fmt.Printf("%v\n", int(total))
		fmt.Printf("%v\n", len(gray.Pix))
		fmt.Println(float32(total) / float32(len(gray.Pix)))
		break
		/*
			out, err := os.Create("./out/" + file.Name())
			if err != nil {
				fmt.Println("cant create fike")
				log.Fatal(err)
			}
			err = jpeg.Encode(out, gray, nil)
			if err != nil {
				fmt.Println("cant create fike")
				log.Fatal(err)
			}
		*/
	}
}
