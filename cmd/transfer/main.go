package main

import (
	"fmt"
	"github.com/artrey/bgo-adv-methods/pkg/card"
	"github.com/artrey/bgo-adv-methods/pkg/transaction"
	"github.com/artrey/bgo-adv-methods/pkg/transfer"
	"math"
)

func main() {
	transactionSvc := transaction.NewService()
	cardSvc := card.NewService("Tinkoff")
	cardSvc.Issue("visa", 2000_00, "RUB", "0001", "...")
	transferSvc := transfer.NewService(cardSvc, transactionSvc, transfer.Commissions{
		FromInner: func(val int64) int64 {
			return int64(math.Max(float64(val*5/1000), 10_00))
		},
		ToInner: func(val int64) int64 {
			return 0
		},
		FromOuterToOuter: func(val int64) int64 {
			return int64(math.Max(float64(val*15/1000), 30_00))
		},
	})
	fmt.Println(transferSvc)
	fmt.Println(transferSvc.Card2Card("0001", "0002", 1000_00))
	fmt.Println(transactionSvc)
}
