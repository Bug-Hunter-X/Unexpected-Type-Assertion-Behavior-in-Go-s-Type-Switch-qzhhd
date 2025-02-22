```go
package main

import "fmt"

func main() {
    var i interface{} = 10
    switch i.(type) {
    case int:
        fmt.Println("int")
    case float64:
        fmt.Println("float64")
    default:
        fmt.Println("default")
    }

    var j interface{} = 10.0
    switch j.(type) {
    case int:
        fmt.Println("int")
    case float64:
        fmt.Println("float64")
    default:
        fmt.Println("default")
    }
}
```