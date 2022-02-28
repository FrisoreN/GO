package event

import (
	"fmt"

	"github.com/FrisoreN/GO/state"
)

func Put(item string) {
	ky, re, ko, mann := state.GetViewState()
	BItem := state.BoatItem

	if BItem == "" {
		if item == "Kylling" {
			if ky == mann && re != 1 && ko != 1 {
				ky = 1
			} else if ky == 1 {
				ky = mann
			} else if ky == 0 && mann == 2 || mann == 0 && ky == 2 {
				fmt.Println("Båten er på feil side")
			} else {
				fmt.Println("Båten har ikke plass")
				return
			}

		} else if item == "Rev" {
			if re == mann && ky != 1 && ko != 1 {
				re = 1
			} else if re == 1 {
				re = mann
			} else if re == 0 && mann == 2 || mann == 0 && re == 2 {
				fmt.Println("Båten er på feil side")
			} else {
				fmt.Println("Båten har ikke plass")
				return
			}

		} else if item == "Korn" {
			if ko == mann && ky != 1 && re != 1 {
				ko = 1
			} else if ko == 1 {
				ko = mann
			} else if ko == 0 && mann == 2 || mann == 0 && ko == 2 {
				fmt.Println("Båten er på feil side")
			} else {
				fmt.Println("Båten har ikke plass")
				return
			}

		} else {
			fmt.Println("Skrive feil, prøv igjen. Skriv Hjelp for flere commandoer")
			return
		}

	} else if BItem != "" {

		if item == "Kylling" {
			ky = mann
		}
		if item == "Rev" {
			re = mann
		}
		if item == "Korn" {
			ko = mann
		}
	} else {
		fmt.Println("Båten har bare plass til to om gangen.")
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
		fmt.Println("NEI NEI! PASS PÅ!!! Kyllingen spiste opp Kornet, og Reven spiste opp Kyllingen. Prøv Igjen!")
		return "GameOver"
	} else if ky == re && re != mann {
		fmt.Println("NEI NEI! PASS PÅ!!! Reven spiste Kyllingen. Prøv Igjen!")
		return "GameOver"
	} else if ky == ko && ko != mann {
		fmt.Println("NEI NEI! PASS PÅ!!! Kyllingen spiste opp Kornet. Prøv Igjen!")
		return "GameOver"
	} else {
		return "Error"
	}
}
