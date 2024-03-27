package checkout

import (
	"encoding/json"
	"fmt"
	"os"
)

/*
PricingSchemaFile is the name of the file that contains the pricing information of items.

package level variable for testing purposes.
*/
var pricingSchemaFile = "pricing.json"

/*
itemPricing struct is used to store the pricing information of an item.
*/
type itemPricing struct {
	UnitPrice     int `json:"unitPrice"`
	OfferQuantity int `json:"offerQuantity"`
	OfferPrice    int `json:"offerPrice"`
}

func getItemPricing(sku string) (itemPricing, error) {
	return itemPricing{}, nil
}

/*
pricingSchema is a map of SKU to itemPricing.

It is parsed from the pricingSchemaFile.
*/
type pricingSchema map[string]itemPricing

func loadPricingSchema() (pricingSchema, error) {
	file, err := os.Open(pricingSchemaFile)
	if err != nil {
		return nil, fmt.Errorf("error opening pricing file: %w", err)
	}
	defer file.Close()

	var schema pricingSchema
	err = json.NewDecoder(file).Decode(&schema)
	if err != nil {
		return nil, fmt.Errorf("error decoding pricing file: %w", err)
	}

	return schema, nil
}
