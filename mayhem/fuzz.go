package fuzz

import "strconv"
import "benchyou/xcommon"

func mayhemit(bytes []byte) int {

    var num int
    if len(bytes) < 1 {
        num = 0
    } else {
        num, _ = strconv.Atoi(string(bytes[0]))
    }

    switch num {
    
    case 0:
        content := string(bytes)
        xcommon.RandString(content)
        return 0

    default:
        content := string(bytes)
        xcommon.MockConf(content)
        return 0
    }
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}