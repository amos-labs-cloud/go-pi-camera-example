package main

import (
	"context"
	"github.com/vladimirvivien/go4vl/device"
	"log"
	"os"
)

func main() {
	log.Println("opening camera device")
	dev, err := device.Open("/dev/video0", device.WithBufferSize(1))
	if err != nil {
		log.Fatal(err)
	}
	defer dev.Close() // close the camera after this function ends

	log.Println("Starting the camera device")
	if err := dev.Start(context.TODO()); err != nil {
		log.Fatal(err)
	}

	// capture frame
	log.Println("Capturing a frame")
	frame := <-dev.GetOutput()

	log.Println("Creating file pic.jpg")
	file, err := os.Create("pic.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log.Println("Writing capture frames to pic.jpg")
	if _, err := file.Write(frame); err != nil {
		log.Fatal(err)
	}
	log.Println("Write complete")
}
