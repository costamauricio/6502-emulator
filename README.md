# 6502 Emulator

Simple emulator of 6502 cpu in Golang.

![image](https://user-images.githubusercontent.com/3957076/189212091-485c172a-b81e-440e-9d9c-fb9549cdd214.png)

## Contents

Packages:

- cpu6502 -> 6502 CPU emulator
- bus -> Simple BUS to attach to the emulator to provide RAM addresses
- debugger -> SDL2 implementation to visualize the current CPU status

## Dependencies

- [go-sdl2](https://github.com/veandco/go-sdl2#requirements) to run the debbugger

## Running the debbuger

```bash
$ go mod tidy
$ make run
```

## References

Based on the oneloanecoder series of NES emulator https://www.youtube.com/watch?v=nViZg02IMQo&list=PLrOv9FMX8xJHqMvSGB_9G9nZZ_4IgteYf

- Datasheet -> [http://archive.6502.org/datasheets/rockwell_r650x_r651x.pdf](http://archive.6502.org/datasheets/rockwell_r650x_r651x.pdf)
- More detailed datasheet with instruction descriptions -> [https://www.princeton.edu/~mae412/HANDOUTS/Datasheets/6502.pdf](https://www.princeton.edu/~mae412/HANDOUTS/Datasheets/6502.pdf)
- Instruction set and Address modes details explained -> [https://www.masswerk.at/6502/6502_instruction_set.html](https://www.masswerk.at/6502/6502_instruction_set.html)
- Vflag (overflow) explanation article -> [http://www.6502.org/tutorials/vflag.html](http://www.6502.org/tutorials/vflag.html)
