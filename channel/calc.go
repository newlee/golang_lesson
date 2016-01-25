package taximeter

import (
	"fmt"
	"strconv"
)

const (
	perMilePrice float64 = 1.5
	startPrice   int     = 6
)

func StartWorker(distanceChan chan int, waitingChan chan int, resultChan chan int) {
	priceChan := make(chan float64)
	go distancePrice(distanceChan, priceChan)
	go waitingPrice(waitingChan, priceChan)
	go roundPrice(priceChan, resultChan)
}

func distancePrice(distanceChan chan int, priceChan chan float64) {
	for {
		distance := <-distanceChan
		priceChan <- startCalc(distance) + mileCalc(distance) + overCalc(distance)
	}
}

func waitingPrice(waitingChan chan int, priceChan chan float64) {
	for {
		perMinsPrice := 0.25
		waitingTime := <-waitingChan
		priceChan <- float64(waitingTime) * perMinsPrice
	}
}

func roundPrice(priceChan chan float64, resultChan chan int) {
	for {
		price := <-priceChan + <-priceChan
		roundPrice := fmt.Sprintf("%.0f", price)
		priceNumber, _ := strconv.ParseInt(roundPrice, 10, 0)
		resultChan <- int(priceNumber)
	}
}

func startCalc(distance int) float64 {
	return float64(startPrice)
}

func mileCalc(distance int) float64 {
	startDistance := 2
	if distance <= startDistance {
		return 0
	}

	return float64(distance-startDistance) * perMilePrice
}

func overCalc(distance int) float64 {
	overDistance := 8
	if distance <= overDistance {
		return 0
	}

	return float64(distance-overDistance) * perMilePrice * 0.5
}
