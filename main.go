package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/billiemerz/checkout-kata/checkout"
)

func main() {

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	checkout := checkout.NewCheckout()

	msg := make(chan string, 1)
	go func() {
		for {
			var s string
			fmt.Scan(&s)
			msg <- strings.ToUpper(s)
		}
	}()

	fmt.Println("Enter SKU to scan, 'TOTAL' to get total, or 'EXIT' / SIGTERM (CTRL/ CMD + c) to exit")

	for {
		select {
		case <-sigs:
			fmt.Printf("exiting, total cart val: %d\n", checkout.GetTotalPrice())
			return

		case s := <-msg:
			switch s {
			case "EXIT":
				fmt.Printf("exiting, total cart val: %d\n", checkout.GetTotalPrice())
				return
			case "TOTAL":
				fmt.Printf("total cart val: %d\n", checkout.GetTotalPrice())
			default:
				err := checkout.Scan(s)
				if err != nil {
					fmt.Println(err)
					continue
				}
				fmt.Printf("scanned: %s\n", s)
			}
		}
	}
}
