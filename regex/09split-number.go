package main
import (
    "fmt"
    "os"
    "regexp"
    "strconv"
    "bufio"
    "strings"
)
func Split(r rune) bool {
    return r == '-' || r == ' '
}

func findAnswer(queries []string) {
    
    expression :=    `^\d{1,3}[-\s]\d{1,3}[-\s]\d{4,10}$`
    re := regexp.MustCompile(expression)
    
    for _, query := range queries {
        match := re.MatchString(query)
        
        if match{
            a := strings.FieldsFunc(query, Split)
            fmt.Print("CountryCode="+a[0]+",")
            fmt.Print("LocalAreaCode="+a[1]+",")
            fmt.Println("Number="+a[2])
        }
    }
}

func main() {
    sc := bufio.NewScanner(os.Stdin)
    sc.Scan()
    
    n, _ := strconv.Atoi(sc.Text())
    
    var queries []string 
    
    for i := 0; i < n; i++ {
        sc.Scan()
        queries = append(queries, sc.Text())
    }
    
    findAnswer(queries)
}