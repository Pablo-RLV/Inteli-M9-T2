package publisher

import (
	"math/rand"
	"strconv"
)

func RandFreezer(max, min int) string {
	var max_limit = max + 10
	var min_limit = min - 10
	var sorted = rand.Intn(max_limit-min_limit) + min_limit
	if sorted < min {
		return strconv.Itoa(-sorted) + "° [ALERTA: Temperatura ALTA]"
	}
	if sorted > max {
		return strconv.Itoa(-sorted) + "° [ALERTA: Temperatura BAIXA]"
	}
	return strconv.Itoa(-sorted) + "°"
}

func RandGeladeira(max, min int) string {
	var max_limit = max + 10
	var min_limit = min - 10
	var sorted = rand.Intn(max_limit-min_limit) + min_limit
	if sorted < min { 
		return strconv.Itoa(sorted) + "° [ALERTA: Temperatura BAIXA]"
	}
	if sorted > max {
		return strconv.Itoa(sorted) + "° [ALERTA: Temperatura ALTA]"
	}
	return strconv.Itoa(sorted) + "°"
}
