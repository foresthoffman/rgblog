/**
 * rgblog.go
 *
 * Copyright (c) 2017 Forest Hoffman. All Rights Reserved.
 * License: MIT License (see the included LICENSE file)
 */

package rgblog

import (
	"fmt"
)

// Formats the output according to the format string and uses the provided interface data to fill
// it in. The output is wrapped with ANSI color escape codes, which can be used to make standard
// output more readable, in supporting shells.
//
// Output color: black
func BkPrintf(f string, a ...interface{}) {
	fmt.Printf("\033[0;30m"+f+"\033[0m", a...)
}

// Formats the output according to the format string and uses the provided interface data to fill
// it in. The output is wrapped with ANSI color escape codes, which can be used to make standard
// output more readable, in supporting shells.
//
// Output color: red
func RPrintf(f string, a ...interface{}) {
	fmt.Printf("\033[0;31m"+f+"\033[0m", a...)
}

// Formats the output according to the format string and uses the provided interface data to fill
// it in. The output is wrapped with ANSI color escape codes, which can be used to make standard
// output more readable, in supporting shells.
//
// Output color: green
func GPrintf(f string, a ...interface{}) {
	fmt.Printf("\033[0;32m"+f+"\033[0m", a...)
}

// Formats the output according to the format string and uses the provided interface data to fill
// it in. The output is wrapped with ANSI color escape codes, which can be used to make standard
// output more readable, in supporting shells.
//
// Output color: yellow
func YPrintf(f string, a ...interface{}) {
	fmt.Printf("\033[0;33m"+f+"\033[0m", a...)
}

// Formats the output according to the format string and uses the provided interface data to fill
// it in. The output is wrapped with ANSI color escape codes, which can be used to make standard
// output more readable, in supporting shells.
//
// Output color: blue
func BlPrintf(f string, a ...interface{}) {
	fmt.Printf("\033[0;34m"+f+"\033[0m", a...)
}

// Formats the output according to the format string and uses the provided interface data to fill
// it in. The output is wrapped with ANSI color escape codes, which can be used to make standard
// output more readable, in supporting shells.
//
// Output color: magenta
func MPrintf(f string, a ...interface{}) {
	fmt.Printf("\033[0;35m"+f+"\033[0m", a...)
}

// Formats the output according to the format string and uses the provided interface data to fill
// it in. The output is wrapped with ANSI color escape codes, which can be used to make standard
// output more readable, in supporting shells.
//
// Output color: cyan
func CPrintf(f string, a ...interface{}) {
	fmt.Printf("\033[0;36m"+f+"\033[0m", a...)
}

// Formats the output according to the format string and uses the provided interface data to fill
// it in. The output is wrapped with ANSI color escape codes, which can be used to make standard
// output more readable, in supporting shells.
//
// Output color: white
func WPrintf(f string, a ...interface{}) {
	fmt.Printf("\033[0;37m"+f+"\033[0m", a...)
}
