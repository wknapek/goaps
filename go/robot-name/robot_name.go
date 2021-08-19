package robotname

import (
	"errors"
	"math/rand"
	"strconv"
)

var usedRobots = make(map[string]bool)

const maxNrOfRobots = 676000

type Robot struct {
	name string
}

func (rob *Robot) Name() (string, error) {
	if len(usedRobots) == maxNrOfRobots {
		return "", errors.New("robot names exhausted")
	}
	if rob.name != "" {
		return rob.name, nil
	} else {
		name := generateName()
		for usedRobots[name] {
			name = generateName()
		}
		rob.name = name
		usedRobots[name] = true
		return rob.name, nil
	}
}

func (rob *Robot) Reset() {
	//delete(usedRobots, rob.name)
	rob.name = ""
}

func generateName() string {
	letters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	return string(letters[rand.Intn(len(letters))]) + string(letters[rand.Intn(len(letters))]) + strconv.Itoa(numbers[rand.Intn(len(numbers))]) + strconv.Itoa(numbers[rand.Intn(len(numbers))]) + strconv.Itoa(numbers[rand.Intn(len(numbers))])
}
