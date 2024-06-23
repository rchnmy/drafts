package splitafterfieldsfunc

import (
    "strings"
)

// splitAfterFieldsFunc is equivalent to `strings.SplitAfter` for multiple separators.
// If `seps` is empty, splitAfterFieldsFunc returns an empty slice.
func splitAfterFieldsFunc(s string, seps ...string) []string {
    switch s {
    case "":
        return nil    
    }
    
    var se int
    for _, sep := range seps {
        for _, r := range s {
            switch sep {
                case string(r):
                    se++
            }
        }
    }

    spl := make([]string, se + 1, len(s))
    str := strings.Builder{}
    i := 0
    for i < se {
        for _, sep := range seps {
            str.WriteString(sep)
        }
        is := strings.IndexAny(s, str.String())
        str.Reset()
        if is < 0 {
            break
        }
        spl[i] = s[:is + 1]
        s = s[is + 1:]
        i++
    }
    spl[i] = s
    return spl[:i]
}
