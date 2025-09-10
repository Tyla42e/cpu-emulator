# CPU Emulator

This project is a CPU emulator that supports the emulation of three different CPU architectures: 6502, 6510, and Z80. Each CPU emulator is implemented in its own package, allowing for modular development and easy maintenance.

## Project Structure

```
cpu-emulator
├── cmd
│   └── main.go          # Entry point of the application
├── pkg
│   ├── cpu6502
│   │   └── cpu6502.go   # Implementation of the 6502 CPU emulator
│   ├── cpu6510
│   │   └── cpu6510.go   # Implementation of the 6510 CPU emulator
│   └── cpuZ80
│       └── cpuZ80.go    # Implementation of the Z80 CPU emulator
├── go.mod                # Module definition for the Go project
└── README.md             # Documentation for the project
```

## Getting Started

To build and run the CPU emulator, follow these steps:

1. **Clone the repository:**
   ```
   git clone <repository-url>
   cd cpu-emulator
   ```

2. **Install dependencies:**
   ```
   go mod tidy
   ```

3. **Run the emulator:**
   ```
   go run cmd/main.go
   ```

## CPU Architectures

### 6502
The 6502 CPU emulator is implemented in the `pkg/cpu6502` package. It provides methods to reset the CPU, execute a single instruction (step), and load a program into memory.

### 6510
The 6510 CPU emulator is implemented in the `pkg/cpu6510` package. Similar to the 6502, it includes methods for resetting the CPU, stepping through instructions, and loading programs.

### Z80
The Z80 CPU emulator is implemented in the `pkg/cpuZ80` package. It features methods for resetting the CPU, executing instructions, and loading programs specific to the Z80 architecture.

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue for any enhancements or bug fixes.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.