#define GLFW_INCLUDE_VULKAN
#include <GLFW/glfw3.h>
#include <GL/glut.h>

#include <stdio.h>
#include "stg.h"

VkInstance instance;

VkInstanceCreateInfo createInfo;

//gcc stg.c -lglfw -lvulkan -ldl -lpthread -lX11 -lXxf86vm -lXrandr -lXi

VkApplicationInfo create_app_info()
{
	VkApplicationInfo app_info = {
		.sType = VK_STRUCTURE_TYPE_APPLICATION_INFO,
		.pApplicationName = "Hello Triangle",
		.applicationVersion = VK_MAKE_VERSION(1, 0, 0),
		.pEngineName = "No Engine",
		.engineVersion = VK_MAKE_VERSION(1, 0, 0),
		.apiVersion = VK_API_VERSION_1_0,
	};

	
	return app_info;
}

VkInstance create_instance(VkApplicationInfo appInfo, VkStructureType s_type)
{
	
	createInfo.sType = s_type;
	createInfo.pApplicationInfo = &appInfo;

	uint32_t glfwExtensionCount = 1;
	

	const char** glfwExtensions;
	glfwExtensions = glfwGetRequiredInstanceExtensions(&glfwExtensionCount);

	int sup;
	sup = glfwVulkanSupported();
	if (sup == GLFW_TRUE)
	{
		printf("supported\n");
       	}
	if (sup == GLFW_FALSE)
	{
		printf("unsupported"); 
	} 
	
	printf("count: %d\n", glfwExtensionCount);
	for (int i = 0; i < glfwExtensionCount; i++)
	{
		printf("extensions: %s\n", glfwExtensions[i]);
	}
		
	createInfo.enabledExtensionCount = glfwExtensionCount;
	createInfo.ppEnabledExtensionNames = glfwExtensions;
	
	createInfo.enabledLayerCount = 0;
	
	VkResult result = vkCreateInstance(&createInfo, NULL, &instance);
	//if (result != VK_SUCCESS)
	//{
	//	printf("create instance failed: %d\n", result);
	//}	

	return instance;
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

//int main() {
//	GLFWwindow* window = init();
//
//	VkApplicationInfo app_info = create_app_info();
//
//	VkInstance instance = createInstance(app_info);
//	render(window);
//	teardown(window, instance);
//	
//	return 0;
//}

