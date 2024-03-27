package checkout

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

/*
PricingSchemaFile is the name of the file that contains the pricing information of items.

package level variable for testing purposes.
*/
var pricingSchemaFile string

/*
init sets the pricingSchemaFile to the correct path.
*/
func init() {
	pricingSchemaFile = path.Join(rootDir(), "pricing.json")
}

/*
itemPricing struct is used to store the pricing information of an item.
*/
type itemPricing struct {
	UnitPrice     int `json:"unitPrice"`
	OfferQuantity int `json:"offerQuantity"`
	OfferPrice    int `json:"offerPrice"`
}

/*
getItemPricing returns the pricing information for an item based on its SKU,
this will be called if the item does not already exist in the basket.

If the item is not found in the pricing schema, an error is returned.

This function could be replaced with a cache lookup/ db query/ call to a
separate service in the real world.
*/
func getItemPricing(sku string) (itemPricing, error) {
	pricingSchema, err := loadPricingSchema()

	if err != nil {
		return itemPricing{}, fmt.Errorf("error loading pricing schema: %w", err)
	}

	itemPricing, found := pricingSchema[sku]
	if !found {
		return itemPricing, fmt.Errorf("item not found in pricing schema")
	}

	return itemPricing, nil
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
