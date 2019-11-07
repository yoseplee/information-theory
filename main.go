package main //import "github.com/yoseplee/information-theory/homework2"

import (
	"fmt"

	"github.com/yoseplee/information-theory/wep"
)

func main() {

	var c1 = wep.Codeword{
		Idx:     1,
		Symbol:  "00",
		Encoded: "000",
		Length:  4,
	}

	var c2 = wep.Codeword{
		Idx:     2,
		Symbol:  "01",
		Encoded: "101",
		Length:  4,
	}

	var c3 = wep.Codeword{
		Idx:     3,
		Symbol:  "10",
		Encoded: "110",
		Length:  4,
	}

	var c4 = wep.Codeword{
		Idx:     4,
		Symbol:  "11",
		Encoded: "111",
		Length:  4,
	}

	codewordArr := [4]wep.Codeword{
		c1, c2, c3, c4,
	}

	lambda := wep.CalculateLambda(codewordArr[:], 0.1)

	for key, val := range lambda {
		fmt.Println(key.GetId(), " ", val)
	}

	fmt.Println("Word Error Probability")

	for i := 1; i < 6; i += 1 {
		crit := float64(i) * 0.1
		fmt.Println("e: 0.", i, " ", wep.CalculateWordErrorProbability(codewordArr[:], lambda, crit))
	}
}
