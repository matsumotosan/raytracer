package geometry

type Interval struct {
	Min, Max float64
}

func (i *Interval) contains(x float64) bool {
	return i.Min <= x && x <= i.Max
}

func (i *Interval) surrounds(x float64) bool {
	return i.Min < x && x < i.Max
}
