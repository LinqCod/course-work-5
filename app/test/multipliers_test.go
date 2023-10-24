package test

import (
	"github.com/linqcod/course-work-5/app/internal/model"
	"github.com/linqcod/course-work-5/app/internal/service"
	"testing"
)

func Test_CalculateEBITDA(t *testing.T) {
	t.Run("ebitda test", func(t *testing.T) {
		comp := new(model.Company)
		comp.Revenue = 500000.0
		comp.TransactionCosts = 300000.0

		actualResult := service.CalculateEBITDA(*comp)
		var expectedResult float32 = 200000.0

		if actualResult != expectedResult {
			t.Errorf("ebitda = %v, and expected result = %v", actualResult, expectedResult)
		}
	})
}

func Test_CalculatePE(t *testing.T) {
	t.Run("p/e test", func(t *testing.T) {
		comp := new(model.Company)
		comp.MarketCapitalization = 500000.0
		comp.AnnualProfit = 100000.0

		actualResult := service.CalculatePE(*comp)
		var expectedResult float32 = 5

		if actualResult != expectedResult {
			t.Errorf("ebitda = %v, and expected result = %v", actualResult, expectedResult)
		}
	})
}

func Test_CalculateEV(t *testing.T) {
	t.Run("e/v test", func(t *testing.T) {
		comp := new(model.Company)
		comp.MarketCapitalization = 500000
		comp.Debentures = 200000
		comp.AvailableFunds = 50000

		actualResult := service.CalculateEV(*comp)
		var expectedResult float32 = 650000

		if actualResult != expectedResult {
			t.Errorf("ebitda = %v, and expected result = %v", actualResult, expectedResult)
		}
	})
}
