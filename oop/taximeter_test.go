package taximeter

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	meter := &Taximeter{
		distanceCalc: StartCalc{nextCalc: MileCalc{nextCalc: OverCalc{nextCalc: EndCalc{}}}},
		waitingCalc:  WaitCalcImpl{},
		rounder:      RounderImpl{},
	}
	meter.Start()
	meter.Run(1).Price().shouldBe(6, t.Error, 1)
	meter.Run(1).Price().shouldBe(6, t.Error, 2)
	meter.Run(1).Price().shouldBe(8, t.Error, 3)
	meter.Run(7).Price().shouldBe(20, t.Error, 10)
	meter.Wait(1).Price().shouldBe(20, t.Error, 10)
	meter.Wait(2).Price().shouldBe(20, t.Error, 10)
	meter.Wait(2).Price().shouldBe(21, t.Error, 10)
}

func (p Price) shouldBe(result Price, output func(args ...interface{}), distance int) {
	if p != result {
		output(fmt.Printf("when distance is: %v, expect:%v, actual: %v\n", distance, result, p))
	}
}
