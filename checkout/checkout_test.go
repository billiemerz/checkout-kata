package checkout

import "testing"

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
