package checkout

import "fmt"

type ICheckout interface {
	Scan(string) error
	GetTotalPrice() int
}

/*
Checkout represents a basket of items, it implements the ICheckout interface.
*/
type Checkout struct {
	basket map[string]*checkoutItem
}

func NewCheckout() Checkout {
	return Checkout{
		basket: make(map[string]*checkoutItem),
	}
}

/*
GetTotalPrice returns the summed price of all items in the basket.
*/
func (c Checkout) GetTotalPrice() int {
	total := 0
	for _, item := range c.basket {
		total += item.getPrice()
	}
	return total
}

func (c Checkout) Scan(sku string) error {

	item, found := c.basket[sku]

	if found {
		item.quantity++
	} else {
		itemPricing, err := getItemPricing(sku)

		if err != nil {
			return fmt.Errorf("error getting item pricing: %w", err)
		}

		c.basket[sku] = &checkoutItem{
			quantity:    1,
			itemPricing: itemPricing,
		}
	}

	return nil
}
