package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	return float32(taux(balance))
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return taux(balance) * balance / 100
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance:
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	var year = 0
	for balance < targetBalance {
		balance = AnnualBalanceUpdate(balance)
		year++
	}
	return year
}

func taux(solde float64) float64 {
	switch {
	case solde < 0:
		return 3.213
	case solde < 1000:
		return 0.5
	case solde < 5000:
		return 1.621
	default:
		return 2.475
	}
}
