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

	fmt.Print(c1.ToString())
	fmt.Print(c2.ToString())
	fmt.Print(c3.ToString())
	fmt.Print(c4.ToString())
	fmt.Println()

	lambda := wep.CalculateLambda(codewordArr[:], 0.1)
	for _, val := range lambda {
		fmt.Println(val.ToString())
	}

}
