# color 🎨
A CSS color format conversion tool.

<p align="center">
 <img src="/screen.png?raw=true" alt="Color screenshot">
</p>

# Installation
```bash
go get github.com/kayex/color/cmd/color
```

# Usage
The color value can either be passed as the sole argument to `color`, or entered interactively after running the program without any arguments.

Color tries to be as accomodating as possible when parsing input. The following formats are known to work:

| Format                | Example                  |
| --------------------- | ------------------------ |
| RGB (plain values)    | 255 255 255              |
| RGB (8-bit channels)  | rgb(255, 255, 255)       |
| RGB (float channels)  | rgb(1.0, 1.0, 1.0)       |
| RGBa                  | rgba(255, 255, 255, 0.5) |
| Hex                   | #ffee00                  |
| Hex (short)           | #fe0                     |
| Hex (plain value)     | fe0                      |
