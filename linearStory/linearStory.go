package main

import "fmt"

type storyPage struct {
	text     string
	nextPage *storyPage
}

func (page *storyPage) playStory() {
	if page == nil {
		return
	}
	fmt.Println(page.text)
	page.nextPage.playStory()
}
func main() {

	page1 := storyPage{"It was a dark and stormy night.", nil}
	page2 := storyPage{"You are alone, and you need to fin the sacred helmet before the bad guys do", nil}
	page3 := storyPage{"You see a troll ahead", nil}

	page1.nextPage = &page2
	page2.nextPage = &page3

	playStory(&page1)

}

// Games with Go Youtube Series - Ep. 03.2
