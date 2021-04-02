package stats

import (
	"reflect"
	"testing"
	"github.com/Muhammad-21/bank/v2/pkg/types")



func TestCategoriesAvg_nil(t *testing.T) {
	var payments []types.Payment
	result:= CategoriesAvg(payments)
	if len(result)!=0{
		t.Error("result len != 0")
	}
}
	
func TestCategoriesAvg_empty(t *testing.T) {
	payments := []types.Payment{}
	result:= CategoriesAvg(payments)
	if len(result)!=0{
		t.Error("result len != 0")
	}
}

func TestCategoriesAvg(t *testing.T) {
	payments := []types.Payment{
		{
			ID: 1,
			Category: "auto",
			Amount: 10000,
		},
		{
			ID: 2,
			Category: "auto",
			Amount: 20000,
		},
		{
			ID: 3,
			Category: "food",
			Amount: 10000,
		},
		{
			ID: 4,
			Category: "food",
			Amount: 10000,
		},
		{
			ID: 5,
			Category: "auto",
			Amount: 10000,
		},
	}
	expected := map[types.Category]types.Money{
		"auto":13333,
		"food":10000,
	}
	result:=CategoriesAvg(payments)
	if !reflect.DeepEqual(expected,result){
		t.Errorf("invalid result, expected: %v, actual: %v",expected,result)
	}
}

func TestPeriodsDynamic(t *testing.T) {
	first:=map[types.Category]types.Money{
		"auto":10,
		"food":20,
	}
	second:=map[types.Category]types.Money{
		"auto":5,
		"food":3,
		"mobile":5,
	}
	expected := map[types.Category]types.Money{
		"auto":-5,
		"food":-17,
		"mobile":5,
	}
	result:=PeriodsDynamic(first,second)
	if !reflect.DeepEqual(expected,result){
		t.Errorf("invalid result, expected: %v, actual: %v",expected,result)
	}
}