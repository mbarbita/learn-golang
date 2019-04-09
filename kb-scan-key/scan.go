package main

import "fmt"
import term "github.com/nsf/termbox-go"

func ScanKey() {
	var response int
	for {
		fmt.Println("This program will activate SkyNet worldwide, are you sure about this?")

		fmt.Scanf("%c", &response) //<--- here
		switch response {
		default:
			fmt.Println("SkyNet launch aborted!")
		case 'y':
			fmt.Println("SkyNet launched")
		case 'Y':
			fmt.Println("SkyNet launched")
		}
	}
}

func ScanKey2() {
	err := term.Init()
	if err != nil {
		panic(err)
	}

	defer term.Close()

	fmt.Println("Enter any key to see their ASCII code or press ESC button to quit")

	for {
		switch ev := term.PollEvent(); ev.Type {
		case term.EventKey:
			switch ev.Key {
			case term.KeyEsc:
				term.Sync()
				fmt.Println("ESC pressed")
				return
			case term.KeyF1:
				term.Sync()
				fmt.Println("F1 pressed")
				panic("panic")
			case term.KeyF2:
				term.Sync()
				fmt.Println("F2 pressed")
			case term.KeyF3:
				term.Sync()
				fmt.Println("F3 pressed")
			case term.KeyF4:
				term.Sync()
				fmt.Println("F4 pressed")
			case term.KeyF5:
				term.Sync()
				fmt.Println("F5 pressed")
			case term.KeyF6:
				term.Sync()
				fmt.Println("F6 pressed")
			case term.KeyF7:
				term.Sync()
				fmt.Println("F7 pressed")
			case term.KeyF8:
				term.Sync()
				fmt.Println("F8 pressed")
			case term.KeyF9:
				term.Sync()
				fmt.Println("F9 pressed")
			case term.KeyF10:
				term.Sync()
				fmt.Println("F10 pressed")
			case term.KeyF11:
				term.Sync()
				fmt.Println("F11 pressed")
			case term.KeyF12:
				term.Sync()
				fmt.Println("F12 pressed")
			case term.KeyInsert:
				term.Sync()
				fmt.Println("Insert pressed")
			case term.KeyDelete:
				term.Sync()
				fmt.Println("Delete pressed")
			case term.KeyHome:
				term.Sync()
				fmt.Println("Home pressed")
			case term.KeyEnd:
				term.Sync()
				fmt.Println("End pressed")
			case term.KeyPgup:
				term.Sync()
				fmt.Println("Page Up pressed")
			case term.KeyPgdn:
				term.Sync()
				fmt.Println("Page Down pressed")
			case term.KeyArrowUp:
				term.Sync()
				fmt.Println("Arrow Up pressed")
			case term.KeyArrowDown:
				term.Sync()
				fmt.Println("Arrow Down pressed")
			case term.KeyArrowLeft:
				term.Sync()
				fmt.Println("Arrow Left pressed")
			case term.KeyArrowRight:
				term.Sync()
				fmt.Println("Arrow Right pressed")
			case term.KeySpace:
				term.Sync()
				fmt.Println("Space pressed")
			case term.KeyBackspace:
				term.Sync()
				fmt.Println("Backspace pressed")
			case term.KeyEnter:
				term.Sync()
				fmt.Println("Enter pressed")
			case term.KeyTab:
				term.Sync()
				fmt.Println("Tab pressed")

			default:
				// we only want to read a single character or one key pressed event
				term.Sync()
				fmt.Println("ASCII : ", ev.Ch)
			}
		case term.EventError:
			panic(ev.Err)
		}
	}
}

func ScanKey3() {
	err := term.Init()
	if err != nil {
		panic(err)
	}

	defer term.Close()

	// fmt.Println("Enter any key to see their ASCII code or press ESC button to quit")

	for {
		switch ev := term.PollEvent(); ev.Type {
		case term.EventKey:
			switch ev.Key {
			case term.KeyEsc:
				fmt.Println("ESC pressed")
				return
			case term.KeyF1:
				fmt.Println("F1 pressed")
				panic("panic")
			case term.KeyF2:
				fmt.Println("F2 pressed")
			case term.KeyF3:
				fmt.Println("F3 pressed")
			case term.KeyF4:
				fmt.Println("F4 pressed")
			case term.KeyF5:
				fmt.Println("F5 pressed")
			case term.KeyF6:
				fmt.Println("F6 pressed")
			case term.KeyF7:
				fmt.Println("F7 pressed")
			case term.KeyF8:
				fmt.Println("F8 pressed")
			case term.KeyF9:
				fmt.Println("F9 pressed")
			case term.KeyF10:
				fmt.Println("F10 pressed")
			case term.KeyF11:
				fmt.Println("F11 pressed")
			case term.KeyF12:
				fmt.Println("F12 pressed")
			case term.KeyInsert:
				fmt.Println("Insert pressed")
			case term.KeyDelete:
				fmt.Println("Delete pressed")
			case term.KeyHome:
				fmt.Println("Home pressed")
			case term.KeyEnd:
				fmt.Println("End pressed")
			case term.KeyPgup:
				fmt.Println("Page Up pressed")
			case term.KeyPgdn:
				fmt.Println("Page Down pressed")
			case term.KeyArrowUp:
				fmt.Println("Arrow Up pressed")
			case term.KeyArrowDown:
				fmt.Println("Arrow Down pressed")
			case term.KeyArrowLeft:
				fmt.Println("Arrow Left pressed")
			case term.KeyArrowRight:
				fmt.Println("Arrow Right pressed")
			case term.KeySpace:
				fmt.Println("Space pressed")
			case term.KeyBackspace:
				fmt.Println("Backspace pressed")
			case term.KeyEnter:
				fmt.Println("Enter pressed")
			case term.KeyTab:
				fmt.Println("Tab pressed")

			default:
				// we only want to read a single character or one key pressed event
				fmt.Println("ASCII : ", ev.Ch)
			}
		case term.EventError:
			panic(ev.Err)
		}
	}
}

func main() {}
