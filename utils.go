package faker

import (
	"math"
	"math/rand"
	"time"
)

func randomChoice(array []string) string {
	rand.Seed(time.Now().UnixNano())
	return array[rand.Intn(len(array))]
}

func randomChoiceNested(array [][]string) []string {
	rand.Seed(time.Now().UnixNano())
	return array[rand.Intn(len(array))]
}

func randomChoiceIndexNested(array [][]string) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(len(array))
}

func randomChoiceIF(array []interface{}) interface{} {
	rand.Seed(time.Now().UnixNano())
	return array[rand.Intn(len(array))]
}

func round(floatNum float64, roundNum int) float64{
	floatRoundNum := float64(roundNum)
	return math.Round(floatNum * floatRoundNum) / floatRoundNum
}