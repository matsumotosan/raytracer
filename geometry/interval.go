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

func (i *Interval) Clamp(x float64) float64 {
	if x < i.Min {
		return i.Min
	} else if x > i.Max {
		return i.Max
	} else {
		return x
	}
}
