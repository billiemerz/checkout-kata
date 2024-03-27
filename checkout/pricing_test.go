package checkout

import (
	"errors"
	"reflect"
	"testing"
)

func TestGetItemPricing(t *testing.T) {
	tTable := []struct {
		name        string
		sku         string
		expected    itemPricing
		expectedErr error
	}{
		{
			name:        "valid item",
			sku:         "A",
			expected:    itemPricing{50, 3, 130},
			expectedErr: nil,
		},
		{
			name:        "invalid item",
			sku:         "C",
			expected:    itemPricing{},
			expectedErr: errors.New("item not found in pricing schema"),
		},
	}

	pricingSchemaFile = "testdata/valid_pricing.json"

	for _, tt := range tTable {
		t.Run(tt.name, func(t *testing.T) {
			pricingSchemaFile = "testdata/valid_pricing.json"
			ip, err := getItemPricing(tt.sku)

			if !reflect.DeepEqual(ip, tt.expected) {
				t.Errorf("Expected item pricing: %v, got: %v", tt.expected, ip)
			}

			if !ErrorContains(err, tt.expectedErr) {
				t.Errorf("expected %v, got %v", tt.expectedErr, err)
			}
		})
	}

}

func TestLoadPricingSchema(t *testing.T) {
	tTable := []struct {
		name            string
		priceSchemaFile string
		expected        pricingSchema
		expectedErr     error
	}{
		{
			name:            "valid schema",
			priceSchemaFile: "testdata/valid_pricing.json",
			expected: pricingSchema{
				"A": itemPricing{
					UnitPrice:     50,
					OfferQuantity: 3,
					OfferPrice:    130,
				},
				"B": itemPricing{
					UnitPrice: 30,
				},
			},
			expectedErr: nil,
		},
		{
			name:            "invalid schema",
			priceSchemaFile: "testdata/invalid_pricing.json",
			expected:        nil,
			expectedErr:     errors.New("error decoding pricing file"),
		},
		{
			name:            "missing schema",
			priceSchemaFile: "testdata/missing_pricing.json",
			expected:        nil,
			expectedErr:     errors.New("error opening pricing file"),
		},
	}

	for _, tt := range tTable {
		t.Run(tt.name, func(t *testing.T) {
			pricingSchemaFile = tt.priceSchemaFile
			schema, err := loadPricingSchema()

			if !reflect.DeepEqual(schema, tt.expected) {
				t.Errorf("Expected schema: %v, got: %v", tt.expected, schema)
			}

			if !ErrorContains(err, tt.expectedErr) {
				t.Errorf("expected %v, got %v", tt.expectedErr, err)
			}
		})
	}
}
