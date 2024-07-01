package strshredder

import (
    "strings"
)

// strShredder rebuilds strings in a slice into strings with a specific length
// (up to 45 bytes in the code below) and puts them back in the original slice.
func strShredder(s []string) []string {
    switch {
    case s == nil:
       return nil
    case len(s) == 1:
       return s
    }

    str := strings.Builder{}
    nx := make(chan int)

    go func(nx chan<- int, s []string) {
       for i := 1; i < len(s); i++ {
           nx <- len(s[i])
        }
        close(nx)
    }(nx, s)

    i, c := 0, 0
loop:
    for i < len(s) {
        if 45 < len(s[i]) {
            break
        }
        for n, x := <- nx; i < len(s) && str.Len() + n < 45; i++ {
            str.WriteString(s[i])
            n, x = <- nx
            if !x {
                for i += 1; i < len(s); {
                    if 45 < len(s[i]) {
                        break
                    }
                    for i < len(s) && str.Len() + len(s[i]) < 45 {
                        str.WriteString(s[i])
                        i++
                    }
                    s[c] = str.String()
                    str.Reset()
                    c++
                }
                break loop
            }
        }
        s[c] = str.String()
        str.Reset()
        c++
    }
    
    for g := c; g < len(s); g++ {
        s[g] = ""
    }

    s = s[:c]
    return s
}
