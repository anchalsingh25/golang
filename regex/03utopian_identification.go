package main
import (
    "fmt"
    "bufio"
    "strconv"
    "os"
    "regexp"
)
func findAnswer(lines []string){
    
    expression := `^[a-z]{0,3}\d{2,8}[A-Z]{3,}$`
    re := regexp.MustCompile(expression) 
        
    for _, line := range lines {
        match := re.MatchString(line)
      
        if match{
            fmt.Println("VALID")
        }else {
            fmt.Println("INVALID")
        } 
        
    }    
}

func main() {
    sc := bufio.NewScanner(os.Stdin)
    sc.Scan()
    n, _ := strconv.Atoi(sc.Text())
    
    var lines []string
    
    for i := 0; i<n; i++ {
        sc.Scan()
        lines = append(lines, sc.Text())
    }
    
    findAnswer(lines)
}
