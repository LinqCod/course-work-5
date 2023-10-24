package test

import (
	"github.com/linqcod/course-work-5/app/internal/model"
	"testing"
)

func Test_ProfitabilityByIdForDeposit(t *testing.T) {
	t.Run("test for one deposit", func(t *testing.T) {
		var deposit model.Deposit
		deposit.PercentageRate = 7
		deposit.InitialAmount = 120
		deposit.NumberOfMonths = 6

		actualResult, _ := deposit.ProfitabilityById(deposit.NumberOfMonths)
		var expectedResult float32 = 3.5

		if actualResult != expectedResult {
			t.Errorf("profit = %v, and expected result = %v", actualResult, expectedResult)
		}
	})
}

func Test_ProfitabilityByIdForShare(t *testing.T) {
	t.Run("test for one share", func(t *testing.T) {
		var share model.Share
		share.PurchasePrice = 324.4
		share.EstimatedSellingPrice = 424.4
		share.ExpectedAmountOfDividends = 50

		actualResult, _ := share.ProfitabilityById(12)
		var expectedResult float32 = 46.239212

		if actualResult != expectedResult {
			t.Errorf("profit = %v, and expected result = %v", actualResult, expectedResult)
		}
	})
}

func Test_ProfitabilityByIdForBond(t *testing.T) {
	t.Run("test for one bond", func(t *testing.T) {
		var bond model.Bond
		bond.AmountOfMonths = 24
		bond.SizeOfCoupon = 50
		bond.NumberOfPayments = 2
		bond.PurchasePrice = 960
		bond.Nominal = 1000

		actualResult, _ := bond.ProfitabilityById(24)
		var expectedResult float32 = 7.2916665

		if actualResult != expectedResult {
			t.Errorf("profit = %v, and expected result = %v", actualResult, expectedResult)
		}
	})
}
