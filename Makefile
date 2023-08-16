.PHONY: run-validator check-deps install-deps-debian install-deps-macos update-xsd update-catalog update-overdue

# Default target
all: run-validator

# Include the validator's Makefile
include tools/validator/Makefile

# Run the validator program
run-validator:
	@read -p "Enter the type of XML (catalog/overdue): " type; \
	read -p "Enter the path to the XML file: " path; \
	if echo "$$path" | grep -q "tools/validator"; then \
		path=`realpath $$path`; \
	fi; \
	cd tools/validator && go run main.go -type=$$type -path=$$path

# Check dependencies
check-deps: check-validator-deps

# Install dependencies on Debian/Ubuntu
install-deps-debian: install-validator-deps-debian

# Install dependencies on macOS
install-deps-macos: install-validator-deps-macos

# Update XSD files
update-xsd: update-catalog update-overdue

# Download and update catalog.xsd
update-catalog:
	@curl -o schemas/catalog.xsd https://docs.killbill.io/latest/catalog.xsd

# Download and update overdue.xsd
update-overdue:
	@curl -o schemas/overdue.xsd https://docs.killbill.io/latest/overdue.xsd