package checkout

import "testing"

func TestCheckoutItem_GetPrice(t *testing.T) {
	tests := []struct {
		name          string
		unitPrice     int
		offerQuantity int
		offerPrice    int
		quantity      int
		expected      int
	}{
		{
			name:          "No offer",
			unitPrice:     10,
			offerQuantity: 0,
			offerPrice:    0,
			quantity:      5,
			expected:      50,
		},
		{
			name:          "With offer, quantity less than offer quantity",
			unitPrice:     7,
			offerQuantity: 6,
			offerPrice:    25,
			quantity:      5,
			expected:      35,
		},
		{
			name:          "With offer, quantity equal to offer quantity",
			unitPrice:     1010,
			offerQuantity: 6,
			offerPrice:    25,
			quantity:      6,
			expected:      25,
		},
		{
			name:          "With offer, quantity greater than offer quantity, remainder < 0.5 * offer quantity",
			unitPrice:     15,
			offerQuantity: 6,
			offerPrice:    25,
			quantity:      7,
			expected:      40,
		},
		{
			name:          "With offer, quantity greater than offer quantity, remainder > 0.5 * offer quantity",
			unitPrice:     8,
			offerQuantity: 6,
			offerPrice:    25,
			quantity:      10,
			expected:      57,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			item := checkoutItem{
				quantity: tt.quantity,
				itemPricing: itemPricing{
					UnitPrice:     tt.unitPrice,
					OfferQuantity: tt.offerQuantity,
					OfferPrice:    tt.offerPrice,
				},
			}

			actualPrice := item.getPrice()
			if actualPrice != tt.expected {
				t.Errorf("Expected price: %d, got: %d", tt.expected, actualPrice)
			}
		})
	}
}
