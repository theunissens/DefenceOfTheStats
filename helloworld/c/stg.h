#ifndef STG_INCLUDE_GRAPHICS
#define STG_INCLUDE_GRAPHICS

GLFWwindow* init();

void render(GLFWwindow* window);

void teardown(GLFWwindow* window, VkInstance instance);

#endif //STG_INCLUDE_GRAPHICS

