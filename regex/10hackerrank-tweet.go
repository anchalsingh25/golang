package main
import (
    "fmt"
    "os"
    "strconv"
    "bufio"
    "regexp"
)

func findAnswer(queries []string){
    
    expression := `(?i)hackerRank`
    re := regexp.MustCompile(expression)
    total := 0
    for _, query := range queries {
      match := re.MatchString(query)
      if match {
          total += 1;
      }   
    }
    fmt.Println(total)
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