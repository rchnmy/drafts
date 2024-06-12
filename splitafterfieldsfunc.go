package splitafterfieldsfunc

import (
    "strings"
)

// SplitAfterFieldsFunc is equivalent to `strings.SplitAfter` for multiple separators.
// If `seps` is empty, SplitAfterFieldsFunc returns an empty slice.
func SplitAfterFieldsFunc(s string, seps ...string) []string {
    var se int
    for _, sep := range seps {
        for _, r := range s {
            switch sep {
                case string(r):
                    se++
            }
        }
    }

    spl := make([]string, se+1, len(s))
    str := strings.Builder{}
    i := 0
    for i < se {
        defer str.Reset()
        for _, sep := range seps {
            str.WriteString(sep)
        }
        is := strings.IndexAny(s, str.String())
        switch {
            case is < 0:
                break
        }
        spl[i] = s[:is+1]
        s = s[is+1:]
        i++
    }
    spl[i] = s
    return spl[:i]
}
