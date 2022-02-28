package state

import "testing"

func TestViewState(t *testing.T) {
	//ViewState() har returnerer ingen verdi. Printer i stedet
	ViewState(0, 0, 0, 0)
}

func TestGetViewState(t *testing.T) {
	// wanted nummer
	n1, n2, n3, n4 := 0, 0, 0, 0

	// got variabel
	v1, v2, v3, v4 := GetViewState()

	if v1 != n1 {
		t.Errorf("Feil, mottok %v i v1, ønsket %v", v1, n1)
	}
	if v2 != n2 {
		t.Errorf("Feil, mottok %v i v2, ønsket %v", v2, n2)
	}
	if v3 != n3 {
		t.Errorf("Feil, mottok %v i v3, ønsket %v", v3, n3)
	}
	if v4 != n4 {
		t.Errorf("Feil, mottok %v i v4, ønsket %v", v4, n4)
	}
}
