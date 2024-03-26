package checkout

type checkoutItem struct {
	itemPricing
	quantity int
}

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
