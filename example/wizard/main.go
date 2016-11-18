package main

import (
	"image"

	"github.com/marcusolsson/tui-go"
)

var logo = `     _____ __ ____  ___   ______________  
    / ___// //_/\ \/ / | / / ____/_  __/  
    \__ \/ ,<    \  /  |/ / __/   / /     
   ___/ / /| |   / / /|  / /___  / /      
  /____/_/ |_|  /_/_/ |_/_____/ /_/     `

func main() {

	userpassBox := tui.NewGrid(0, 0)
	userpassBox.SetBorder(true)
	userpassBox.SetSizePolicy(tui.Expanding, tui.Minimum)
	userpassBox.AppendRow(tui.NewLabel("User"), tui.NewLabel("Password"))
	userpassBox.AppendRow(tui.NewEntry(), tui.NewEntry())

	loginBtn := tui.NewButton("[Login]")
	loginBtn.SetFocused(true)
	registerBtn := tui.NewButton("[Register]")

	btnGroup := tui.NewHBox(
		tui.NewSpacer(),
		tui.NewPadder(loginBtn, image.Point{1, 0}),
		tui.NewPadder(registerBtn, image.Point{1, 0}),
	)
	btnGroup.SetSizePolicy(tui.Expanding, tui.Minimum)

	loginBox := tui.NewVBox(
		tui.NewPadder(tui.NewLabel(logo), image.Point{10, 1}),
		tui.NewPadder(tui.NewLabel("Welcome to Skynet! Login or register."), image.Point{12, 0}),
		userpassBox,
		btnGroup,
	)
	loginBox.SetBorder(true)

	wrapper := tui.NewVBox(
		tui.NewSpacer(),
		loginBox,
		tui.NewSpacer(),
	)
	wrapper.SetSizePolicy(tui.Minimum, tui.Expanding)

	root := tui.NewHBox(
		tui.NewSpacer(),
		wrapper,
		tui.NewSpacer(),
	)
	root.SetSizePolicy(tui.Expanding, tui.Expanding)

	// Start the application.
	ui := tui.New(root)
	if err := ui.Run(); err != nil {
		panic(err)
	}
}
