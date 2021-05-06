package main

func main() {
	cards := newDeck()
	cards.saveToFile("1.txt")
	cards.print()
	cardsFromFile := newDeckFromFile("1.txt")
	cardsFromFile.print()
	cardsFromFile.shuffle()
	cardsFromFile.print()

	/*
		cards := newDeck()
		fmt.Println(cards.toString())
		cards.saveToFile("1.txt")


		hand, remainingCards := deal(cards, 5)

		hand.print()
		remainingCards.print()
	*/
}
