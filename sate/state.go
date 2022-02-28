package state

import "fmt"

var BoatItem string
var ky int
var re int
var ko int
var mann int

func ViewState(Item1 int, Item2 int, Item3 int, Item4 int) {

	ky = Item1
	re = Item2
	ko = Item3
	mann = Item4

	Vestky := ""
	Vestre := ""
	Vestko := ""

	Østky := ""
	Østre := ""
	Østko := ""

	BoatItem := ""

	//Kylling
	if ky == 1 {
		BoatItem = "Kylling"
	} else if ky == 0 {
		Vestky = "Kylling"
	} else {
		Østky = "Kylling"
	}

	//Rev
	if re == 1 {
		BoatItem = "Rev"
	} else if re == 0 {
		Vestre = "Rev"
	} else {
		Østre = "Rev"
	}

	//Korn
	if ko == 1 {
		BoatItem = "Korn"
	} else if ko == 0 {
		Vestko = "Korn"
	} else {
		Østko = "Korn"
	}

	//Båt + Menneske
	if mann == 0 {
		fmt.Printf("VEST [%v %v %v ]---\\ \\_Mann_%v_/ _________________/---[ %v %v %v] ØST \n", Vestky, Vestre, Vestko, BoatItem, Østky, Østre, Østko)
	} else {
		fmt.Printf("VEST [%v %v %v ]---\\_________________ \\_Mann_%v_/ /---[ %v %v %v] ØST\n", Vestky, Vestre, Vestko, BoatItem, Østky, Østre, Østko)
	}

}

func GetViewState() (int, int, int, int) {
	return ky, re, ko, mann
}

func GetBoatitem() string {
	return BoatItem
}

func GetWin() string {
	if ky == 2 && re == 2 && ko == 2 && mann != 0 {
		return "Win"
	} else {
		return "Nono"
	}

}
