# CPU Emulator

A modular emulator framework for classic CPUs, supporting the strategy pattern and extensible instruction sets.

## Project Structure

```
cpu-emulator/
├── pkg/
│   ├── cpu/           # Common CPU interface and instruction definitions
│   │   └── instruction.go
│   ├── cpu6502/       # MOS 6502 CPU implementation and instructions
│   │   ├── cpu6502.go
│   │   ├── instructions.go
│   │   └── cpu6502_test.go
│   ├── cpu6510/       # MOS 6510 CPU implementation and instructions
│   │   ├── cpu6510.go
│   │   └── instructions.go
│   ├── cpu65c02/      # WDC 65C02 CPU implementation and instructions
│   │   ├── cpu65c02.go
│   │   └── instructions.go
│   ├── cpuz80/        # Zilog Z80 CPU implementation and instructions
│   │   ├── cpuz80.go
│   │   └── instructions.go
│   ├── cpu68000/      # Motorola 68000 CPU implementation and instructions
│   │   ├── cpu68000.go
│   │   └── instructions.go
│   └── memory/        # Memory abstraction
│       └── memory.go
└── main.go            # Example entry point
```

## Features

- **Strategy Pattern:** Swap CPU implementations at runtime via a common interface.
- **Extensible Instruction Sets:** Each CPU has its own instruction array, easily extended or overridden.
- **Modular Design:** Add new CPUs by creating a new package and implementing the interface.
- **Unit Tests:** Example tests included for instruction correctness.
- **Memory Abstraction:** Easily configure memory size and behavior.

## Adding a New CPU

1. Create a new package in `pkg/` (e.g., `cpu8080`).
2. Implement the `cpu.CPU` interface.
3. Define an instruction set array for the new CPU.
4. Add addressing mode helpers and flag helpers as needed.

## Example Usage

```go
import (
    "emulator/pkg/cpu6502"
)

func main() {
    cpu := cpu6502.New(65536)
    program := []uint8{0xA9, 0x42, 0xAA, 0x00} // LDA #$42; TAX; BRK
    cpu.LoadProgram(program, 0x8000)
    cpu.(*cpu6502.CPU6502).Memory.Write(0xFFFC, 0x00)
    cpu.(*cpu6502.CPU6502).Memory.Write(0xFFFD, 0x80)
    cpu.Reset()
    for i := 0; i < 3; i++ {
        cpu.Step()
    }
}
```

## Testing

Run all tests with:

```sh
go test ./pkg/...
```

## Contributing

- Fork and submit pull requests for new CPUs or improvements.
- Please include tests for new instructions or features.

---

**This project is a