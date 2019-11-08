package homework2

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"sort"
	"strconv"
)

type HuffmanCode struct{}

type Entropy struct {
	AlphabetCount map[string]int
	TotalAlphabet int
	Support       []string
	EmpiricalMass map[string]float64
	Entropy       float64
}

func PrintEntropy(entropy Entropy) string {
	var printData string

	printData = "< Alphabet Count >\n"
	for start := 'a'; start <= 'z'; start++ {
		count := entropy.AlphabetCount[string(start)]
		printData += string(start)
		printData += " - "
		printData += strconv.Itoa(count)
		printData += "\n"
	}
	printData += "----------------------------------------------------------------------------------------"
	printData += "\n"

	printData += "< Total Alphabet Count >\n"
	printData += strconv.Itoa(entropy.TotalAlphabet)
	printData += "\n"
	printData += "----------------------------------------------------------------------------------------"
	printData += "\n"

	printData += "< Support >\n"
	sortedSupport := entropy.Support
	sort.Strings(sortedSupport)
	for _, val := range sortedSupport {
		printData += val
		printData += " / "
	}
	printData += "\n"
	printData += "----------------------------------------------------------------------------------------"
	printData += "\n"

	printData += "< Empirical Mass - alphabet order >\n"
	keys := make([]string, len(entropy.EmpiricalMass))
	i := 0
	for key := range entropy.EmpiricalMass {
		keys[i] = key
		i++
	}
	sort.Strings(keys)

	for _, key := range keys {
		printData += key
		printData += " - "
		printData += fmt.Sprintf("%f\n", entropy.EmpiricalMass[key])
	}
	printData += "----------------------------------------------------------------------------------------"
	printData += "\n"

	printData += "< Entropy >\n"
	printData += fmt.Sprintf("%f\n", entropy.Entropy)
	printData += "----------------------------------------------------------------------------------------"
	printData += "\n"

	return printData
}

func CalcEntropy(entropy *Entropy) {
	var result float64
	result = 0.0

	for _, val := range entropy.EmpiricalMass {
		mass := val
		information := math.Log2(1 / mass)
		result += mass * information
	}
	entropy.Entropy = result
}

func CalcEmpiricalMass(entropy *Entropy) {
	empiricalMass := make(map[string]float64)
	// entropy.EmpiricalMass = make(map[string]float64)

	for key, val := range entropy.AlphabetCount {
		empiricalMass[key] = float64(val) / float64(entropy.TotalAlphabet)
	}
	entropy.EmpiricalMass = empiricalMass
}

func MakeSupport(entropy *Entropy) {
	entropy.Support = make([]string, len(entropy.AlphabetCount))
	i := 0
	for val := range entropy.AlphabetCount {
		entropy.Support[i] = val
		i++
	}
}

func IsAlphabetRange(character byte) bool {
	if character < 65 || character > 122 {
		return false
	}
	return true
}

func IsUpperCharacter(character byte) bool {
	if character < 91 && character > 64 {
		return true
	}
	return false
}

func FileOut(dataToWrite string, fileName string) {

	err := ioutil.WriteFile(fmt.Sprint("./", fileName), []byte(dataToWrite), 0644)
	if err != nil {
		log.Panic("file write error!")
	}

	fmt.Println("file out done")
}

func DoHomework2() {
	entireTextData := []string{
		"harryPotter1.txt",
		"harryPotter2.txt",
		"harryPotter3.txt",
		"harryPotter4.txt",
		"harryPotter5.txt",
		"harryPotter6.txt",
		"harryPotter7.txt",
		"harryPotter8.txt",
		"harryPotter9.txt",
		"harryPotter10.txt",
		"harryPotter11.txt",
		"harryPotter12.txt",
		"harryPotter13.txt",
		"harryPotter14.txt",
		"harryPotter15.txt",
		"harryPotter16.txt",
		"harryPotter17.txt",
		"harryPotter18.txt",
	}

	for _, txtName := range entireTextData {
		var wholeText []byte
		var entropy Entropy
		entropy.AlphabetCount = make(map[string]int)

		readFilePath := fmt.Sprint("./homework2/", txtName)
		if textFromFile, err := ioutil.ReadFile(readFilePath); err != nil {
			log.Fatal("file read failed")
		} else {
			wholeText = textFromFile
		}

		for _, character := range wholeText {
			if !IsAlphabetRange(character) {
				continue
			}
			if IsUpperCharacter(character) {
				character += 32
			}

			// alphabet count
			if entropy.AlphabetCount[string(character)] == 0 {
				entropy.AlphabetCount[string(character)] = 1
			} else {
				entropy.AlphabetCount[string(character)]++
			}
			entropy.TotalAlphabet++
		}
		MakeSupport(&entropy)
		CalcEmpiricalMass(&entropy)
		CalcEntropy(&entropy)

		entropyData := PrintEntropy(entropy)
		FileOut(entropyData, fmt.Sprint("result_", txtName))
	}
}
