package gl

/*
#cgo LDFLAGS: -lglfw -lvulkan -ldl -lpthread -lX11 -lXxf86vm -lXrandr -lXi
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

VkResult create_instance(VkInstance instance)
{
	struct VkApplicationInfo appInfo;
    	appInfo.sType = VK_STRUCTURE_TYPE_APPLICATION_INFO;
    	appInfo.pApplicationName = "Hello Triangle";
    	appInfo.applicationVersion = VK_MAKE_VERSION(1, 0, 0);
    	appInfo.pEngineName = "No Engine";
    	appInfo.engineVersion = VK_MAKE_VERSION(1, 0, 0);
    	appInfo.apiVersion = VK_API_VERSION_1_0;

	struct VkInstanceCreateInfo createInfo;
	createInfo.sType = VK_STRUCTURE_TYPE_INSTANCE_CREATE_INFO;
	createInfo.pApplicationInfo = &appInfo;

	uint32_t glfwExtensionCount = 0;
	const char** glfwExtensions;

	glfwExtensions = glfwGetRequiredInstanceExtensions(&glfwExtensionCount);

	createInfo.enabledExtensionCount = glfwExtensionCount;
	createInfo.ppEnabledExtensionNames = glfwExtensions;

	createInfo.enabledLayerCount = 0;

	printf("WTF\n");
	VkResult result = vkCreateInstance(&createInfo, NULL, &instance);
	printf("res: %d", result);
	if (result != VK_SUCCESS)
	{
		printf("HEELELELLELELELELEL\n");
		// Print out all extensions, so we know what extensions we have.
		uint32_t extensionCount = 0;
		vkEnumerateInstanceExtensionProperties(NULL, &extensionCount, NULL);

		printf("size: %d\n", extensionCount);

		VkExtensionProperties extensions[extensionCount];
		VkResult exResult = vkEnumerateInstanceExtensionProperties(NULL, &extensionCount, extensions);
		if (exResult != VK_SUCCESS)
		{
			printf("failed: %d\n", exResult);
		}


		printf("size: %d\n", extensionCount);

		for (int i = 0; i < extensionCount; i++) {
    			printf("%s\n", extensions[i].extensionName);
		}
	}

	//vkDestroyInstance(instance, NULL);

	return result;
}


*/
import "C"
import "fmt"

func Init() {
	initWindow()
}

var instance C.VkInstance

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
	C.vkDestroyInstance(instance, nil)

	C.glfwDestroyWindow(window)

	C.glfwTerminate()
}

// createInstance initialises the Vulkan library and feeds information about the
// application to the Vulkan library. This information will be provided to the
// driver and may be helpful in optimizing the application.
func createInstance() {
	/*appInfo := C.VkApplicationInfo{}
	appInfo.sType = C.VK_STRUCTURE_TYPE_APPLICATION_INFO
	appInfo.pApplicationName = C.CString("Hello Triangle")
	appInfo.applicationVersion = C.vk_version(1, 0, 0)
	appInfo.pEngineName = C.CString("No Engine")
	appInfo.engineVersion = C.vk_version(1, 0, 0)
	appInfo.apiVersion = C.VK_API_VERSION_1_0

	createInfo := C.VkInstanceCreateInfo{}
	createInfo.sType = C.VK_STRUCTURE_TYPE_INSTANCE_CREATE_INFO
	createInfo.pApplicationInfo = &appInfo

	var glfwExtensionCount C.uint
	var glfwExtensions **C.char

	glfwExtensions = C.glfwGetRequiredInstanceExtensions(&glfwExtensionCount)

	createInfo.enabledExtensionCount = glfwExtensionCount
	createInfo.ppEnabledExtensionNames = glfwExtensions

	createInfo.enabledLayerCount = 0

	result := C.vkCreateInstance(&createInfo, nil, &instance)
	if result != C.VK_SUCCESS {
		panic(fmt.Sprintf("failed to create instance, result: %v", result))
	}*/

	result := C.create_instance(instance)
	if result != C.VK_SUCCESS {
		panic(fmt.Sprintf("failed to create instance, result: %v", result))
	}

}
