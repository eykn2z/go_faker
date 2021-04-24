package faker

import (
	"math"
	"math/rand"
	"time"
)

func randomIndex(arrayLength int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(arrayLength)
}

func randomChoice(array []string) string {
	return array[randomIndex(len(array))]
}

func randomChoiceNested(array [][]string) []string {
	return array[randomIndex(len(array))]
}

func randomChoiceIndexNested(array [][]string) int {
	return randomIndex(len(array))
}

func randomChoiceIF(array []interface{}) interface{} {
	return array[randomIndex(len(array))]
}

func round(floatNum float64, roundNum int) float64 {
	floatRoundNum := float64(roundNum)
	return math.Round(floatNum*floatRoundNum) / floatRoundNum
}
