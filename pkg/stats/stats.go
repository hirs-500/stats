package stats

import (
	"github.com/hirs-500/bank/v2/pkg/types"
	
)

// Avg рассчитывает среднюю сумму платежа 
func Avg (payments []types.Payment)  types.Money {
	 var total types.Money
	 payment := len(payments)
	 for _, v := range payments {
		 if v.Status == "FAIL"{
			 return 0
		 }
		 total+= v.Amount}
		 
		 return total/types.Money(payment)
		 
	 }
	 //TotalInCategory - находить сумму пакупок определённой котегории
	 func TotalInCategory(payments []types.Payment, category types.Category ) types.Money {
		var categorySum types.Money
		for _, v := range payments {
			if v.Category == category{
				if v.Status == "FAIL"{
					return 0 
				}
			   categorySum+=v.Amount
			}}
			
			return categorySum
		
		
	}