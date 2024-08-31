package stats

import "github.com/Yessentemir256/bank/v2/pkg/bank/types"

// TotalInCategory находит сумму покупок в определенной категории.
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := types.Money(0)
	for _, payment := range payments {
		if payment.Status == types.StatusFail {
			continue
		}
		if payment.Category != category {
			continue
		}
		if payment.Amount <= 0 {
			continue
		}
		sum += payment.Amount
	}
	return sum

}

// Total рассчитывает сумму всех платежей.
func Total(payments []types.Payment) types.Money {
	sum := types.Money(0)
	for _, payment := range payments {
		if payment.Status == types.StatusFail {
			continue
		}
		if payment.Amount <= 0 {
			continue
		}
		sum += payment.Amount
	}
	return sum
}

// Avg рассчитывает среднюю сумму платежа.
func Avg(payments []types.Payment) types.Money {
	validPaymentsCount := 0
	sum := types.Money(0)
	for _, payment := range payments {
		if payment.Status == types.StatusFail {
			continue
		}
		if payment.Amount <= 0 {
			continue
		}
		sum += payment.Amount
		validPaymentsCount++
	}
	if validPaymentsCount == 0 {
		return 0
	}
	avg := sum / types.Money(validPaymentsCount)
	return avg
}

// FilterByCategory возвращает платежи в указанной категории.
func FilterByCategory(payments []types.Payment, category types.Category) []types.Payment {
	var filtered []types.Payment
	for _, payment := range payments {
		if payment.Category == category {
			filtered = append(filtered, payment)
		}
	}

	return filtered
}

// CategoriesTotal возвращает сумму платежей по каждой категории.
func CategoriesTotal(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}

	for _, payment := range payments {
		categories[payment.Category] += payment.Amount
	}

	return categories
}

// CategoriesAvg считает среднюю сумму платежа по каждой категории.
func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}

	for _, payment := range payments {
		categories[payment.Category] = TotalInCategory2(payments, payment.Category) / NumberOfPaymentsCategory(payments, payment.Category)
	}

	return categories

}

// TotalInCategory2 находит сумму покупок в определенной категории.
func TotalInCategory2(payments []types.Payment, category types.Category) types.Money {
	sum := types.Money(0)

	for _, payment := range payments {

		if payment.Category != category {
			continue
		}

		sum += payment.Amount
	}
	return sum
}

// NumberOfPaymentsCategory считает количесво категорий в платеже.
func NumberOfPaymentsCategory(payments []types.Payment, category types.Category) types.Money {
	sum := types.Money(0)

	for _, payment := range payments {
		if payment.Category == category {
			sum++
		}

	}
	return sum
}
