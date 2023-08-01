package main
import (
    "fmt"
    "regexp"
)

func main() {
    var n int
    fmt.Scan(&n)
    
    expression := `^[A-Z]{5}\d{4}[A-Z]{1}$`
    re := regexp.MustCompile(expression) 
    
    for i:=0; i<n; i++ {
        var str string
        fmt.Scan(&str)
        match := re.MatchString(str)
        if match {
            fmt.Println("YES")
        } else{
            fmt.Println("NO")
        }
    }
}
