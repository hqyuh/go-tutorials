package main

func main() {
	//cards := newDeck()
	//
	//hand, remainingCards := deal(cards, 5)
	//hand.print()
	//remainingCards.print()
	//
	//fmt.Println(cards.toString())
	//err := cards.saveToFile("my_cards")
	//if err != nil {
	//	return
	//}

	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()
}
