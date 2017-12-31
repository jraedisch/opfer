package main

import (
	"reflect"
	"sort"
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

func TestShuffleNames(t *testing.T) {
	sorted := []string{"A", "B", "C", "D"}
	shuffled := shuffleNames(sorted)
	sort.Strings(shuffled)
	if !reflect.DeepEqual(sorted, shuffled) {
		t.Errorf("shuffled does not contain the right names: %v", shuffled)
	}
}
