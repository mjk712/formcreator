package main

import (
	"fmt"
	"html/template"
	"os"
)

func addhtml() {

	file, err := os.Create("hello.html")

	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()

	type htmlTest struct {
		DeviceName   string
		DeviceNumber string
		Data         string
		ReportNumber string
	}

	data := htmlTest{
		DeviceName:   mapHtmlNames[deviceName],
		DeviceNumber: mapHtmlNames[deviceNumber],
		Data:         "29.05.2023",
		ReportNumber: "0001",
	}
	tmpl, _ := template.ParseFiles("report.html")
	tmpl.Execute(file, data)

}
