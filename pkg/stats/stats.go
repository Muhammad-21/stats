package stats

import (
	"github.com/Muhammad-21/bank/v2/pkg/types"
)
func Avg(payments []types.Payment)  types.Money{
	total:=types.Money(0)
	counter:=0
	for i:=0; i<len(payments);i++{
		payment:=payments[i]
		if payment.Status!=types.StatusFail{
			total+=payment.Amount
			counter++
		}
	}
	average:=total/types.Money(counter)
	return types.Money(average)
}

func TotalInCategory(payments []types.Payment, category types.Category)  types.Money{
	total:=types.Money(0)
	for i:=0; i<len(payments);i++{
		payment:=payments[i]
		if category == payment.Category && payment.Status!=types.StatusFail{
		total+=payment.Amount
		}
	}
	return total		
}