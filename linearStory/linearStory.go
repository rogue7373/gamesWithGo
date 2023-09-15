package main

import "fmt"

type storyPage struct {
	text     string
	nextPage *storyPage
}

func (page *storyPage) playStory() {
	for page != nil {
		fmt.Println(page.text)
		page = page.nextPage
	}

}

// Linked lists
func (page *storyPage) addToEnd(text string) {
	for page.nextPage != nil {
		page = page.nextPage

	}
	page.nextPage = &storyPage{text, nil}

}

// Add new node
func (page *storyPage) addAfter(text string) {
	newPage := &storyPage{text, page.nextPage}
	page.nextPage = newPage
}

// Delete to be added (coming back to this)

func main() {

	page1 := storyPage{"It was a dark and stormy night.", nil}
	page1.addToEnd("You are alone, and you need to fin the sacred helmet before the bad guys do")
	page1.addToEnd("You see a troll ahead")

	// Testing adding a page after the initial page 1
	page1.addAfter("Testing addAfter")

	// Functions - has a return value - may also execute commands
	// Procedures - has no return value, just executes commands
	// Method call - Functions that are attached to a struct/object/etc

	page1.playStory()

}

// Games with Go Youtube Series - Ep. 04 "gameswithgo.org"
