package exercises

func Rps(p1, p2 string) string {
	p1Winner := "Player 1 won!"
	p2Winner := "Player 2 won!"
	draw := "Draw!"

	if p1 == p2 {
		return draw
	} else if p1 == "scissors" && p2 == "paper" {
		return p1Winner
	} else if p1 == "paper" && p2 == "rock" {
		return p1Winner
	} else if p1 == "rock" && p2 == "scissors" {
		return p1Winner
	}

	return p2Winner
}
