# Makefile

# Set the binary name and installation directory
BINARY_NAME=kbcmd
INSTALL_DIR=/usr/local/bin
BASH_COMPLETION_DIR=/etc/bash_completion.d

# Set the repository location and the desired directory to clone to
REPO=github.com/killbill/kbcli
CLONE_DIR=./killbill
KBCMD_DIR=$(CLONE_DIR)/kbcmd

# Default target to fetch and build
all: fetch build install bash-completion config show-config

# Clone the kbcli repository
fetch:
	@echo "Cloning kbcli repository..."
	@if [ ! -d "$(CLONE_DIR)" ]; then git clone https://$(REPO).git $(CLONE_DIR); fi

# Build the kbcli binary
build:
	@echo "Building..."
	@cd $(KBCMD_DIR) && go build

# Install the kbcmd binary
install: build
	@echo "Installing to $(INSTALL_DIR)..."
	@cp $(KBCMD_DIR)/kbcmd $(INSTALL_DIR)/$(BINARY_NAME)

# Setup bash completion
bash-completion:
ifeq ($(shell uname), Linux)
	@echo "Setting up bash completion for Linux..."
	@cp $(KBCMD_DIR)/bash-autocomplete $(BASH_COMPLETION_DIR)/kbcmd
else ifeq ($(shell uname), Darwin)
	@echo "Setting up bash completion for macOS..."
	@echo "source $(KBCMD_DIR)/bash-autocomplete" >> $(HOME)/.bash_profile
	@echo "Please restart your terminal or run 'source $(HOME)/.bash_profile'"
endif

# Environment variable configuration
config:
	@echo "#!/bin/bash" > set_env.sh; \
	read -p "Enter value for KB_API_KEY (default: bob): " key; \
	echo "export KB_API_KEY=$${key:-bob}" >> set_env.sh; \
	read -p "Enter value for KB_API_SECRET (default: lazar): " secret; \
	echo "export KB_API_SECRET=$${secret:-lazar}" >> set_env.sh; \
	read -p "Enter value for KB_API_CREATED_BY (default: `whoami`): " created_by; \
	echo "export KB_API_CREATED_BY=$${created_by:-`whoami`}" >> set_env.sh; \
	read -p "Enter value for KB_HOST (default: 127.0.0.1:8080): " host; \
	echo "export KB_HOST=$${host:-127.0.0.1:8080}" >> set_env.sh; \
	chmod +x set_env.sh; \
	printf "\033[1;31mConfiguration script created as 'set_env.sh'. To apply the environment variables, run 'source set_env.sh'.\033[0m\n"



# Show current configuration
show-config:
	@echo "Displaying the current configuration for the kbcli environment variables:"
	@echo "---------------------------------------------------------"
	@echo "API Key: $$KB_API_KEY"
	@echo "API Secret: $$KB_API_SECRET"
	@echo "Created By: $$KB_API_CREATED_BY"
	@echo "Host Address: $$KB_HOST"
	@echo "---------------------------------------------------------"
	@echo "Note: If any variable is empty or not updated, it hasn't been set in the current session."
	@echo "To set them, you can use the 'make config' command if you haven't already and then 'source set_env.sh'."

# Clean up the binary if it exists and the cloned repo
clean:
	@echo "Cleaning..."
	@rm -f $(KBCMD_DIR)/kbcmd
	@rm -rf $(CLONE_DIR)
