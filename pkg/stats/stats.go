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
