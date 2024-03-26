package checkout

/*
itemPricing struct is used to store the pricing information of an item.

It is parsed from the pricing.json file.
*/
type itemPricing struct {
	UnitPrice     int `json:"unitPrice"`
	OfferQuantity int `json:"offerQuantity"`
	OfferPrice    int `json:"offerPrice"`
}
