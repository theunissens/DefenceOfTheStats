#define GLFW_INCLUDE_VULKAN
#include <GLFW/glfw3.h>
#include <GL/glut.h>

#include <stdio.h>
#include "stg.h"

VkInstance instance;

//gcc stg.c -lglfw -lvulkan -ldl -lpthread -lX11 -lXxf86vm -lXrandr -lXi

VkInstance createInstance()
{
	// remove soon.
	glfwInit();

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

	int sup;
	sup = glfwVulkanSupported();
	if (sup == GLFW_TRUE)
	{
		printf("supported");
	}
	if (sup == GLFW_FALSE)
	{
		printf("unsupported");
	}


	printf("count: %d\n", glfwExtensionCount);
	printf("extensions: %c\n", **glfwExtensions);
	
	createInfo.enabledExtensionCount = glfwExtensionCount;
	//createInfo.enabledExtensionCount = 1;
	createInfo.ppEnabledExtensionNames = glfwExtensions;
	
	createInfo.enabledLayerCount = 0;
	
	VkResult result = vkCreateInstance(&createInfo, NULL, &instance);
	//if (result != VK_SUCCESS)
	//{
	//	printf("create instance failed: %d\n", result);
	//}	

	//return instance;
}

GLFWwindow* init()
{
	glfwInit();

    	glfwWindowHint(GLFW_CLIENT_API, GLFW_NO_API);
    	GLFWwindow* window = glfwCreateWindow(800, 600, "Vulkan window", NULL, NULL);

    	uint32_t extensionCount = 0;
    	vkEnumerateInstanceExtensionProperties(NULL, &extensionCount, NULL);

    	printf("%d extensions supported\n", extensionCount);
	return window;	
}

void render(GLFWwindow* window)
{
	while(!glfwWindowShouldClose(window)) {
        	glfwPollEvents();
    	}
}

void teardown(GLFWwindow* window, VkInstance instance)
{
	vkDestroyInstance(instance, NULL);

	glfwDestroyWindow(window);
	glfwTerminate();

}

int main() {
	//GLFWwindow* window = init();
	VkInstance instance = createInstance();
//	render(window);
	//teardown(window, instance);
	
	//vkDestroyInstance(instance, NULL);

	return 0;
}

