package main

func main() {
	cards := deck{"Ace of cards", getDiamonds()}
	cards = append(cards, "Ace of Spades")

	cards.print()
}

func getDiamonds() string {
	return "Ace of Diamonds"
}
