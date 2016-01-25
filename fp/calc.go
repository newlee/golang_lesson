package taximeter

import (
	"fmt"
	"strconv"
)

const (
	perMilePrice float64 = 1.5
	startPrice   int     = 6
)

var DistancePrice = func(distance int) float64 {
	return startCalc(distance) + mileCalc(distance) + overCalc(distance)
}

var WaitingPrice = func(waitingTime int) float64 {
	perMinsPrice := 0.25
	return float64(waitingTime) * perMinsPrice
}

var RoundPrice = func(price float64) int64 {
	roundPrice := fmt.Sprintf("%.0f", price)
	priceNumber, _ := strconv.ParseInt(roundPrice, 10, 0)
	return priceNumber
}

var startCalc = func(distance int) float64 {
	return float64(startPrice)
}

var mileCalc = func(distance int) float64 {
	startDistance := 2
	if distance <= startDistance {
		return 0
	}

	return float64(distance-startDistance) * perMilePrice
}

var overCalc = func(distance int) float64 {
	overDistance := 8
	if distance <= overDistance {
		return 0
	}

	return float64(distance-overDistance) * perMilePrice * 0.5
}
