package wep

import (
	"fmt"
	"math"
	"sort"
)

var SetT [8]string

func init() {
	SetT = [8]string{
		"000",
		"001",
		"010",
		"011",
		"100",
		"101",
		"110",
		"111",
	}
}

type Codeword struct {
	Idx     int
	Symbol  string
	Encoded string
	Length  int
}

type Likelihood struct {
	Y            string
	ThisCodeword Codeword
	Rate         float64
}

func (l Likelihood) ToString() string {
	return fmt.Sprintf("%s | %s - %f", l.Y, l.ThisCodeword.GetId(), l.Rate)
}

func (c *Codeword) ToString() string {
	return fmt.Sprintf("Codeword - Idx: %d, Symbol: %s, Encoded Symbol: %s\n", c.Idx, c.Symbol, c.Encoded)
}

func (c *Codeword) CalculateLikelihoodRate(compareCodeword Codeword, Y string, errorRate float64) float64 {
	cMiss, cMatch := c.CompareCode(Y)
	compareMiss, compareMatch := compareCodeword.CompareCode(Y)

	N := math.Pow(errorRate, float64(cMiss)) * math.Pow(1-errorRate, float64(cMatch))
	D := math.Pow(errorRate, float64(compareMiss)) * math.Pow(1-errorRate, float64(compareMatch))
	result := N / D

	return result
}

//CompareCode compares two args with 3 length of digits
func (c *Codeword) CompareCode(toCompare string) (miss int, match int) {
	myCode := []byte(c.Encoded)
	toCompareCode := []byte(toCompare)

	for idx, val := range myCode {
		if val == toCompareCode[idx] {
			match++
		} else {
			miss++
		}
	}

	return miss, match
}

func (c *Codeword) GetId() string {
	return fmt.Sprintf("C%d", c.Idx)
}

func CalculateLambda(codewordArr []Codeword, errorRate float64) []Likelihood {

	lambda := make([]Likelihood, 0, 8)
	e := 0.1
	for _, val := range SetT {
		likelihoods := make(map[float64]Codeword)
		for _, i := range codewordArr {
			for _, j := range codewordArr {
				if i.Idx == j.Idx {
					continue
				}
				got := i.CalculateLikelihoodRate(j, val, e)
				likelihoods[got] = i
				// fmt.Printf("c%d / c%d for y = %s :: %f\n", i.Idx, j.Idx, val, got)
			}
			// fmt.Println()
		}
		// fmt.Println()
		keys := make([]float64, 0, len(likelihoods))
		for k := range likelihoods {
			keys = append(keys, k)
		}
		sort.Float64s(keys)

		/*
			for _, k := range keys {
				fmt.Printf("%s | C%d - %f\n", val, likelihoods[k].Idx, k)
			}
		*/
		// lambda[keys[len(keys)-1]] = likelihoods[keys[len(keys)-1]]
		tmpLikelihood := Likelihood{
			Y:            val,
			ThisCodeword: likelihoods[keys[len(keys)-1]],
			Rate:         keys[len(keys)-1],
		}
		lambda = append(lambda, tmpLikelihood)
	}
	return lambda
}

//Cube returns a value that multiplied three times.
func Cube(target float64) float64 {
	return math.Pow(target, 3)
}

func CalculateWordErrorProbability(codewordArr []Codeword, lambda []Likelihood, errorRate float64) float64 {
	equallyLikelyCodeword := 1 / math.Pow(2, float64(len(codewordArr)))

	var sum float64

	for _, y := range SetT {
		for _, i := range codewordArr {
			for _, j := range codewordArr {
				if i.Idx == j.Idx {
					continue
				}
			}
		}
	}
	return sum * equallyLikelyCodeword
}
