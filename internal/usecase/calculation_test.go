package usecase

import (
	"reflect"
	"testing"
)

func TestCalculatePacksBasic(t *testing.T) {
	packs := []int{250, 500, 1000, 2000, 5000}

	result, err := CalculatePacks(250, packs)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if result.UsedSum != 250 {
		t.Errorf("Expected UsedSum=250 got=%d", result.UsedSum)
	}
	expectedPacks := map[int]int{250: 1}
	if !reflect.DeepEqual(result.Packs, expectedPacks) {
		t.Errorf("Expected packs %v got %v", expectedPacks, result.Packs)
	}

	result, err = CalculatePacks(251, packs)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	// Expected best solution: 1x500 = 500 total, overhead=249
	if result.UsedSum != 500 {
		t.Errorf("Expected UsedSum=500, got=%d", result.UsedSum)
	}
	expectedPacks = map[int]int{500: 1}
	if !reflect.DeepEqual(result.Packs, expectedPacks) {
		t.Errorf("Expected packs %v got %v", expectedPacks, result.Packs)
	}
}
