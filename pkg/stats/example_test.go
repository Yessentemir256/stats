package stats

import (
	"fmt"
	"github.com/Yessentemir256/bank/pkg/bank/types"
)

func ExampleTotal() {
	fmt.Println(Total([]types.Payment{
		{
			Amount: 1_000_00,
		},
	}))
	fmt.Println(Total([]types.Payment{
		{
			Amount: 1_000_00,
		},
		{
			Amount: 2_000_00,
		},
	}))
	fmt.Println(Total([]types.Payment{
		{
			Amount: 1_000_00,
		},
		{
			Amount: -500_00, // This negative amount should be ignored by Total
		},
	}))

	// Output:
	// 100000
	// 300000
	// 100000
}

func ExampleAvg() {
	fmt.Println(Avg([]types.Payment{
		{
			Amount: 1_000_00,
		},
	}))
	fmt.Println(Avg([]types.Payment{
		{
			Amount: 1_000_00,
		},
		{
			Amount: 2_000_00,
		},
	}))
	fmt.Println(Avg([]types.Payment{
		{
			Amount: 1_000_00,
		},
		{
			Amount: 2_000_00,
		},
		{
			Amount: 3_000_00,
		},
	}))

	// Output:
	// 100000
	// 150000
	// 200000
}

func ExampleTotalInCategory() {
	// Пример 1: один платеж в категории
	fmt.Println(TotalInCategory([]types.Payment{
		{
			Amount:   1_000_00,
			Category: "Food",
		},
	}, "Food"))

	// Пример 2: несколько платежей, все в нужной категории
	fmt.Println(TotalInCategory([]types.Payment{
		{
			Amount:   1_000_00,
			Category: "Food",
		},
		{
			Amount:   2_000_00,
			Category: "Food",
		},
	}, "Food"))

	// Пример 3: несколько платежей, только некоторые в нужной категории
	fmt.Println(TotalInCategory([]types.Payment{
		{
			Amount:   1_000_00,
			Category: "Food",
		},
		{
			Amount:   2_000_00,
			Category: "Transport",
		},
	}, "Food"))

	// Пример 4: платежи в категории, но с отрицательными суммами
	fmt.Println(TotalInCategory([]types.Payment{
		{
			Amount:   -500_00,
			Category: "Food",
		},
		{
			Amount:   2_000_00,
			Category: "Food",
		},
	}, "Food"))

	// Пример 5: нет платежей в категории
	fmt.Println(TotalInCategory([]types.Payment{
		{
			Amount:   1_000_00,
			Category: "Transport",
		},
	}, "Food"))

	// Output:
	// 100000
	// 300000
	// 100000
	// 200000
	// 0
}
