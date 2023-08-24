package vars

import (
	"fmt"
	"time"
)

var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func SetOtherVariables() {
	Nombre = "Pedro"
	Estado = true
	Sueldo = 1522.2
	Fecha = time.Now()
}

func ShowOtherVariables() {
	fmt.Println("Nombre: ", Nombre)
	fmt.Println("Estado: ", Estado)
	fmt.Println("Sueldo: ", Sueldo)
	fmt.Println("Hoy es: ", Fecha.Day(), " de", Fecha.Month(), " del", Fecha.Year())
}
