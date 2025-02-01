```go
package main

import (

        "fmt"
)

func main() {
        var m map[string]int
        // Check if m is nil before accessing elements
        if m == nil {
                m = make(map[string]int)
        }
        m["key"] = 10
        fmt.Println(m["key"])

        //Alternative solution using the 'ok' idiom
        val, ok := m["anotherkey"]
        if ok {
                fmt.Println("value is", val)
        }else{
                fmt.Println("Key not found")
        }
}
```