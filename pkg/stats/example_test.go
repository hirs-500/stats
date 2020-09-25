package stats

import (
	"github.com/hirs-500/bank/v2/pkg/types"
	"fmt"
	
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
	// Output: 2533
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