# XML Validator

This program provides a tool to validate XML files against predefined XSD schemas. Currently, it supports validation for two types of XML: `catalog` and `overdue`.

## What It Does

The XML Validator checks the structure and content of an XML file against a specified XSD schema to ensure its validity. If the XML file does not adhere to the expected structure outlined in the XSD schema, the program will output an error message detailing the reason for the mismatch.

## Dependencies

- **Go:** The program is written in Go and requires the Go toolchain to be installed.
- **libxml2:** This is a C library for parsing XML documents. The program uses Go bindings for `libxml2` to perform XML validation against XSD schemas.

## Installation & Setup

1. Ensure you have Go installed on your system. If not, download and install it from [the official Go website](https://golang.org/).
2. Install the `libxml2` library. Depending on your operating system:
   - **Debian/Ubuntu**: 
     ```bash
     sudo apt-get install libxml2-dev
     ```
   - **macOS**:
     ```bash
     brew install libxml2
     ```

3. Clone the repository to your local machine.
4. Navigate to the project root and run the `Makefile` to check and install any missing dependencies:
   ```bash
   make
   ```

## How to Use

Navigate to the `tools/validator` directory and run the program with the following flags:

- **-type**: The type of XML you want to validate. This can be either `catalog` or `overdue`.
- **-path**: The path to the XML file you wish to validate.

Example usage:

```
go run main.go -type=catalog -path=/path/to/catalog.xml
```

OR 

```
go run main.go -type=overdue -path=/path/to/overdue.xml
```

Replace `main.go` with the name of the Go file and `/path/to/catalog.xml` with the path to your XML file.

## Makefile Options

The provided `Makefile` helps automate dependency checks and installations. Here are the available options:

- `make`: This default command checks for the necessary dependencies and installs them if they're missing.
- `make check-deps`: Checks if all required dependencies are installed.
- `make install-deps-debian`: Installs the required dependencies on Debian/Ubuntu systems.
- `make install-deps-macos`: Installs the required dependencies on macOS systems.
- `make run`: Runs the Go program. Ensure you've provided the correct paths and flags.

---

Save the above content in a `README.md` file in your project root. This will guide users on how to use the XML Validator tool, its dependencies, and how to leverage the provided Makefile.