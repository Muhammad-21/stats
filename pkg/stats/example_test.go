package stats

import (
	"github.com/Muhammad-21/bank/v2/pkg/types"
	"fmt"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			Amount: 10_000,
			Status: "OK",
		},
		{
			Amount: 20_000,
			Status: types.StatusFail,
		},
		{
			Amount: 30_000,
			Status: "OK",
		},
	}
	average:=Avg(payments)
	fmt.Println(average)
//Output:
//20000
}

func ExampleTotalInCategory(){
	payments := []types.Payment{
		{
			Category: "aaa",
			Amount: 10_000,
			Status: "OK",
		},
		{
			Category: "aaa",
			Amount: 20_000,
			Status: types.StatusFail,
		},
		{
			Category: "bbb",
			Amount: 20_000,
			Status: types.StatusFail,
		},
	}
	fmt.Println(TotalInCategory(payments,"aaa"))
//Output:
//10000
}
