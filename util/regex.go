package util

import (
    "regexp"
)

func RegexpCompileSimple(spec string) *regexp.Regexp {
    regex, err := regexp.Compile(spec)
    Check(err)
    return regex
}
