package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	mapHtmlNames = make(map[string]string)

	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Print("Err open txt")
	}
	defer file.Close()

	strDeviceName := "Поверка "

	iDeviceNumber := "Номер БУ-3ПС: "

	data := bufio.NewScanner(file)

	for data.Scan() {
		str := data.Text()
		if strings.Contains(str, strDeviceName) {
			//fmt.Println(str[15:])
			mapHtmlNames[deviceName] = str[15:]
			fmt.Println(mapHtmlNames[deviceName])
		}
		if strings.Contains(str, iDeviceNumber) {

			mapHtmlNames[deviceNumber] = str[23:]
			fmt.Println(mapHtmlNames[deviceNumber])
		}
	}
	addhtml()
}
