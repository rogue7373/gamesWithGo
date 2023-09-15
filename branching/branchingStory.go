package main

// Branching yes/no game depending on yes/no. Based on binary tree

import (
	"bufio"
	"fmt"
	"os"
)

type storyNode struct {
	text    string
	yesPath *storyNode
	noPath  *storyNode
}

func (node *storyNode) play() {
	fmt.Println(node.text)

	if node.yesPath != nil && node.noPath != nil {
		scanner := bufio.NewScanner(os.Stdin)

		// if the user typed yes if they typed no if they typed something else we will capture with a for loop to redirect them.

		for {
			scanner.Scan()
			answer := scanner.Text()

			if answer == "yes" {
				node.yesPath.play()
				break
			} else if answer == "no" {
				node.noPath.play()
				break
			} else {
				fmt.Println("That answer was not an option! Please answer yes or no.")
			}
		}
	}
}

func main() {
	root := storyNode{"You are at the entrance to a dark cave. Do you want to go in the cave", nil, nil}
	winning := storyNode{"You have won!", nil, nil}
	losing := storyNode{"You have lost!", nil, nil}
	root.yesPath = &losing
	root.noPath = &winning

	root.play()

}
