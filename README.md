# Go WebAssembly Virtual Machine

This project demonstrates a simple Virtual Machine implemented in Go and compiled to WebAssembly (WASM). The VM has `registers`, and you can `set` and `get` values from those registers using JavaScript.

## Prerequisites

To run this project, you need the following:

- Go compiler
- `wasm_exec.js` - This is a JavaScript glue code provided by the Go distribution.

## Project Structure

The project consists of the following files:

- `main.go`: Contains the Go code for the virtual machine implementation.
- `js-wasm-vm-interop.html`: Contains the HTML and JavaScript code for the user interface.
- `wasm_exec.js`: JavaScript glue code provided by the Go distribution.

## Virtual Machine Implementation (Go)

The virtual machine implementation is contained in the `main.go` file. Here's an overview of the components:

- `VirtualMachine`: Represents a simple virtual machine with registers. It has the following methods:
  - `NewVirtualMachine()`: Creates a new instance of `VirtualMachine`.
  - `SetRegister(name string, value int)`: Sets the value of a register.
  - `GetRegister(name string) int`: Retrieves the value of a register.
- `setupVM(vm *VirtualMachine)`: Exposes the virtual machine methods to JavaScript.

## User Interface (HTML/JS)

The user interface is implemented in the `js-wasm-interop.html` file. Here's an overview of the components:

- Input Field: Allows the user to enter a value.
- "Set in Register" Button: Sets the entered value in the virtual machine register.
- "Get from Register" Button: Retrieves the value from the virtual machine register and displays it.

## Build and Run

To build and run the project, follow these steps:

1. Clone this repository to your local machine.

2. Build the Go code to WebAssembly using the following command:

   ```bash
   GOOS=js GOARCH=wasm go build -o go-wasm-vm.wasm
   ```

   (alternatively you can run the `build.sh` script).

   This command builds the Go code and generates the WebAssembly binary file `go-wasm-vm.wasm`.

3. Import the latest `wasm_exec.js` in the project:

    ```bash
    cp $(go env GOROOT)/misc/wasm/wasm_exec.js .
    ```

4. Start a local web server in the project directory. For example, you can use the Python built-in HTTP server:

   ```bash
   python -m http.server
   ```

5. Open a web browser and visit `http://localhost:8000`. You should see the virtual machine interface.

6. Enter a value in the input field and click `Set in Register` to set the value in the virtual machine register.

7. Click `Get from Register` to retrieve the value from the virtual machine register. The value will be displayed in the browser console and shown in an alert.

## License

This project is licensed under the [MIT License](LICENSE).

## Notes

Please ensure that you have the `wasm_exec.js` file in the same directory as the HTML file and the `go-wasm-vm.wasm` file. Follow the instructions in the README.md file to build and run the project.

If you have any further questions or need additional assistance, please let me know!
