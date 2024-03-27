package checkout

import (
	"errors"
	"testing"
)

func TestCheckout_GetTotalPrice(t *testing.T) {
	tests := []struct {
		name             string
		startingCheckout Checkout
		expectedTotal    int
	}{
		{
			name: "Empty basket",
			startingCheckout: Checkout{
				basket: make(map[string]*checkoutItem),
			},
			expectedTotal: 0,
		},
		{
			name: "Multiple items in basket, 1 using offer, 1 not using offer",
			startingCheckout: Checkout{
				basket: map[string]*checkoutItem{
					"A": {
						itemPricing: itemPricing{
							UnitPrice:     50,
							OfferQuantity: 3,
							OfferPrice:    130,
						},
						quantity: 2,
					},
					"B": {
						itemPricing: itemPricing{
							UnitPrice:     30,
							OfferQuantity: 2,
							OfferPrice:    45,
						},
						quantity: 3,
					},
				},
			},
			expectedTotal: 175,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.startingCheckout
			total := c.GetTotalPrice()

			if total != tt.expectedTotal {
				t.Errorf("expected total price: %d, got: %d", tt.expectedTotal, total)
			}
		})
	}
}

func TestCheckout_Scan(t *testing.T) {
	tests := []struct {
		name             string
		sku              string
		expected         int
		startingCheckout Checkout
		expectedErr      error
	}{
		{
			name:        "Item already in basket",
			sku:         "A",
			expected:    2,
			expectedErr: nil,
			startingCheckout: Checkout{
				basket: map[string]*checkoutItem{
					"A": {
						itemPricing: itemPricing{
							UnitPrice:     50,
							OfferQuantity: 3,
							OfferPrice:    130,
						},
						quantity: 1,
					},
				},
			},
		},
		{
			name:             "Item not in basket",
			sku:              "B",
			expected:         1,
			expectedErr:      nil,
			startingCheckout: NewCheckout(),
		},
		{
			name:             "Item not in basket or pricing schema",
			sku:              "D",
			expected:         0,
			expectedErr:      errors.New("error getting item pricing"),
			startingCheckout: NewCheckout(),
		},
	}

	pricingSchemaFile = "testdata/valid_pricing.json"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.startingCheckout
			err := c.Scan(tt.sku)

			if i, ok := c.basket[tt.sku]; ok && i.quantity != tt.expected {
				t.Errorf("expected quantity: %d, got: %d", tt.expected, c.basket[tt.sku].quantity)
			}

			if !ErrorContains(err, tt.expectedErr) {
				t.Errorf("expected %v, got %v", tt.expectedErr, err)
			}
		})
	}
}
