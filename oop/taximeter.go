package taximeter

type DistanceCalc interface {
	Price(distance int) float64
}

type WaitCalc interface {
	Price(waitingTime int) float64
}

type Rounder interface {
	Price(price float64) int64
}

type Taximeter struct {
	distance     int
	waitingTime  int
	distanceCalc DistanceCalc
	waitingCalc  WaitCalc
	rounder      Rounder
}

type Price int

func (t *Taximeter) Start() {
	t.distance = 0
	t.waitingTime = 0
}

func (t *Taximeter) Run(miles int) *Taximeter {
	t.distance += miles
	return t
}

func (t *Taximeter) Wait(time int) *Taximeter {
	t.waitingTime += time
	return t
}

func (t *Taximeter) Price() Price {
	price := t.distanceCalc.Price(t.distance) + t.waitingCalc.Price(t.waitingTime)
	return Price(t.rounder.Price(price))
}
