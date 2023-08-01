package main
import (
    "fmt"
    "regexp"
    "bufio"
    "strconv"
    "os"
)

func findAnswer(content string, queries []string){
    
    for _, query := range queries {
        expression := fmt.Sprintf(`\w+%s\w+`,query)
        re := regexp.MustCompile(expression) 
        match := re.FindAllString(content, -1)
        fmt.Println(len(match))  
    }    
}

func main() {
    sc := bufio.NewScanner(os.Stdin)
    sc.Scan()
    
    n, _ := strconv.Atoi(sc.Text())
    
    content := ""
    
    for i := 0; i<n; i++ {
        sc.Scan()
        lines := sc.Text()
        content += lines + "\n"
    }
    
    sc.Scan()
    q, _ := strconv.Atoi(sc.Text())
    var queries []string
    
    for i:=0; i<q; i++ {
        sc.Scan()
        queries = append(queries, sc.Text());
    }
    
    findAnswer(content, queries)
}
