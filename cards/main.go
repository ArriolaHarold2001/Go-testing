package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
	print(deal(cards, 3))
}