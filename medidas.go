package figuras

import "fmt"

type Medidas interface {
	medidas() float64
}

func Mostrarmedidas(med Medidas) {
	fmt.Println(med.medidas())
}
