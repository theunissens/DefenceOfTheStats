package gl

/*
#include <GL/glut.h>
#include <stdio.h>

extern void goStart(unsigned char, int, int);
*/
import "C"
import "fmt"

//export goStart
func goStart(key C.uchar, x C.int, y C.int) {
	fmt.Println("You know it ;)!!!")
	switch key {
	case 27:
		//os.Exit(0)
		C.exit(0)
	}
}
