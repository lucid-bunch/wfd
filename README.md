# wfd
![build status](https://github.com/lucid-bunch/wfd/actions/workflows/go.yml/badge.svg)

## Description

wfd (What's For Dinner) is a tool that helps with dinner planning. This tool is designed to simplify the process of deciding what to cook for dinner.

## Features

- Generates recipes: This tool generates a list of recipes that meet your criteria.
- Cooldown store: This feature allows you to set a cooldown period for certain recipes to avoid repetition.
- Block store: This feature allows you to block certain recipes that you do not want to cook.

## Prerequisites

Before you can run this application, you need to have Go installed on your machine. Here's how you can install it:

### Windows

1. Visit the [Go Downloads page](https://golang.org/dl/)
2. Download the Microsoft Windows installer
3. Open the installer and follow the prompts

### macOS

1. Visit the [Go Downloads page](https://golang.org/dl/)
2. Download the Apple macOS installer
3. Open the installer and follow the prompts

### Linux

1. Visit the [Go Downloads page](https://golang.org/dl/)
2. Download the archive suitable for your distribution
3. Extract the archive to `/usr/local` (you may need to use `sudo`)
4. Add `/usr/local/go/bin` to the PATH environment variable

You can verify the installation by opening a command prompt and running `go version`. This should display the installed version of Go.

## Installation

1. Clone the repository: `git clone https://github.com/lucid-bunch/wfd.git`
2. Navigate to the project directory: `cd wfd`
3. Install the dependencies: `go get ./...`

## Usage

Run the program with the command: `go run main.go`

You can use the `-t` flag to specify the type of recipes you want to generate. For example, `go run main.go -t vegetarisk` will generate vegetarian (vegetarisk in Swedish) recipes. If the `-t` flag is not provided, the default value is `all`, which means all types of recipes will be generated.

Use the `-h` flag to display help information about the command-line flags. For example, `go run main.go -h` will display a list of all available flags and their descriptions.

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change. Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)
