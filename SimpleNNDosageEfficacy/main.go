// This is a very simple program of Neural Network
// It is about predicting the efficacy of a drug dose. It has been found that when the dosage is low efficacy is low.
// Also if the dosage is high efficy is low. However with medium dose efficy is high.
// For simplicity we consider high efficacy as 1 and low efficacy as 0
// This NN has only one hidden layer
// Activation function used here is Soft Plus.
// For full story go through - https://www.youtube.com/watch?v=CqOfi41LfDw&list=PLblh5JKOoLUIxGDQs4LFFD--41Vzf-ME1&index=2
package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

// One layer neural network with sigmoid function as activation function

type Coordinate struct {
	input float64
	x     float64
	y     float64
}

func main() {
	dosage := []float64{0.0, 0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1}
	var finalElementsTop []Coordinate
	var finalElementsBottom []Coordinate

	fmt.Println("Top flow - Top Neuron 1 + Top Neuron 2")
	//Top Neuron-1
	elementsTop := ApplyActivationFuncWithWeightAndBias(dosage, -34.4, 2.14, "SOFT_PLUS")
	for _, elem := range elementsTop {
		fmt.Print(elem)
	}
	//Top Neuron-2
	fmt.Println("\nScaled values")
	for _, elem := range ScaleY(elementsTop, -1.30) {
		fmt.Print(elem)
		finalElementsTop = append(finalElementsTop, elem)
	}

	fmt.Println("Bottom flow - Bottom Neuron 1 + Bottom Neuron 2")
	elementsBot := ApplyActivationFuncWithWeightAndBias(dosage, -2.52, 1.29, "SOFT_PLUS")
	for _, elem := range elementsBot {
		fmt.Print(elem)
	}
	fmt.Println("\nScaled values")
	for _, elem := range ScaleY(elementsBot, 2.28) {
		fmt.Print(elem)
		finalElementsBottom = append(finalElementsBottom, elem)
	}

	//Combine to find our efficacy vs dosage
	fmt.Println("")
	for i := 0; i < len(finalElementsTop) && i < len(finalElementsBottom); i++ {
		fmt.Println("dose:", finalElementsTop[i].input, "efficacy:", finalElementsTop[i].y+finalElementsBottom[i].y-0.58)
	}

}

// This function takes a list of floats, weight and bias, along with activation function type
func ApplyActivationFuncWithWeightAndBias(inputList []float64, weight float64, bias float64, ActivationFunction string) []Coordinate {

	var c []Coordinate
	for _, i := range inputList {
		x := math.Round((i*weight+bias)*100) / 100
		err, y := ApplyActivatioNfunction(x, ActivationFunction)
		if err != nil {
			fmt.Println("Error found in activation function")
		} else {
			c = append(c, Coordinate{i, x, y})
		}
	}
	return c
}

func ApplyActivatioNfunction(x float64, actType string) (err error, y float64) {

	switch strings.ToUpper(actType) {
	case "SOFT_PLUS":
		return nil, SoftPlusActivationFunc(x)
	default:
		return errors.New("Unimplemented a valid activation function used: " + actType), 0.0
	}

}

func ScaleY(coordinates []Coordinate, scaleFactor float64) (newCorrdinates []Coordinate) {
	var opCoord []Coordinate
	for _, c := range coordinates {
		c.y = math.Round(c.y*scaleFactor*100) / 100
		opCoord = append(opCoord, c)
	}
	return opCoord
}

func SoftPlusActivationFunc(x float64) (y float64) {
	y = math.Round(math.Log(1+math.Exp(x))*100) / 100
	return y
}

