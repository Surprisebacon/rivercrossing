package event

import "testing"

func TestPut(t *testing.T) {//hva forventer jeg? 
wanted :="[kylling rev korn ---\\ \\_korn_//____________/---]"
got := Put("korn") // Hva fikk jeg?
if got != wanted {t.Errorf("feil, fikk %q, ønsket %q.", got, wanted) 
} 
} 
