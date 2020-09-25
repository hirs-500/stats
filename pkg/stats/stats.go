package stats

import (
	"github.com/hirs-500/bank/v2/pkg/types"
	
)

// Avg рассчитывает среднюю сумму платежа 
func Avg (payments []types.Payment )  types.Money {
	 var total types.Money
	 i := 0
	 for _, v := range payments {
		 if v.Status == types.StatusFail{
			 continue
		 }
		 total+= v.Amount
		i++}
		 
		 return total/types.Money(i)
		 
	 }
	 //TotalInCategory - находить сумму пакупок определённой котегории
	 func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
		var categorySum types.Money
		for _, v := range payments {
			if  v.Status == types.StatusFail {
			 continue }
			if v.Category == category{
				categorySum+=v.Amount
				}
			}
			return categorySum
			}	
		
			
		
		
		
	