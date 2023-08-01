package main
import (
    "fmt"
    "bufio"
    "strconv"
    "regexp"
    "os"
)

func findAnswer(queries []string){
    
    expression1 := `hackerrank$`
    expression2 := `^hackerrank`
    
    for _, query := range queries{
        re := regexp.MustCompile(expression1)
        match1 := re.MatchString(query)
        
        re = regexp.MustCompile(expression2)
        match2 := re.MatchString(query)
        
        if match1 && match2 {
            fmt.Println(0)
            continue
        }
        if match1 {
            fmt.Println(2)
        } else if match2 {
            fmt.Println(1)
        } else {
            fmt.Println(-1)
        }
    }
}

func main() {
    sc := bufio.NewScanner(os.Stdin)
    sc.Scan()
    
    n, _ := strconv.Atoi(sc.Text())
    
    var queries []string
    for i:=0; i<n; i++ {
        sc.Scan()
        queries = append(queries, sc.Text())
    }
    
    findAnswer(queries)
}