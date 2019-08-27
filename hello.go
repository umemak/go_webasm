package main
import (
	"fmt"
	"syscall/js"
)
func main() {
	fmt.Println("Go launched.")
	document := js.Global().Get("document")
	h1 := document.Call("createElement", "h1")
	h1.Set("innerHTML", "Hello, WASM World")
	body := document.Get("body")
	body.Call("appendChild", h1)
}

