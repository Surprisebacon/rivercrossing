package state

import "testing"

//test function implemented in line whit the Golang's testing tool
func TestViewState(t testing.T) {
wanted := "[kylling rev korn hs ---\\ \\__/____________/---]"
state := ViewState(); 
if state !=wanted {
t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
}
}
