package main

import (
	"fmt"
	"html/template"
	"os"
)

type htmlTest struct {
	DeviceName           string
	DeviceNumber         string
	Data                 string
	ReportNumber         string
	MpmeNumber           string
	Temperature          string
	RelativeHumidity     string
	AthmosphericPressure string
	UnitOfPress          string
	SpeedTabl            interface{}
	SpeedUpTabl          interface{}
	PressTabl            interface{}
	Press2Tabl           interface{}
	//----------Определение основной абсолютной погрешности измерений пройденного пути-------
	Way      string
	WayError string
	//----------Определение основной абсолютной погрешности отсчета текущего времени-------
	Time      string
	TimeError string
	//----------Определение основной абсолютной погрешности измерений перемещения транспортного средства от заданной отметки----
	Travel      string
	TravelError string
}

func addhtml() {

	file, err := os.Create("hello.html")

	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()

	data := htmlTest{
		DeviceName:           mapHtmlNames[deviceName],
		DeviceNumber:         mapHtmlNames[deviceNumber],
		MpmeNumber:           mapHtmlNames[mpmeNumber],
		Temperature:          mapHtmlNames[temperature],
		RelativeHumidity:     mapHtmlNames[relativeHumidity],
		AthmosphericPressure: mapHtmlNames[athmosphericPressure],
		UnitOfPress:          mapHtmlNames[unitOfPress],
		Way:                  mapHtmlNames[way],
		WayError:             mapHtmlNames[wayError],
		Time:                 mapHtmlNames[time],
		TimeError:            mapHtmlNames[timeError],
		Travel:               mapHtmlNames[travel],
		TravelError:          mapHtmlNames[travelError],

		Data:         "29.05.2023",
		ReportNumber: "0001",
		SpeedTabl: []SpeedTable{
			{SetValue: tableSetValues[0], GetValue: tableGetValues[0], ErrorValue: tableErrorValues[0]},
			{SetValue: tableSetValues[1], GetValue: tableGetValues[1], ErrorValue: tableErrorValues[1]},
			{SetValue: tableSetValues[2], GetValue: tableGetValues[2], ErrorValue: tableErrorValues[2]},
			{SetValue: tableSetValues[3], GetValue: tableGetValues[3], ErrorValue: tableErrorValues[3]},
			{SetValue: tableSetValues[4], GetValue: tableGetValues[4], ErrorValue: tableErrorValues[4]},
			{SetValue: tableSetValues[5], GetValue: tableGetValues[5], ErrorValue: tableErrorValues[5]},
			{SetValue: tableSetValues[6], GetValue: tableGetValues[6], ErrorValue: tableErrorValues[6]},
			{SetValue: tableSetValues[7], GetValue: tableGetValues[7], ErrorValue: tableErrorValues[7]},
		},
		SpeedUpTabl: []SpeedUpTable{
			{SetValue: table2SetValues[0], GetValue: table2GetValues[0], ErrorValue: table2ErrorValues[0]},
			{SetValue: table2SetValues[1], GetValue: table2GetValues[1], ErrorValue: table2ErrorValues[1]},
			{SetValue: table2SetValues[2], GetValue: table2GetValues[2], ErrorValue: table2ErrorValues[2]},
			{SetValue: table2SetValues[3], GetValue: table2GetValues[3], ErrorValue: table2ErrorValues[3]},
			{SetValue: table2SetValues[4], GetValue: table2GetValues[4], ErrorValue: table2ErrorValues[4]},
			{SetValue: table2SetValues[5], GetValue: table2GetValues[5], ErrorValue: table2ErrorValues[5]},
			{SetValue: table2SetValues[6], GetValue: table2GetValues[6], ErrorValue: table2ErrorValues[6]},
			{SetValue: table2SetValues[7], GetValue: table2GetValues[7], ErrorValue: table2ErrorValues[7]},
		},
		PressTabl: []PressTable{
			{SetValue: table3SetValues[0], GetValue: table3GetValues[0], ErrorValue: table3ErrorValues[0]},
			{SetValue: table3SetValues[1], GetValue: table3GetValues[1], ErrorValue: table3ErrorValues[1]},
			{SetValue: table3SetValues[2], GetValue: table3GetValues[2], ErrorValue: table3ErrorValues[2]},
			{SetValue: table3SetValues[3], GetValue: table3GetValues[3], ErrorValue: table3ErrorValues[3]},
			{SetValue: table3SetValues[4], GetValue: table3GetValues[4], ErrorValue: table3ErrorValues[4]},
			{SetValue: table3SetValues[5], GetValue: table3GetValues[5], ErrorValue: table3ErrorValues[5]},
			{SetValue: table3SetValues[6], GetValue: table3GetValues[6], ErrorValue: table3ErrorValues[6]},
			{SetValue: table3SetValues[7], GetValue: table3GetValues[7], ErrorValue: table3ErrorValues[7]},
		},
		Press2Tabl: []Press2Table{
			{SetValue: table4SetValues[0], GetValue: table4GetValues[0], ErrorValue: table4ErrorValues[0]},
			{SetValue: table4SetValues[1], GetValue: table4GetValues[1], ErrorValue: table4ErrorValues[1]},
			{SetValue: table4SetValues[2], GetValue: table4GetValues[2], ErrorValue: table4ErrorValues[2]},
			{SetValue: table4SetValues[3], GetValue: table4GetValues[3], ErrorValue: table4ErrorValues[3]},
			{SetValue: table4SetValues[4], GetValue: table4GetValues[4], ErrorValue: table4ErrorValues[4]},
			{SetValue: table4SetValues[5], GetValue: table4GetValues[5], ErrorValue: table4ErrorValues[5]},
			{SetValue: table4SetValues[6], GetValue: table4GetValues[6], ErrorValue: table4ErrorValues[6]},
			{SetValue: table4SetValues[7], GetValue: table4GetValues[7], ErrorValue: table4ErrorValues[7]},
		},
	}

	tmpl, _ := template.ParseFiles("report.html")

	err = tmpl.Execute(file, data)
	fmt.Println(err)

}
