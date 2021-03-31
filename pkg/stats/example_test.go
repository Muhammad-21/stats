package stats

import (
	"github.com/Muhammad-21/bank/pkg/types"
	"fmt"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			Amount: 10_000,
		},
		{
			Amount: 20_000,
		},
	}
	average:=Avg(payments)
	fmt.Println(average)
//Output:
//15000
}

func ExampleTotalInCategory(){
	payments := []types.Payment{
		{
			Category: "aaa",
			Amount: 10_000,
		},
		{
			Category: "aaa",
			Amount: 20_000,
		},
		{
			Category: "bbb",
			Amount: 20_000,
		},
	}
	fmt.Println(TotalInCategory(payments,"aaa"))
//Output:
//30000
}
