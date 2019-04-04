package shopping

import (
	"github.com/nsarychev/coding-problems-go/intro/shopping/db"
)

// PriceCheck yolo
func PriceCheck(itemID int) (float64, bool) {
	item := db.LoadItem(itemID)
	if item == nil {
		return 0, false
	}
	return item.Price, true
}
