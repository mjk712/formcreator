package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// --------------------Таблица - Определение основной абсолютной погрешности измерений скорости движения---------
var table1Values []int
var tableSetValues []string
var tableGetValues []string
var tableErrorValues []string

// --------------------Таблица - Определение основной абсолютной погрешности измерений ускорения---------
var table2Values []int
var table2SetValues []string
var table2GetValues []string
var table2ErrorValues []string

// --------------------Таблица - Определение основной абсолютной погрешности измерений давления в тормозной магистрали (по первому каналу)-------
var table3Values []int
var table3SetValues []string
var table3GetValues []string
var table3ErrorValues []string

// --------------------Таблица - Определение основной приведенной погрешности измерений давления в тормозной магистрали (по второму каналу)-------
var table4Values []int
var table4SetValues []string
var table4GetValues []string
var table4ErrorValues []string

// --------------------Таблица - Определение основной приведенной погрешности измерений давления в тормозной магистрали (по третьему каналу)-------
var table5Values []int
var table5SetValues []string
var table5GetValues []string
var table5ErrorValues []string

var counter int
var counter2 int
var counter3 int
var counter4 int
var counter5 int

func main() {

	mapHtmlNames = make(map[string]string)

	file, err := os.Open("txtbu3pv.txt")
	if err != nil {
		fmt.Print("Err open txt")
	}
	defer file.Close()

	strDeviceName := "Поверка "

	iMpmeNumber := "Номер МПМЭ: "

	strTemperature := "Температура окружающего воздуха: "

	strRelativeHumidity := "Относительная влажность окружающего воздуха: "

	strAthmosphericPressure := "Атмосферное давление: "

	strUnitOfPress := "Контроль давлений (кПа)"

	strTableValues := "	Имитируемая: "

	str2TableValues := "	Имитируемое:  "

	str3TableValues := "	Имитируемое: "

	strPathTable := "	Имитируемый: 20000 м"

	strTimeTable := "	Имитируемое: 1800 c"

	strTravelTable := "	Имитируемое: 100 м, пройденное:"

	strSurname := "Фамилия проверяющего: "

	data := bufio.NewScanner(file)

	var txtlines []string

	// Сканируем txt файл
	for data.Scan() {
		str := data.Text()
		txtlines = append(txtlines, data.Text())

		//cвич кейс для табличных значений

		//------------------------------Эта табличка первая иначе ломает всё-------------------------------------------------------------------------
		if strings.Contains(str, strTravelTable) {
			value := strings.Replace(str, "	Имитируемое: 100 м, пройденное: ", "", -1)
			value2 := strings.Replace(value, ", погрешность: ", "|", -1)
			s := strings.Split(value2, "|")
			mapHtmlNames[travel] = s[0]
			mapHtmlNames[travelError] = s[1]
			fmt.Println(mapHtmlNames[travel], mapHtmlNames[travelError])
			str = ""
		}
		//------------------------------Тут занимаемся поиском нужных переменных и записью их в файл----------------------------------
		if strings.Contains(str, strDeviceName) {
			value := strings.Replace(str, strDeviceName, "", -1)
			mapHtmlNames[deviceName] = value
			fmt.Println(mapHtmlNames[deviceName])
		}
		//--------------переменные которые зависят от device name--------------------
		strResult := "	" + mapHtmlNames[deviceName] + " "
		iDeviceNumber := "Номер " + mapHtmlNames[deviceName] + ": "
		//Вот как тут парс лучше делать
		if strings.Contains(str, strUnitOfPress) {
			value := strings.Replace(str, "Контроль давлений ", "", -1)
			value2 := strings.Replace(value, "(", "", -1)
			value3 := strings.Replace(value2, ")", "", -1)
			mapHtmlNames[unitOfPress] = value3
			fmt.Println(mapHtmlNames[unitOfPress])
		}
		if strings.Contains(str, strResult) {
			value := strings.Replace(str, "	"+mapHtmlNames[deviceName]+" ", "", -1)
			mapHtmlNames[result] = value
			fmt.Println(mapHtmlNames[result])
		}
		if strings.Contains(str, strSurname) {
			value := strings.Replace(str, "Фамилия проверяющего: ", "", -1)
			mapHtmlNames[surname] = value
			fmt.Println(mapHtmlNames[surname])
		}
		if strings.Contains(str, iDeviceNumber) {
			value := strings.Replace(str, iMpmeNumber, "", -1)
			mapHtmlNames[deviceNumber] = value
			fmt.Println(mapHtmlNames[deviceNumber])
		}
		if strings.Contains(str, iMpmeNumber) {
			mapHtmlNames[mpmeNumber] = str[21:]
			fmt.Println(mapHtmlNames[mpmeNumber])
		}
		if strings.Contains(str, strTemperature) {
			mapHtmlNames[temperature] = str[62:]
			fmt.Println(mapHtmlNames[temperature])
		}
		if strings.Contains(str, strRelativeHumidity) {
			mapHtmlNames[relativeHumidity] = str[85:]
			fmt.Println(mapHtmlNames[relativeHumidity])
		}
		if strings.Contains(str, strAthmosphericPressure) {
			mapHtmlNames[athmosphericPressure] = str[41:]
			fmt.Println(mapHtmlNames[athmosphericPressure])
		}

		//------------------------------Тут идут последние таблицы по старому методу, чтобы лучше разбираться в этом коде посмотри этот гайд https://www.youtube.com/watch?v=dQw4w9WgXcQ----------------------

		if strings.Contains(str, strPathTable) {
			value := strings.Replace(str, "	Имитируемый: 20000 м, пройденный: ", "", -1)
			value2 := strings.Replace(value, ", погрешность: ", "|", -1)
			value3 := strings.Replace(value2, "м", "", -1)
			s := strings.Split(value3, "|")
			mapHtmlNames[way] = s[0]
			mapHtmlNames[wayError] = s[1]
			//fmt.Println(mapHtmlNames[way], mapHtmlNames[wayError])
		}
		if strings.Contains(str, strTimeTable) {
			value := strings.Replace(str, "	Имитируемое: 1800 c, измеренное: ", "", -1)
			value2 := strings.Replace(value, ", погрешность: ", "|", -1)
			s := strings.Split(value2, "|")
			mapHtmlNames[time] = s[0]
			mapHtmlNames[timeError] = s[1]
			//fmt.Println(mapHtmlNames[time], mapHtmlNames[timeError])
		}
	}
	//--------------------------------------------Начало записи в таблицы------------------
	var value2name string
	var value3name string
	var value4tableName string
	var value5tableName string

	switch mapHtmlNames[deviceName] {

	case "БУ-3ПВ":
		value3name = ", измеренное: "
		value2name = ", на индикаторе:  "
		value5tableName = "Контроль самопроизвольного ухода локомотива"
		value4tableName = "Давление по каналу 3"

	case "БУ-3ПС":
		value2name = ", измеренное: "
		value3name = ", измеренное: "
		value4tableName = "Давление по каналу 3"
		value5tableName = "Контроль частоты вращения вала ДГУ (об/мин)"

	case "БУ-4":
		value2name = ", на индикаторе:  "
		value3name = ", измеренное: "
		value4tableName = "Контроль времени установления показаний скорости"
	}
	//---------------------------------------Заполнение первой таблицы----------------------
	for i, v := range txtlines {
		if v == "Контроль ускорения (м/с2)" {
			break
		}
		//------------------та которая ломает-----------
		if strings.Contains(v, strTravelTable) {
			value := strings.Replace(v, "	Имитируемое: 100 м, пройденное: ", "", -1)
			value2 := strings.Replace(value, ", погрешность: ", "|", -1)
			s := strings.Split(value2, "|")
			mapHtmlNames[travel] = s[0]
			mapHtmlNames[travelError] = s[1]
			fmt.Println(mapHtmlNames[travel], mapHtmlNames[travelError])
			txtlines[i] = ""
		}
		//------------------А тут уже 1 таблица-----------
		if strings.Contains(v, strTableValues) {
			value := strings.Replace(v, "	Имитируемая: ", "", -1)
			value2 := strings.Replace(value, ", на индикаторе: ", "|", -1)
			value3 := strings.Replace(value2, ", погрешность: ", "|", -1)
			value4 := strings.Replace(value3, ".", ",", -1)
			s := strings.Split(value4, "|")
			tableSetValues = append(tableSetValues, s[0])
			tableGetValues = append(tableGetValues, s[1])
			tableErrorValues = append(tableErrorValues, s[2])
			txtlines[i] = ""
			counter++
			table1Values = append(table1Values, counter)
		}
	}
	//---------------------------------------Заполнение второй таблицы----------------------
	for i, v := range txtlines {
		if v == "Контроль сигнала ТЯГА" {
			break
		}
		if strings.Contains(v, str2TableValues) {
			value := strings.Replace(v, "	Имитируемое:  ", "", -1)
			value2 := strings.Replace(value, value2name, "|", -1)
			value3 := strings.Replace(value2, ", погрешность: ", "|", -1)
			value4 := strings.Replace(value3, ".", ",", -1)
			s := strings.Split(value4, "|")
			table2SetValues = append(table2SetValues, s[0])
			table2GetValues = append(table2GetValues, s[1])
			table2ErrorValues = append(table2ErrorValues, s[2])
			txtlines[i] = ""
			counter2++
			table2Values = append(table2Values, counter2)
		}
	}
	//---------------------------------------Заполнение третьей таблицы----------------------
	for i, v := range txtlines {
		if v == "Давление по каналу 2" {
			break
		}
		if strings.Contains(v, str3TableValues) {
			value := strings.Replace(v, "	Имитируемое: ", "", -1)
			value2 := strings.Replace(value, value3name, "|", -1)
			value3 := strings.Replace(value2, ", погрешность: ", "|", -1)
			value4 := strings.Replace(value3, ".", ",", -1)
			s := strings.Split(value4, "|")
			table3SetValues = append(table3SetValues, s[0])
			table3GetValues = append(table3GetValues, s[1])
			table3ErrorValues = append(table3ErrorValues, s[2])
			txtlines[i] = ""
			counter3++
			table3Values = append(table3Values, counter3)
		}
	}
	//---------------------------------------Заполнение четвёртой таблицы----------------------
	for i, v := range txtlines {
		if v == value4tableName {
			break
		}
		if strings.Contains(v, str3TableValues) {
			value := strings.Replace(v, "	Имитируемое: ", "", -1)
			value2 := strings.Replace(value, ", измеренное: ", "|", -1)
			value3 := strings.Replace(value2, ", погрешность: ", "|", -1)
			value4 := strings.Replace(value3, ".", ",", -1)
			s := strings.Split(value4, "|")
			table4SetValues = append(table4SetValues, s[0])
			table4GetValues = append(table4GetValues, s[1])
			table4ErrorValues = append(table4ErrorValues, s[2])
			txtlines[i] = ""
			counter4++
			table4Values = append(table4Values, counter4)
		}
	}
	//---------------------------------------Заполнение пятой таблицы----------------------
	for i, v := range txtlines {
		if v == value5tableName {
			break
		}
		if strings.Contains(v, str3TableValues) {
			value := strings.Replace(v, "	Имитируемое: ", "", -1)
			value2 := strings.Replace(value, ", измеренное: ", "|", -1)
			value3 := strings.Replace(value2, ", погрешность: ", "|", -1)
			value4 := strings.Replace(value3, ".", ",", -1)
			s := strings.Split(value4, "|")
			if mapHtmlNames[deviceName] == "БУ-4" || mapHtmlNames[deviceName] == "БУ-3ПВ" {
				table5SetValues = append(table5SetValues, "", "", "", "", "", "", "", "")
				table5GetValues = append(table5GetValues, "", "", "", "", "", "")
				table5ErrorValues = append(table5ErrorValues, "", "", "", "", "", "")
			} else {
				table5SetValues = append(table5SetValues, s[0])
				table5GetValues = append(table5GetValues, s[1])
				table5ErrorValues = append(table5ErrorValues, s[2])
			}
			txtlines[i] = ""
			counter5++
			table5Values = append(table5Values, counter5)
		}
	}
	if mapHtmlNames[deviceName] == "БУ-3ПВ" {
		tableSetValues = append(tableSetValues, "")
		tableGetValues = append(tableGetValues, "")
		tableErrorValues = append(tableErrorValues, "")
	}

	err = addhtml()
	fmt.Println(err)
}
