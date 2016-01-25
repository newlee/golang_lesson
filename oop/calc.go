package taximeter

import (
	"fmt"
	"strconv"
)

type WaitCalcImpl struct {
}

type RounderImpl struct {
}

func (r RounderImpl) Price(price float64) int64 {
	roundPrice := fmt.Sprintf("%.0f", price)
	priceNumber, _ := strconv.ParseInt(roundPrice, 10, 0)
	return priceNumber
}

func (c WaitCalcImpl) Price(waitingTime int) float64 {
	perMinsPrice := 0.25
	return float64(waitingTime) * perMinsPrice
}

type StartCalc struct {
	nextCalc DistanceCalc
}

func (c StartCalc) Price(distance int) float64 {
	startPrice := 6
	startDistance := 2
	return float64(startPrice) + c.nextCalc.Price(distance-startDistance)
}

type MileCalc struct {
	nextCalc DistanceCalc
}

func (c MileCalc) perMilePrice() float64 {
	return 1.5
}
func (c MileCalc) Price(distance int) float64 {
	if distance <= 0 {
		return 0
	}

	return float64(distance)*c.perMilePrice() + c.nextCalc.Price(distance)
}

type OverCalc struct {
	MileCalc
	nextCalc DistanceCalc
}

func (c OverCalc) Price(distance int) float64 {
	overDistance := 6
	if distance <= overDistance {
		return 0
	}

	return float64(distance-overDistance)*c.perMilePrice()*0.5 + c.nextCalc.Price(distance)
}

type EndCalc struct {
	nextCalc DistanceCalc
}

func (c EndCalc) Price(distance int) float64 {
	return 0
}
