package figuras

type Cuadrado struct {
	Ancho float64
	Alto  float64
}

func (c *Cuadrado) medidas() float64 {
	return c.Ancho * c.Alto
}
