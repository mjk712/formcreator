package main

import (
	"fmt"
	"html/template"
	"os"
	"strconv"
	"strings"
)

var htmlPath string

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
	SpeedError           string
	SpeedUpError         string
	Channel1AbsError     string
	Channel2AbsError     string
	Channel3AbsError     string
	Result               string
	Surname              string
	SpeedTabl            interface{}
	SpeedUpTabl          interface{}
	PressTabl            interface{}
	Press2Tabl           interface{}
	Press3Tabl           interface{}
	//----------Определение основной абсолютной погрешности измерений пройденного пути-------
	Way      string
	WayError string
	//----------Определение основной абсолютной погрешности отсчета текущего времени-------
	Time      string
	TimeError string
	//----------Определение основной абсолютной погрешности измерений перемещения транспортного средства от заданной отметки----
	Travel      string
	TravelError string
	//------------test test test---------------
	TestTest string
}

func maxErrElem(s []string) string {
	var max float64
	max = 0
	for i, element := range s {
		if i > 8 {
			break
		} // Было бы круто менять тут значение динамически под размер таблицы
		el := strings.Replace(element, ",", ".", 1)
		a, err := strconv.ParseFloat(el, 64)
		if err != nil {
			return err.Error()
		}
		if a > max {
			max = a
		}
	}
	redMax := fmt.Sprintf("%.1f", max)
	out := strings.Replace(redMax, ".", ",", 1)
	return out
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
		Result:               mapHtmlNames[result],
		Surname:              mapHtmlNames[surname],
		SpeedError:           maxErrElem(tableErrorValues),
		SpeedUpError:         maxErrElem(table2ErrorValues),
		Channel1AbsError:     maxErrElem(table3ErrorValues),
		Channel2AbsError:     maxErrElem(table4ErrorValues),
		Channel3AbsError:     maxErrElem(table5ErrorValues),
		TestTest:             "22",

		Data:         "29.05.2023",
		ReportNumber: "0001",
		SpeedTabl: []SpeedTable{
			{SetValue: tableSetValues[0], GetValue: tableGetValues[0], ErrorValue: tableErrorValues[0]},
			{SetValue: tableSetValues[1], GetValue: tableGetValues[1], ErrorValue: tableErrorValues[1]},
			{SetValue: tableSetValues[2] + ",0", GetValue: tableGetValues[2] + ",0", ErrorValue: tableErrorValues[2]},
			{SetValue: tableSetValues[3] + ",0", GetValue: tableGetValues[3] + ",0", ErrorValue: tableErrorValues[3]},
			{SetValue: tableSetValues[4] + ",0", GetValue: tableGetValues[4] + ",0", ErrorValue: tableErrorValues[4]},
			{SetValue: tableSetValues[5] + ",0", GetValue: tableGetValues[5] + ",0", ErrorValue: tableErrorValues[5]},
			{SetValue: tableSetValues[6] + ",0", GetValue: tableGetValues[6] + ",0", ErrorValue: tableErrorValues[6]},
			{SetValue: tableSetValues[7] + ",0", GetValue: tableGetValues[7] + ",0", ErrorValue: tableErrorValues[7]},
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
		Press3Tabl: []Press3Table{
			{SetValue: table5SetValues[0], GetValue: table5GetValues[0], ErrorValue: table5ErrorValues[0]},
			{SetValue: table5SetValues[1], GetValue: table5GetValues[1], ErrorValue: table5ErrorValues[1]},
			{SetValue: table5SetValues[2], GetValue: table5GetValues[2], ErrorValue: table5ErrorValues[2]},
			{SetValue: table5SetValues[3], GetValue: table5GetValues[3], ErrorValue: table5ErrorValues[3]},
			{SetValue: table5SetValues[4], GetValue: table5GetValues[4], ErrorValue: table5ErrorValues[4]},
			{SetValue: table5SetValues[5], GetValue: table5GetValues[5], ErrorValue: table5ErrorValues[5]},
			{SetValue: table5SetValues[6], GetValue: table5GetValues[6], ErrorValue: table5ErrorValues[6]},
			{SetValue: table5SetValues[7], GetValue: table5GetValues[7], ErrorValue: table5ErrorValues[7]},
		},
	}

	switch mapHtmlNames[deviceName] {
	case "БУ-3ПС":
		htmlPath = "report_bu3ps.html"

	case "БУ-3ПА":
		htmlPath = "report_bu3pa3pv.html"

	case "БУ-3ПВ":
		htmlPath = "report_bu3pa3pv.html"

	case "БУ-4":
		htmlPath = "report_bu4.html"
	}

	tmpl, _ := template.ParseFiles(htmlPath)

	err = tmpl.Execute(file, data)
	fmt.Println(err)

}
