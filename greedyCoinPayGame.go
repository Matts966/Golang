package main
import (
    "fmt"
    "math"
    )
func main(){
    var a int
    fmt.Scan(&a)
    ret := ""
    for i := 9; i >= 0; i-- {
        var b int
        b = int(a / int(math.Pow(2.0,float64(i))))
        a = int(a % int(math.Pow(2.0,float64(i))))
        ret = fmt.Sprintf("%d",b) + ret
    }
    fmt.Println(ret+"\n")
}
