package main

import (
	"fmt"
)

//func NbMonths(startPriceOld, startPriceNew, savingPermonth int, percentLossByMonth float64) [2]int {
//	months := 0
//var loss int

//	if startPriceOld >= startPriceNew {
//		return [2]int{months, startPriceOld - startPriceNew}
//	}
//	for true {
//		months++
//		if months%2 == 0 {
//			percentLossByMonth += 0.5
//		}
//
//		startPriceOld -= int((float64(startPriceOld) * percentLossByMonth) / 100)
//		startPriceNew -= int((float64(startPriceNew) * percentLossByMonth) / 100)
//		if (months*savingPermonth+startPriceOld)-startPriceNew >= 0 {
//			//return [2]int{months, savingPermonth + startPriceOld - startPriceNew}
//			break
//		}
//	}
//	return [2]int{months, math.Round(months*savingPermonth + startPriceOld)) - startPriceNew}
//
//}

func NbMonths(startPriceOld, startPriceNew, savingperMonth int, percentLossByMonth float64) [2]int {
	months := 0
	priceOld := float64(startPriceOld)
	priceNew := float64(startPriceNew)

	for ; priceOld+float64(months*savingperMonth) < priceNew; months++ {
		if months%2 == 1 {
			percentLossByMonth += 0.5
		}
		priceOld -= priceOld * percentLossByMonth / 100.0
		priceNew -= priceNew * percentLossByMonth / 100.0
	}

	return [2]int{months, int(priceOld + float64(months*savingperMonth) - priceNew + 0.5)}
}
func main() {

	fmt.Println(NbMonths(2000, 8000, 1000, 1.5))
}
