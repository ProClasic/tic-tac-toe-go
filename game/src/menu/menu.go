package menu

import (
	"fmt"
	"ttt/src/game"
	"ttt/src/utils"

	"github.com/nsf/termbox-go"
)

var (
	menuCursor int
	menuOpts   [3]string = [3]string{
		"host new game",
		"join game",
		"exit",
	}
)

func renderMenu() {
	utils.Clrscr()
	fmt.Print("             TIC-TAC-TOE\n\n")
	fmt.Printf("Hello %s\n\n", game.Username)
	for idx, opt := range menuOpts {
		if idx == menuCursor {
			fmt.Print("-> ")
		} else {
			fmt.Print("   ")
		}
		fmt.Print(opt, "\n")
	}
}

func listenKeyboard(optChan chan int) {
	termbox.SetInputMode(termbox.InputEsc)

	for {
		switch e := termbox.PollEvent(); e.Type {
		case termbox.EventKey:
			switch e.Key {
			case termbox.KeyArrowDown:
				menuCursor++
			case termbox.KeyArrowUp:
				menuCursor--
			case termbox.KeyEnter:
				optChan <- menuCursor
				return
			}
			if l := len(menuOpts); menuCursor < 0 {

				menuCursor = l - 1
			} else if menuCursor >= l {
				menuCursor = 0
			}
			renderMenu()
		}
	}
}

func Start() {
	fmt.Print("enter your name: ")
	fmt.Scanln(&game.Username)
	if err := termbox.Init(); err != nil {
		panic(err)
	}

	renderMenu()

	optChan := make(chan int)

	go listenKeyboard(optChan)

	menuCursor := <-optChan

	termbox.Close()
	utils.Clrscr()

	switch menuCursor {
	case 0:
		game.Host()
	case 1:
		game.Join()
	case 2:
		fmt.Println("thanks buddy, very cool.")
		fmt.Println("exiting...")
		return
	}
}
