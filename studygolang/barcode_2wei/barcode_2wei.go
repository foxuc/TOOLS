package main

import (
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"image"
	"image/png"
	"log"
	"os"
	"time"
)

func writePng(filename string, img image.Image) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	err = png.Encode(file, img)
	// err = jpeg.Encode(file, img, &jpeg.Options{100})      //图像质量值为100，是最好的图像显示
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	log.Println(file.Name())
}

func main() {
	args := os.Args
	base64 :=""
	NowTime :=time.Now().Format("2006_01_02_15_04_05")
	log.Println("===  20181018 ver:0.1 by hi==")
	if args == nil || len(args) !=2  {
		log.Println("URL Error !")
		return
	}else {
		base64=args[1]
	}

	//base64 := "http://www.baidu.com"

	log.Println("Original data:", base64)
	code, err := qr.Encode(base64, qr.L, qr.Unicode)
	// code, err := code39.Encode(base64)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Encoded data: ", code.Content())

	if base64 != code.Content() {
		log.Fatal("data differs")
	}

	code, err = barcode.Scale(code, 300, 300)
	if err != nil {
		log.Fatal(err)
	}
		writePng("Out_barcode_"+NowTime+".png", code)
}
