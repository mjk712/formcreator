package main

import "fmt"

const (
	deviceName   = "DeviceName"
	deviceNumber = "DeviceNumber"
)

var mapHtmlNames map[string]string

func Hello() {
	mapHtmlNames = make(map[string]string)
	fmt.Println("Hello")
}
