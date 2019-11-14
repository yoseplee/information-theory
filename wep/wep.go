package wep

import (
	"fmt"
	"math"
	"sort"
)

var (
	SetT        [8]string
	c1          *Codeword
	c2          *Codeword
	c3          *Codeword
	c4          *Codeword
	codewordArr [4]Codeword
)

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

	c1 = &Codeword{
		Idx:     1,
		Symbol:  "00",
		Encoded: "000",
		Length:  4,
	}

	c2 = &Codeword{
		Idx:     2,
		Symbol:  "01",
		Encoded: "101",
		Length:  4,
	}

	c3 = &Codeword{
		Idx:     3,
		Symbol:  "10",
		Encoded: "110",
		Length:  4,
	}

	c4 = &Codeword{
		Idx:     4,
		Symbol:  "11",
		Encoded: "111",
		Length:  4,
	}

	codewordArr = [4]Codeword{
		*c1, *c2, *c3, *c4,
	}
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

func CalculateLambda(codewordArr []Codeword, errorRate float64) map[Codeword][]string {

	lambda := make(map[Codeword][]string)
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
		key := likelihoods[keys[len(keys)-1]]
		lambda[key] = append(lambda[key], val)

	}
	return lambda
}

//Cube returns a value that multiplied three times.
func Cube(target float64) float64 {
	return math.Pow(target, 3)
}

func CalculateWordErrorProbability(codewordArr []Codeword, errorRate float64) float64 {

	fmt.Println("Lambda by the maximum likelihood decoding rule")
	lambda := CalculateLambda(codewordArr[:], errorRate)

	for key, val := range lambda {
		fmt.Println(key.GetId(), " ", val)
	}

	equallyLikelyCodeword := 1 / math.Pow(2, float64(len(codewordArr)))

	var sum float64
	for _, codeword := range codewordArr {
		for key, val := range lambda {
			if codeword.Idx == key.Idx {
				continue
			}
			var thisMiss, thisMatch int
			for _, codes := range val {
				miss, match := codeword.CompareCode(codes)
				thisMiss += miss
				thisMatch += match
			}
			sum += math.Pow(errorRate, float64(thisMiss)) * math.Pow(1-errorRate, float64(thisMatch))
		}
	}

	return sum * equallyLikelyCodeword
}

func CalculateUnionBound(codewordArr []Codeword, errorRate float64) float64 {

	fmt.Println("Lambda by the maximum likelihood decoding rule")
	lambda := CalculateLambda(codewordArr[:], errorRate)

	for key, val := range lambda {
		fmt.Println(key.GetId(), " ", val)
	}

	equallyLikelyCodeword := 1 / math.Pow(2, float64(len(codewordArr)))

	var sum float64
	for _, codeword := range codewordArr {
		for key, val := range lambda {
			if codeword.Idx == key.Idx {
				continue
			}
			var thisMiss, thisMatch int
			for _, codes := range val {
				miss, match := codeword.CompareCode(codes)
				thisMiss += miss
				thisMatch += match
			}
			// sum += math.Pow(errorRate, float64(thisMiss)) * math.Pow(1-errorRate, float64(thisMatch))
			sum += math.Pow(errorRate, float64(thisMiss)) * math.Pow(1-errorRate, float64(thisMatch))
		}
	}

	return sum * equallyLikelyCodeword
}

func CalculateBhattacharyyaBound(codewordArr []Codeword, errorRate float64) float64 {

	fmt.Println("Lambda by the maximum likelihood decoding rule")
	lambda := CalculateLambda(codewordArr[:], errorRate)

	for key, val := range lambda {
		fmt.Println(key.GetId(), " ", val)
	}

	equallyLikelyCodeword := 1 / math.Pow(2, float64(len(codewordArr)))

	var sum float64
	for _, codeword := range codewordArr {
		for key, val := range lambda {
			if codeword.Idx == key.Idx {
				continue
			}
			var thisMiss, thisMatch int
			for _, codes := range val {
				miss, match := codeword.CompareCode(codes)
				thisMiss += miss
				thisMatch += match
			}
			sum += math.Sqrt(math.Pow(1-errorRate, 3) * math.Pow(errorRate, float64(thisMiss)) * math.Pow(1-errorRate, float64(thisMatch)))
		}
	}

	return sum * equallyLikelyCodeword
}

func DoHomework4() {

	fmt.Println()
	fmt.Println("Word Error Probability")

	for i := 0; i < 6; i++ {
		crit := float64(i) * 0.1
		fmt.Printf("e: 0.%d - %f\n", i, CalculateWordErrorProbability(codewordArr[:], crit))
		fmt.Println()
	}
}

func DoHomework5() {
	fmt.Println()
	fmt.Println("Word Error Probability")

	for i := 0; i < 6; i++ {
		crit := float64(i) * 0.1
		fmt.Printf("e: 0.%d - %f\n", i, CalculateBhattacharyyaBound(codewordArr[:], crit))
		fmt.Println()
	}
}
