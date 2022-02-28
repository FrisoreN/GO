package event

import (
	"fmt"

	"github.com/FrisoreN/GO/state"
)

func Put(item string) {
	ky, re, ko, mann := state.GetViewState()
	BItem := state.BoatItem

	if BItem == "" {
		if item == "kylling" {
			if ky == mann && re != 1 && ko != 1 {
				ky = 1
			} else if ky == 1 {
				ky = mann
			} else if ky == 0 && mann == 2 || mann == 0 && ky == 2 {
				fmt.Println("Båten er på feil side")
			} else {
				fmt.Println("Båten er full")
				return
			}

		} else if item == "rev" {
			if re == mann && ky != 1 && ko != 1 {
				re = 1
			} else if re == 1 {
				re = mann
			} else if re == 0 && mann == 2 || mann == 0 && re == 2 {
				fmt.Println("Båten er på feil side")
			} else {
				fmt.Println("Båten er full")
				return
			}

		} else if item == "rorn" {
			if ko == mann && ky != 1 && re != 1 {
				ko = 1
			} else if ko == 1 {
				ko = mann
			} else if ko == 0 && mann == 2 || mann == 0 && ko == 2 {
				fmt.Println("Båten er på feil side")
			} else {
				fmt.Println("Båten er full")
				return
			}

		} else {
			fmt.Println("Skrive feil. Skriv Hjelp for flere commandoer")
			return
		}

	} else if BItem != "" {

		if item == "kylling" {
			ky = mann
		}
		if item == "rev" {
			re = mann
		}
		if item == "korn" {
			ko = mann
		}
	} else {
		fmt.Println("Båten har bare plass til to objekter av om gangen.")
	}

	// Oppdater viewstate
	state.ViewState(ky, re, ko, mann)
}

func Crossriver() string {
	ky, re, ko, mann := state.GetViewState()

	if mann == 0 {
		mann = 2
	} else if mann != 0 {
		mann = 0
	}

	state.ViewState(ky, re, ko, mann)

	if ky == re && re == ko && re != mann {
		fmt.Println("Nei, PASS PÅ! Kyllingen har spiste opp kornet, og reven spiste opp kyllingen. Prøv på nytt!")
		return "Du tapte"
	} else if ky == re && re != mann {
		fmt.Println("Nei, PASS PÅ! Reven spiste Kyllingen. Prøv på nytt!")
		return "Du tapte"
	} else if ky == ko && ko != mann {
		fmt.Println("Nei, PASS PÅ! Kyllingen spiste opp Kornet. Prøv på nytt!")
		return "Du tapte"
	} else {
		return "Error"
	}
}
