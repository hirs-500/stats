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
			
		// FilterByCategory - возврашает платежи указанной котегории
		func FilterByCategory (payments []types.Payment, category types.Category)[]types.Payment{
			var filtred []types.Payment
			for _, payment := range payments {
				if payment.Category == category { 
					filtred = append(filtred, payment)
				}
				
			}
			return filtred
		}
	// CategoriesAvg считивает средную сумму платежа
	func CategoriesAvg(payments []types.Payment) map[types.Category] types.Money {
		categories := make(map[types.Category]types.Money)
		
		 for _, payment := range payments {
			 if _,er:=categories[payment.Category]; er{continue}
			 filtered := FilterByCategory(payments, payment.Category)
			 categories[payment.Category]=Avg(filtered)
		
			 
		 }
		 return categories
	}
				//PeriodsDynamic - расчитывает сезоный расход
		func PeriodsDynamic(first map[types.Category]types.Money, second map[types.Category]types.Money) map[types.Category]types.Money{
			categoriesSeasons:= map[types.Category]types.Money{}
	if len(first)>=len(second){
		for  v := range first {
			categoriesSeasons[v]=second[v]-first[v]
			
		}
		return categoriesSeasons

	}
			for  v := range second {
			categoriesSeasons[v]=second[v]-first[v]
			}
		return categoriesSeasons
	}
		
		
		
