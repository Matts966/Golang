package main
import (
    "fmt"
    "unicode/utf8"
    )
func main(){
    // Here your code !
    var a int
    fmt.Scan(&a)
    s := ""
    s = fmt.Sprintf("%b", a)
    var size int
    size = utf8.RuneCountInString(s)
    ret := ""
    for i := 0; i < size; i++ {
        ret = string(s[i]) + ret
    }
    size = utf8.RuneCountInString(ret)
    for i := 0; i < 10-size; i++ {
        ret = ret+"0"
    }
    fmt.Println(ret+"\n")
}

