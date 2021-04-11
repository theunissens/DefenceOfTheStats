package gl

// We are missing '-lXxf86vm' and '-lXi' from the LDFLAGS. We don't have them installed
// and we do not know what they do...

/*
#cgo LDFLAGS: -lglfw -lvulkan -ldl -lpthread -lX11 -lXrandr
#include <GL/glut.h>
#include <stdio.h>
#include <_cgo_export.h>

#define GLFW_INCLUDE_VULKAN
#include <GLFW/glfw3.h>

// vk_version wraps VK_MAKE_VERSION. VK_MAKE_VERSION is a Macro. cgo does not play well
// with macros, and we therefore wrap it.
uint vk_version(const int major, const int minor, const int patch)
{
	return VK_MAKE_VERSION(major, minor, patch);
}

void exit_window(unsigned char key, int x, int y)
{
	//esc(key, x, y);
}

void key_press(int key, int x, int y)
{
        //keyPress(key, x, y);
}


*/
import "C"

func Init() {
	initWindow()
}

var (
	boxX, boxY C.double
)

func initWindow() {
	/*a := C.CString("a")
	var i C.int

	C.glutInit(&i, &a)

	C.glutInitWindowSize(800, 800);
	C.glutCreateWindow(C.CString("ABGR extension"))

	C.glutKeyboardFunc((*[0]byte)(C.exit_window))
	C.glutSpecialFunc((*[0]byte)(C.key_press))

	C.glutMainLoop()*/

	C.glfwInit()

	C.glfwWindowHint(C.GLFW_CLIENT_API, C.GLFW_NO_API)
	C.glfwWindowHint(C.GLFW_RESIZABLE, C.GLFW_TRUE)

	name := C.CString("Vulkan window")
	window := C.glfwCreateWindow(800, 600, name, nil, nil)

	createInstance()

	//uint32_t extensionCount = 0;
	//vkEnumerateInstanceExtensionProperties(nullptr, &extensionCount, nullptr);

	// Keep the application running until either the program is closed, or an error occurs.
	for C.glfwWindowShouldClose(window) == 0 {
		C.glfwPollEvents()
	}

	// Cleanup.
	C.glfwDestroyWindow(window)

	C.glfwTerminate()
}

// createInstance initialises the library and feeds information about the application
// to the Vulkan library. This information will be provided to the driver and may be
// helpful in optimizing the application.
func createInstance() {
	appInfo := &C.VkApplicationInfo{}
	appInfo.sType = C.VK_STRUCTURE_TYPE_APPLICATION_INFO
	appInfo.pApplicationName = C.CString("Hello Triangle")
	appInfo.applicationVersion = C.vk_version(1, 0, 0)
	appInfo.pEngineName = C.CString("No Engine")
	//appInfo.engineVersion = C.VK_MAKE_VERSION(1, 0, 0)
	appInfo.apiVersion = C.VK_API_VERSION_1_0
}
