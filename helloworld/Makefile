install:
	@echo "Installing Vulkan"

	# Install Vulkan tools
	sudo apt install vulkan-tools libvulkan-dev vulkan-validationlayers-dev spirv-tools
	# Install GLFW, used to created platform agnostic window.
	sudo apt install libglfw3-dev
	# Install GLM, a library containing linear algebra operations.
	sudo apt install libglm-dev

	# https://packages.debian.org/sid/x11/xorg-dev
	# This is needed for '-lXxf86vm' and '-lXi'.
	sudo apt-get install xorg-dev

	@echo "Vulkan installation complete"
build:
	@echo "Creating empty text files..."

	go tool cgo -exportheader gl/_cgo_export.h gl/move.go
	go build 
run:
	./helloworld
