package gl

/*
#cgo LDFLAGS: -lGL -lglut -lGLU -lm
#include <GL/glut.h>
#include <stdio.h>
#include <_cgo_export.h>

void doThing(unsigned char key, int x, int y)
{
	// glEnd();
	goStart(key, x, y);
}

*/
import "C"
import (
	"unsafe"
)

func Init() {
	//C.myinit()
	initWindow()
}

var (
	boxX, boxY C.double
)

func initWindow() {
	a := C.CString("a")
	var i C.int
	//	a := *C.char(&aa)
	//ii := 0
	//	i := C.int(0)
	C.glutInit(&i, &a)
	C.glutCreateWindow(C.CString("ABGR extension"))

	keyhitfunc := keyhit
	keyhitfunc1 := &keyhitfunc
	keyhitfunc2 := unsafe.Pointer(keyhitfunc1)

	//	_ = escfunc2
	_ = keyhitfunc2

	C.glutKeyboardFunc((*[0]byte)(C.doThing)) // use esc here
	//C.glutKeyboardFunc((*[0]byte)(escfunc2)) // use esc here
	//C.glutSpecialFunc((*[0]byte)(C.keyhit))
	//C.glutSpecialFunc((*[0]byte)(keyhitfunc2))
	C.glutMainLoop()
}

func draw() {
	C.glMatrixMode(C.GL_PROJECTION)
	C.glLoadIdentity()

	C.glColor3d(0.5, 0.5, 0.5)
	C.glBegin(C.GL_POLYGON)

	C.glVertex2d(boxX, boxY)
	C.glVertex2d(boxX+0.2, boxY)
	C.glVertex2d(boxX+0.2, boxY+0.2)
	C.glVertex2d(boxX, boxY+0.2)

	C.glEnd()

	C.glFinish()
}

func keyhit(key C.int, xx C.int, yy C.int) {
	C.glClear(C.GL_COLOR_BUFFER_BIT)

	switch key {
	case C.GLUT_KEY_UP:
		boxY = boxY + 0.1
		draw()

		break

	case C.GLUT_KEY_DOWN:
		boxY = boxY - 0.1
		draw()

		break

	case C.GLUT_KEY_LEFT:
		boxX = boxX - 0.1
		draw()

		break
	case C.GLUT_KEY_RIGHT:
		boxX = boxX + 0.1
		draw()

		break
	}
}
