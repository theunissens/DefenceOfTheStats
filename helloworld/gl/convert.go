package gl

/*
#cgo LDFLAGS: -lglfw -lvulkan -ldl -lpthread -lX11 -lXxf86vm -lXrandr -lXi
#include <stdio.h>
#define GLFW_INCLUDE_VULKAN
#include <GLFW/glfw3.h>
#include "stg.h"

// vk_version wraps VK_MAKE_VERSION. VK_MAKE_VERSION is a Macro. cgo does not play well
// with macros, and we therefore wrap it.
uint32_t vk_version(const int major, const int minor, const int patch)
{
	return VK_MAKE_VERSION(major, minor, patch);
}


*/
import "C"

func Init() {
	C.glfwInit()

	var extensionCount C.uint
	C.vkEnumerateInstanceExtensionProperties(nil, &extensionCount, nil)

	println("%v extensions supported", extensionCount)
}

func CreateWindow(name string, width, height int) *C.GLFWwindow {
	C.glfwWindowHint(C.GLFW_CLIENT_API, C.GLFW_NO_API)
	return C.glfwCreateWindow(C.int(width), C.int(height), C.CString(name), nil, nil)
}

func CreateInstance(appName, engineName string) C.VkInstance {
	appInfo := C.VkApplicationInfo{
		sType:              C.VK_STRUCTURE_TYPE_APPLICATION_INFO,
		pApplicationName:   C.CString(appName),
		applicationVersion: C.vk_version(1, 0, 0),
		pEngineName:        C.CString(engineName),
		engineVersion:      C.vk_version(1, 0, 0),
		apiVersion:         C.VK_API_VERSION_1_0,
	}

	instance := C.create_instance(appInfo, C.VK_STRUCTURE_TYPE_INSTANCE_CREATE_INFO)
	println("instance created")
	return instance
}

func Wait(window *C.GLFWwindow) {
	for C.glfwWindowShouldClose(window) == 0 {
		C.glfwPollEvents()
	}
}

func Teardown(window *C.GLFWwindow, instance C.VkInstance) {
	C.vkDestroyInstance(instance, nil)

	C.glfwDestroyWindow(window)
	C.glfwTerminate()
}
