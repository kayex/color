# color 🍭

<p align="center">
 A color conversion tool.
</p>
<p align="center">
 <img src="/screen.png?raw=true" alt="Color screenshot">
</p>

# Installation
```bash
go get github.com/kayex/color/cmd/color
```
... or download a pre-compiled binary from the [releases page](https://github.com/kayex/color/releases).

# Usage
Enter any color value as the sole argument to `color`, or enter interactive mode by running the program without any arguments.

Most color strings need to be quoted when passed directly on the command line.

```
$ color "#ffee00"
```

Choose an output format by pressing <kbd>1</kbd> - <kbd>4</kbd> followed by <kbd>Enter</kbd> and the converted color value will be copied to the system clipboard.

# Supported color formats
```
[x] RGB (8-bit channels)
[x] RGB (float channels)
[x] RGBA
[x] Hex triplet
[x] Hex triplet (short)
[ ] HSL
[ ] HSLA
```

Color tries to be as accomodating as possible when parsing input. All of the following color strings are parsed correctly:
```
255 255 255
rgb 255 255 255
rgb(255,255,255)
rgb(255, 255, 255)
rgb(1.0, 1.0, 1.0)
rgba(255, 255, 255, 0.5)
rgba 0 0 0 0.5
#f0a
#ff00aa
```
