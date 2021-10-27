package figuras

type Circulo struct {
	Radio float64
}

func (c *Circulo) medidas() float64 {
	return c.Radio * c.Radio
}
