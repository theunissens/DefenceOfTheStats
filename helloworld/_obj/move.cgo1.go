// Code generated by cmd/cgo; DO NOT EDIT.

//line /home/taffer/dev/projects/helloworld/gl/move.go:1:1
package gl

/*
#include <GL/glut.h>
#include <stdio.h>

extern void goStart(unsigned char, int, int);
*/
import _ "unsafe"
import "fmt"

//export goStart
func goStart(key  /*line :13:18*/_Ctype_uchar /*line :13:25*/, x  /*line :13:29*/_Ctype_int /*line :13:34*/, y  /*line :13:38*/_Ctype_int /*line :13:43*/) {
	fmt.Println("You know it ;)!!!")
	switch key {
	case 27:
		//os.Exit(0)
		( /*line :18:3*/_Cfunc_exit /*line :18:8*/)(0)
	}
}
