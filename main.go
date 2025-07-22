package main

import (
	"fmt"
	"image/jpeg"
	"os"
	"path/filepath"

	"golang.org/x/image/webp"
)

func main() {
	for index, fileName := range os.Args {
		if index != 0 {
			fileInfo, err := os.Stat(fileName)
			if err == nil {
				if filepath.Ext(fileInfo.Name()) == ".webp" {
					fileObject, err := os.Open(fileName)
					if err != nil {
						fmt.Println("File isn't valid.")
						os.Exit(1)
					}
					defer fileObject.Close()

					webpFile, err := webp.Decode(fileObject)
					if err != nil {
						fmt.Println("WebP isn't valid.")
						os.Exit(1)
					}

					jpegName := fileName[0:len(fileName)-len(filepath.Ext(fileName))] + ".jpeg"

					jpegFile, err := os.Create(jpegName)
					if err != nil {
						fmt.Println("Error:", err)
						os.Exit(1)
					}
					defer jpegFile.Close()

					err = jpeg.Encode(jpegFile, webpFile, &jpeg.Options{})
					if err != nil {
						fmt.Println("Encoding doesn't work.")
						os.Exit(1)
					}
					fmt.Println("Conversion is completed")
				} else {
					fmt.Println("File isn't WebP.")
					os.Exit(1)
				}
			} else if os.IsNotExist(err) {
				fmt.Println("File doesn't exist.")
				os.Exit(1)
			} else {
				fmt.Println("Unknow file error.")
				os.Exit(1)
			}
		}
	}
}
