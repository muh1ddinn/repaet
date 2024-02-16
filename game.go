package main

import (
	"fmt"
	"math/rand"
)

type Player struct {
	Name   string
	Chance int
}

type Game struct {
	Number int
	Player Player
}

func (g Game) Newgame(pl Player, p int) Game {

	return Game{
		Number: rand.Intn(p),
		Player: pl,
	}
}

func (p Player) Newplayer(name string, chance int) Player {

	return Player{

		Name:   name,
		Chance: chance,
	}

}

func getPlayername() string {

	var Playername string
	fmt.Println("please enter your name : ")
	_, err := fmt.Scanln(&Playername)
	if err != nil {
		panic(err)
	}
	return Playername

}

func (g Game) Stargame() {

	var numb int
	chance := g.Player.Chance

	for i := 1; i <= chance; i++ {

		fmt.Printf("Your %v try. Please enter the number: ", i)
		_, err := fmt.Scanln(&numb)
		if err != nil {
			fmt.Println("Please enter a valid number.")
			continue
		}
		if g.Number == numb {
			fmt.Println("you are win ")
			break
		} else if i == int(chance) {
			fmt.Println("game over you have lose:: okay man do not sad cm on man start game again !!!!! so number is : ", g.Number)
		} else {
			g.Player.Chance--
			fmt.Printf("you did not find .you have %v chances.try again: \n", g.Player.Chance)
		}

	}

}

func main() {

	var o, p int
	fmt.Println("game is starting ! ")

	fmt.Println("You can choose chance. When you choose chance n*n, it means if you choose number 4, your random number will be between 0 and 15 (exclusive).")
	fmt.Print("Please enter the number of chances: ")
	_, err := fmt.Scanln(&o)
	if err != nil || o <= 0 {
		fmt.Println("Please enter a valid positive number for the chance.")
		return
	}
	p = o * o

	pll := Player{}
	pll = pll.Newplayer(getPlayername(), o)

	gamee := Game{}
	gamee = gamee.Newgame(pll, p)
	gamee.Stargame()
}
