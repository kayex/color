# color 🎨

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
...or download a pre-compiled binary from the [releases page](https://github.com/kayex/color/releases).

# Usage
The color value can either be passed as the sole argument to `color`, or entered interactively after running the program without any arguments.

```
$ color "#ffee00"
```

or

```
$ color
> #ffee00
```

# Color formats
Color tries to be as accomodating as possible when parsing input. All of the following examples are supported:
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
