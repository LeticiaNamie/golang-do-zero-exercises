package exercises

import (
	"strconv"
	"strings"
)

func Points(games []string) int {
	var points int

	for i := 0; i < 3; i++ {
		points += CalculatePoints(games[i])
	}

	return points
}

func CalculatePoints(game string) int {
	s := strings.Split(game, ":")

	var pointsX int
	var pointsY int

	pointsX, _ = strconv.Atoi(s[0])
	pointsY, _ = strconv.Atoi(s[1])

	if pointsX > pointsY {
		return 3
	} else if pointsX == pointsY {
		return 1
	}

	return 0
}
