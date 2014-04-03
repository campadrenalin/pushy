package main

import (
    "github.com/nsf/gothic"
)

var init_code = `
package require Tk

wm title . "Pushy"
grid [ttk::frame .grid -padding "3 3 12 12"] -column 0 -row 0 -sticky nwes
grid [ttk::frame .buttons -padding "3 3 12 12"] -column 0 -row 1 -sticky nes

grid [ttk::label .grid.commitlabel -text "Description of the changes:"] -column 1 -row 1 -sticky e
grid [text .grid.commitentry -width 40 -height 7] -column 2 -row 1 -sticky w

grid [ttk::label .grid.branchlabel -text "Branch to push to:"] -column 1 -row 2 -sticky e
grid [ttk::combobox .grid.branch -values [list design]] -column 2 -row 2 -sticky w
.grid.branch set design

set projectmenu [tk_optionMenu .grid.project project bc abcc nudge]
grid [ttk::label .grid.projectlabel -text "Project:"] -column 1 -row 3 -sticky e
grid .grid.project -column 2 -row 3 -sticky w

grid [ttk::button .buttons.push -text "Push these changes to server" -command { git::Push [.grid.commitentry get 1.0 end] [.grid.branch get] } ] -column 1 -row 1
`

func main() {
    interpreter := gothic.NewInterpreter(init_code)
    interpreter.RegisterCommands("git", &Git{})
    <-interpreter.Done
}
