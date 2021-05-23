package main

import "fmt"

func main() {
	animalWeightMap := make(map[string]int)
	animalWeightMap["panda"] = 100;
	animalWeightMap["monkey"] = 30;
	animalWeightMap["bird"] = 2;

	for animal, weight := range animalWeightMap {
		fmt.Printf("animal:%s, weight: %d\n", animal, weight)
	}

	fmt.Println("---------------------------------------");

	for animal := range animalWeightMap {
		fmt.Printf("aninal:%s, weight: %d\n", animal, animalWeightMap[animal])
	}
	fmt.Println("---------------------------------------");

	var animal = "panda"
	weight, ok := animalWeightMap[animal]

	if ok {
		fmt.Println("animal:", animal, "weight:", weight);
	}

}
