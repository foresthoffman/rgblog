/**
 * doc.go
 *
 * Copyright (c) 2017 Forest Hoffman. All Rights Reserved.
 * License: MIT License (see the included LICENSE file)
 */

/*
Package rgblog provides functions to color standard output using ANSI escape codes. See,
https://en.wikipedia.org/wiki/ANSI_escape_code#Colors, for more detail on ANSI codes.

Here's an example:

    import(
        rgb "github.com/foresthoffman/rgblog"
    )

    func main() {
        rgb.YPrintf("Hello, %s!\n", "World") // the format gets wrapped with "\033[0;33m"
                                             // and "\033[0m", before being sent to
                                             // standard output via fmt.Printf()
    }
*/
package rgblog
