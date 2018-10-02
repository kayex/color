# color 🎨
A CSS color format conversion tool.

![Color screenshot](/screen.png?raw=true&cache=2 "Color")

# Installation
```bash
go get github.com/kayex/color/cmd/color
```

# Usage
You can either pass the color string as the sole argument to `color`, or run it without any arguments to enter interactive mode.

### Passing color string as argument
```
$ color "rgb(232, 112, 96)"

 Input (RGB)	rgb(232, 112, 96)

 [1] sRGB	15233120
 [2] Hex	#e87060
 [3] RGB	232 112 96
 [4] RGB	rgb(232, 112, 96)
 
Copy> 
```

### Entering color string in interactive mode
```
$ color
> #abc

 Input (hex)	#aabbcc

 [1] sRGB	11189196
 [2] Hex	#aabbcc
 [3] RGB	170 187 204
 [4] RGB	rgb(170, 187, 204)
 
Copy>
```
