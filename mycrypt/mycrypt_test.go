package mycrypt

import (
	"reflect"
	"testing"
)



func TestKrypter(t *testing.T) {

	type test struct {
		inputMessage []rune
		chiffer      int
		want         []rune
	}
	tests := []test{
		{inputMessage: []rune("w"), chiffer: 4, want: []rune("æ")},
		{inputMessage: []rune("0"), chiffer: 4, want: []rune("4")},
		{inputMessage: []rune("Kjevik;SN39040;18.03.2022 01:50;6"), chiffer: 4, want: []rune("NnizmocXS7D484c5C 47 6466d45b94cA")},
		{inputMessage: []rune("NnizmocXS7D484c5C 47 6466d45b94cA"), chiffer: len(ALF_SEM03) - 4, want: []rune("Kjevik;SN39040;18.03.2022 01:50;6")},
	}

	for _, tc := range tests {
		got := Krypter(tc.inputMessage, ALF_SEM03, tc.chiffer)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("Feil, for chiffer %d, fikk %q, ønsket %q.", tc.chiffer, got, tc.want)
		}
	}
}


func TestSokIAlfabetet(t *testing.T) {
	type test struct {
		input rune
		want  int
	}
	tests := []test{
		{input: 'æ', want: 26},
		{input: 'a', want: 0},
	}

	for _, tc := range tests {
		got := sokIAlfabetet(tc.input, ALF_SEM03)
		if got != tc.want {
			t.Errorf("Feil, fikk %d, ønsket %d.", got, tc.want)
		}
	}
}
