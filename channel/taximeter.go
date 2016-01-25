package taximeter

type Taximeter struct {
	distance     int
	waitingTime  int
	distanceChan chan int
	waitingChan  chan int
	resultChan   chan int
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
	t.distanceChan <- t.distance
	t.waitingChan <- t.waitingTime
	price := <-t.resultChan
	return Price(price)
}
