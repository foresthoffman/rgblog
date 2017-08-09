## RgbLog

The rgblog package provides a set of functions that wrap the standard `fmt.Printf` function. Each function provided applies a different color to the formatted output, using ANSI escape codes. This package assumes that your standard output (i.e. your terminal) has support for such escape codes.

[ANSI Codes](https://en.wikipedia.org/wiki/ANSI_escape_code#Colors) from Wikipedia:

Color | Foreground Code | Background Code
--- | --- | ---
Black | 30 | 40
Red | 31 | 41
Green | 32 | 42
Yellow | 33 | 43
Blue | 34 | 44
Magenta | 35 | 45
Cyan | 36 | 46
White | 37 | 47

### Installation

Run `go get github.com/foresthoffman/rgblog`

### Importing

Import the package by including `github.com/foresthoffman/rgblog` in your import block.

e.g.

```Go
package main

import(
	...
	"github.com/foresthoffman/rgblog"
)
```

If you want to reference the package via an alias, so that you don't have to type `rgblog.Xxx()` every time, try this:

```Go
import(
	...
	rgb "github.com/foresthoffman/rgblog"
)
```

### What will it look like for me?

If you want to see what the output will look like without having to write a quick Go program to use all the functions, run the code below in your shell.

```bash
echo -e "
\033[0;37mNormal Colors\033[0m
Black\t\033[0;30mThe quick brown fox jumped over the lazy dog\033[0m
Red\t\033[0;31mThe quick brown fox jumped over the lazy dog\033[0m
Green\t\033[0;32mThe quick brown fox jumped over the lazy dog\033[0m
Yellow\t\033[0;33mThe quick brown fox jumped over the lazy dog\033[0m
Blue\t\033[0;34mThe quick brown fox jumped over the lazy dog\033[0m
Magenta\t\033[0;35mThe quick brown fox jumped over the lazy dog\033[0m
Cyan\t\033[0;36mThe quick brown fox jumped over the lazy dog\033[0m
White\t\033[0;37mThe quick brown fox jumped over the lazy dog\033[0m

\033[1;37mBold/Bright Colors\033[0m
Black\t\033[1;30mThe quick brown fox jumped over the lazy dog\033[0m
Red\t\033[1;31mThe quick brown fox jumped over the lazy dog\033[0m
Green\t\033[1;32mThe quick brown fox jumped over the lazy dog\033[0m
Yellow\t\033[1;33mThe quick brown fox jumped over the lazy dog\033[0m
Blue\t\033[1;34mThe quick brown fox jumped over the lazy dog\033[0m
Magenta\t\033[1;35mThe quick brown fox jumped over the lazy dog\033[0m
Cyan\t\033[1;36mThe quick brown fox jumped over the lazy dog\033[0m
White\t\033[1;37mThe quick brown fox jumped over the lazy dog\033[0m
"
```

_That's all, enjoy!_
