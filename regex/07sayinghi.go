package main
import (
    "fmt"
    "strconv"
    "os"
    "bufio"
    "regexp"
)

func findAnswer(lines []string){
    
    var ans []string
    expression := `^[Hh][Ii][\s][^Dd]`
    re := regexp.MustCompile(expression)
    
    for _, query := range lines {
        match := re.MatchString(query)
        if match {
            ans = append(ans, query)
        }
    }
    
    for _, pri := range ans {
        fmt.Println(pri)
    }
}


func main() {
    sc := bufio.NewScanner(os.Stdin)
    sc.Scan()
    
    n, _ := strconv.Atoi(sc.Text())
    
    var lines []string
    
    for i:=0; i<n; i++ {
        sc.Scan()
        lines = append(lines, sc.Text())
    }
    findAnswer(lines)
}