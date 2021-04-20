package main

import "github.com/helloworld/gl"

func main() {
	gl.Init()
	window := gl.CreateWindow("STG", 800, 600)
	instance := gl.CreateInstance("Hello Triangle", "No Engine")
	gl.Wait(window)
	gl.Teardown(window, instance)
}
