package ejercicios

import "strconv"

func StringToInt(txt string) (int, string) {

	entero, err := strconv.Atoi(txt)
	if err != nil {
		return 0, "Hubo un error"
	}

	if entero > 100 {
		return entero, "es mayor a 100"
	} else {
		return entero, "es menor a 100"
	}

}
