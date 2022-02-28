package event

import (
	"testing"

	"github.com/FrisoreN/GO/state"
)

func TestMoveItem(t *testing.T) {
	state.ViewState(0, 0, 0, 0)

	wanted := ""
	Put(wanted)

	bItem := state.GetBoatitem()

	if bItem != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", bItem, wanted)
	}
}

func TestCross(t *testing.T) {
	wanted := "GameOver"
	got := Crossriver()
	t.Log(got)

	if got != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", got, wanted)
	}
}
