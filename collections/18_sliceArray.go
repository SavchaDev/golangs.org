package collections

import "fmt"

type newPlanets []string

func (terra newPlanets) varyPlanet() {
	for i, val := range terra {
		terra[i] = "new " + val
	}
}

// ChangePlanets изменение списка планет
func ChangePlanets() {
	planets := []string{
		"Меркурий", "Венера", "Земля", "Марс",
		"Юпитер", "Сатурн", "Уран", "Нептун",
	}
	newPlanets(planets[3:4]).varyPlanet()
	newPlanets(planets[6:]).varyPlanet()

	fmt.Print(planets)
}
