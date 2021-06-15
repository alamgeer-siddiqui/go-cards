package main

// import "fmt"

func main() {
	cards := newDeck()
	// cards.saveToFile("my_files")

	// fmt.Println(cards.toString())
	// cards:=newDeckFromFile("my_cards")
	// cards.print()

	cards.shuffle()
	cards.print()
	
}

// func main(){
// 	cards:= deck{"Ace of diamonds",newCard()}
// 	cards=append(cards, "six of spades")

// 	cards.print()

// }

// func newCard() string{
// 	return "Five of diamonds"
// }
