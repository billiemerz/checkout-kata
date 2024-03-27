package checkout_test

import (
	"testing"

	"github.com/billiemerz/checkout-kata/checkout"
)

/*
Tests the public API of the checkout package.
*/

func TestExampleCheckout(t *testing.T) {
	tTable := []struct {
		name      string
		skus      []string
		expected  int
		numErrors int
	}{

		{
			name:      "Empty basket",
			skus:      []string{},
			expected:  0,
			numErrors: 0,
		},
		{
			name:      "all of the same item",
			skus:      []string{"A", "A", "A", "A", "A"},
			expected:  230,
			numErrors: 0,
		},
		{
			name:      "mix of items, offers and no offers",
			skus:      []string{"A", "A", "A", "A", "B", "C"},
			expected:  230,
			numErrors: 0,
		},
		{
			name:      "mix of items, some don't exist",
			skus:      []string{"Banana", "A", "A", "A", "B", "C", "E", "Chicken", "A"},
			expected:  230,
			numErrors: 3,
		},
	}

	for _, tt := range tTable {
		t.Run(tt.name, func(t *testing.T) {
			c := checkout.NewCheckout()

			numErrors := 0

			for _, sku := range tt.skus {
				err := c.Scan(sku)
				if err != nil {
					numErrors++
				}
			}

			total := c.GetTotalPrice()

			if total != tt.expected {
				t.Errorf("expected total price: %d, got: %d", tt.expected, total)
			}

			if numErrors != tt.numErrors {
				t.Errorf("expected number of errors: %d, got: %d", tt.numErrors, numErrors)
			}
		})
	}
}
