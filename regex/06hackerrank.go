package main
import (
    "fmt"
    "regexp"
    "bufio"
    "os"
    "strconv"
)

func findAnswer(queries []string){
    content := "(C|CPP|JAVA|PYTHON|PERL|PHP|RUBY|CSHARP|HASKELL|CLOJURE|BASH|SCALA|ERLANG|CLISP|LUA|BRAINFUCK|JAVASCRIPT|GO|D|OCAML|R|PASCAL|SBCL|DART|GROOVY|OBJECTIVEC)"
    
    expression := fmt.Sprintf(`^[1-9]\d{4}\s%s$`,content)
    re := regexp.MustCompile(expression)
    
    for _, query := range queries {
        match := re.MatchString(query)
        if match {
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
    
    var queries []string
    for i:=0; i<n; i++{
        sc.Scan()
        queries = append(queries, sc.Text())
    }
    findAnswer(queries)
}
