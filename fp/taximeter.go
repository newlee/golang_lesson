package taximeter

type Taximeter struct {
	distance     int
	waitingTime  int
	distanceCalc func(distance int) float64
	waitingCalc  func(waitingTime int) float64
	roundCalc    func(price float64) int64
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
	price := t.distanceCalc(t.distance) + t.waitingCalc(t.waitingTime)
	return Price(t.roundCalc(price))
}
