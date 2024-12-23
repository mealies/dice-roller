package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"sort"
	"strconv"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	dice := flag.String("d", "d6", "Type of dice to roll. Format: dx where X is an Integer.  Default: d6")
	numRoll := flag.Int("n", 1, "The number of die to roll.  Default: 1")
	sum := flag.Bool("s", false, "Get the sum of all the dice rolls")
	advantage := flag.Bool("adv", false, "Roll the dice with advantage")
	disadvantage := flag.Bool("dis", false, "Roll the dice with disadvantage")
	flag.Parse()

	matched, _ := regexp.Match("^d\\d+", []byte(*dice))
	//	fmt.Printf("You chose a %s\n", *dice)
	//fmt.Println(matched)

	if matched == true {
		rolls := rollDice(dice, numRoll)
		printDice(rolls)
		if *sum == true {
			diceSum := sumDice(rolls)
			fmt.Printf("The sum of the was %d\n", diceSum)
		}
		if *advantage == true {
			roll := rollWithAdvantage(rolls)
			fmt.Printf("The roll with advantage was %d\n", roll)
		}
		if *disadvantage == true {
			roll := rollWithDisadvantage(rolls)
			fmt.Printf("The roll with disadvantage was %d\n", roll)
		}
	} else {
		log.Fatalf("Improper format for dice. Format should be dX where X is an Integer")
	}
}

func rollDice(dice *string, times *int) []int {
	var rolls []int
	diceSides := (*dice)[1:]
	d, err := strconv.Atoi(diceSides)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < *times; i++ {
		rolls = append(rolls, rand.Intn(d)+1)
	}
	return rolls
}

func printDice(rolls []int) {

	for i, dice := range rolls {
		fmt.Printf("Roll %d was %d\n", i+1, dice)
	}
}

func sumDice(rolls []int) int {
	sum := 0
	for _, roll := range rolls {
		sum += roll
	}
	return sum
}

func rollWithAdvantage(rolls []int) int {
	sort.Ints(rolls)
	return rolls[len(rolls)-1]
}

func rollWithDisadvantage(rolls []int) int {
	sort.Ints(rolls)
	return rolls[0]
}
