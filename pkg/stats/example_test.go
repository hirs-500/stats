package stats

import (
	"fmt"
	"github.com/hirs-500/bank/pkg/types"
)
func ExampleAvg ()  {
	payments :=[]types.Payment{
		{ID: 2,
		Amount: 36_00,
		Category: "Cat", },
		{ID: 2,
		Amount: 24_00,
		Category: "Cat", },
		{ID: 2,
		Amount: 52_00,
		Category: "Cat", },
	}
	fmt.Println(Avg(payments))
	// Output: 3733
}

func ExampleTotalCategory ()  {
	payments := []types.Payment{
		{ID:     1,
		Amount:  52_00,
		Category: "Club",},
		{ID:     1,
		Amount:  48_00,
		Category: "Club",},
		{ID:     1,
		Amount:  52_00,
		Category: "Cafe",},}
		
		fmt.Println(TotalInCategory(payments, "Club"))
		// Output: 10000

}