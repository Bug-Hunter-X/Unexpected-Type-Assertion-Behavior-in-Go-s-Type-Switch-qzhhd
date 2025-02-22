```go
package main

import (

        "fmt"
        "reflect"
)

func main() {
        var i interface{} = 10
        switch reflect.TypeOf(i).Kind() {
        case reflect.Int:
                fmt.Println("int")
        case reflect.Float64:
                fmt.Println("float64")
        default:
                fmt.Println("default")
        }

        var j interface{} = 10.0
        switch reflect.TypeOf(j).Kind() {
        case reflect.Int:
                fmt.Println("int")
        case reflect.Float64:
                fmt.Println("float64")
        default:
                fmt.Println("default")
        }
}
```