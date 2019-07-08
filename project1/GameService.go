package service

import "strconv"

type Game struct {
	x    int
	y    int
	str1 string
	str2 string
}

func Init(x int, y int, str1 string, str2 string) *Game {
	return &Game{x, y, str1, str2}
}
func (game Game) Judge(num int) string {
	if num%game.x == 0 && num%game.y == 0 {
		return game.str1 + game.str2
	}
	if num%game.x == 0 {
		return game.str1
	} else if num%game.y == 0 {
		return game.str2
	}
	return strconv.Itoa(num)
}
