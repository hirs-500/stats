package stats

import (
	"github.com/hirs-500/bank/v2/pkg/types"
	"fmt"
	"reflect"
	"testing"
	
)
func ExampleAvg ()  {
	payments :=[]types.Payment{
		{ID: 2,
		Amount: 36_00,
		Category: "Cat",
	    Status:     "FAIL",},
		{ID: 2,
		Amount: 24_00,
		Category: "Cat",
		Status:    "OK", 
	      },
		{ID: 2,
		Amount: 52_00,
		Category: "Cat",
	    Status:    "OK" },
	}
	fmt.Println(Avg(payments ))
	// Output: 3800
}

func ExampleTotalCategory ()  {
	payments := []types.Payment{
		{ID:     1,
		Amount:  52_00,
		Category: "Club",
	     Status:  "FAIL",},
		{ID:     1,
		Amount:  48_00,
		Category: "Club",
	    Status:   "OK",},
		{ID:     1,
		Amount:  52_00,
		Category: "Cafe",
		Status:  "OK",},}
		
		fmt.Println(TotalInCategory(payments, "Club"))
		// Output: 4800

}

func TestCategoriesAvguser(t *testing.T) {
	payments := []types.Payment{
		{ID: 1,  Category: "auto", Amount: 1_000_00},
		 {ID: 2,  Category: "food", Amount: 2_000_00},
		 {ID: 3,  Category: "auto", Amount: 3_000_00},
		 {ID: 4,  Category: "auto", Amount: 4_000_00},
		 {ID:  5,  Category: "fun",  Amount: 5_000_00},

		
	}
	expected := map[types.Category]types.Money{
		"auto":2_666_66,
		"food":2_000_00,
		"fun" :5_000_00,}
result := CategoriesAvg(payments)
if !reflect.DeepEqual(expected,result){
	t.Errorf("invalid result, expected: %v, actual: %v", expected, result)}
}
func TestPeriodDynamic_positive (t *testing. T)  {
	first := map[types.Category]types.Money{
		"auto": 25,
		 "food": 35,
		}
	second :=map[types.Category]types.Money{
		"auto": 45,
		"food": 55,
		}

	
	expected :=map[types.Category] types.Money {
		"auto": 20,
		"food": 20,
	}
	difference :=PeriodsDynamic(first, second)
	if !reflect.DeepEqual(expected, difference){
		t.Errorf("difference > %v \n expected > %v", difference,expected)
	}
}