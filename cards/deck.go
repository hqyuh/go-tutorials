package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// thêm vào trước để hàm main gọi được func này
func (d deck) print() {
	for i, value := range d {
		fmt.Println(i, value)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	// []string(d)
	return strings.Join(d, ", ")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 06666)
}

func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ", ")
	return s
}

func (d deck) shuffle() {
	// v1
	//for i := range d {
	//	newPosition := rand.Intn(len(d) - 1)
	//	// swap
	//	d[i], d[newPosition] = d[newPosition], d[i]
	//}
	//

	// v2
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		// swap
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
