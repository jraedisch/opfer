package main

import (
	"reflect"
	"testing"
)

func TestDistributeVictims(t *testing.T) {
	names := []string{"Moriarty", "Skeletor", "Cruella de Vil"}
	expected := map[string]string{"Moriarty": "Skeletor", "Skeletor": "Cruella de Vil", "Cruella de Vil": "Moriarty"}
	actual := distributeVictims(names)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected\n%v\ngot\n%v", expected, actual)
	}
}
