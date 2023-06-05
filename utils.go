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
	for _, element := range s {

		el := strings.Replace(element, ",", ".", 1)
		if el == "" {
			el = "0.0"
		}
		a, err := strconv.ParseFloat(el, 64)
		if err != nil {
			return err.Error()
		}
		if a > max {
			max = a
		}
	}
	redMax := fmt.Sprintf("%.2f", max)
	out := strings.Replace(redMax, ".", ",", 1)
	return out
}

func addhtml() error {

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
		SpeedTabl:    fillSpeedTable(),
		SpeedUpTabl:  fillSpeedUpTable(),
		PressTabl:    fillPressTable(),
		Press2Tabl:   fillPress2Table(),
		Press3Tabl:   fillPress3Table(),
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
	return err

}
