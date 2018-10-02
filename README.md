# color
A color format conversion tool. Supports copying converted values to clipboard.

![Color screenshot](/screen.png?raw=true "Color")

# Installation
```bash
go get github.com/kayex/color/cmd/color
```

# Usage

### Color as program argument
```
$ color "rgb(232, 112, 96)"

 Input (RGB)	rgb(232, 112, 96)

 [1] sRGB	15233120
 [2] Hex	#e87060
 [3] RGB	232 112 96
 [4] RGB	rgb(232, 112, 96)
 
Copy> 
```

### Interactive mode
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
