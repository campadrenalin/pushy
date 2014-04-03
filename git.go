package main

import "fmt"

type Git struct {
}

func (g *Git) TCL_Push(msg, branch string) {
    fmt.Printf("Pushing to branch %s, with message:\n\n%s\n", branch, msg)
}
