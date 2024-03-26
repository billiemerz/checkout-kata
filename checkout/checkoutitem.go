package checkout

/*
checkoutItem stores data about an item with a given SKU in the basket.
This includes the pricing information and the quantity of the item.
*/
type checkoutItem struct {
	itemPricing
	quantity int
}

/*
getPrice returns the total price of the item based on the quantity and pricing information.

If the item has an offer, the price is calculated as follows:
- Add items at offer price (quotient of (quantity/ offer quantity) * offer price)
- Add remaining items at unit price (remainder of (quantity/ offer quantity) * unit price)

If the item has no offer, the price is calculated as:
- Add all items at unit price
*/
func (c checkoutItem) getPrice() int {

	total := 0

	// only apply offer if OfferQuantity is greater than 0
	if c.OfferQuantity > 0 {
		// add items at offer price (quotient of (quantity/ offer quantity) * offer price)
		total += (c.quantity / c.OfferQuantity) * c.OfferPrice

		// add remaining items at unit price (remainder of (quantity/ offer quantity) * unit price)
		total += (c.quantity % c.OfferQuantity) * c.UnitPrice
	} else {
		// no offer, add items at unit price
		total += c.quantity * c.UnitPrice
	}

	return total
}
