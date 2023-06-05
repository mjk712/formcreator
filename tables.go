package main

type SpeedTable struct {
	SetValue   string
	GetValue   string
	ErrorValue string
}
type SpeedUpTable struct {
	SetValue   string
	GetValue   string
	ErrorValue string
}

type PressTable struct {
	SetValue   string
	GetValue   string
	ErrorValue string
}

type Press2Table struct {
	SetValue   string
	GetValue   string
	ErrorValue string
}

type Press3Table struct {
	SetValue   string
	GetValue   string
	ErrorValue string
}

func fillSpeedTable() []SpeedTable {
	var a []SpeedTable
	for i, _ := range table1Values {
		s := SpeedTable{SetValue: tableSetValues[i], GetValue: tableGetValues[i], ErrorValue: tableErrorValues[i]}
		if i >= 2 {
			s := SpeedTable{SetValue: tableSetValues[i] + ",0", GetValue: tableGetValues[i] + ",0", ErrorValue: tableErrorValues[i]}
			a = append(a, s)
			continue
		}
		a = append(a, s)
	}
	return a
}

func fillSpeedUpTable() []SpeedUpTable {
	var a []SpeedUpTable
	for i, _ := range table2Values {
		s := SpeedUpTable{SetValue: table2SetValues[i], GetValue: table2GetValues[i], ErrorValue: table2ErrorValues[i]}
		a = append(a, s)
	}
	return a
}

func fillPressTable() []PressTable {
	var a []PressTable
	for i, _ := range table3Values {
		s := PressTable{SetValue: table3SetValues[i], GetValue: table3GetValues[i], ErrorValue: table3ErrorValues[i]}
		a = append(a, s)
	}
	return a
}

func fillPress2Table() []Press2Table {
	var a []Press2Table
	for i, _ := range table4Values {
		s := Press2Table{SetValue: table4SetValues[i], GetValue: table4GetValues[i], ErrorValue: table4ErrorValues[i]}
		a = append(a, s)
	}
	return a
}

func fillPress3Table() []Press3Table {
	var a []Press3Table
	for i, _ := range table5Values {
		s := Press3Table{SetValue: table5SetValues[i], GetValue: table5GetValues[i], ErrorValue: table5ErrorValues[i]}
		a = append(a, s)
	}
	return a
}
