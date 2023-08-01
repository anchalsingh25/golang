package main

import (
    "fmt"
    "regexp"
    "strings"
    "bufio"
    "os"
    "strconv"
    "sort"
)

func main() {
  
    sc := bufio.NewScanner(os.Stdin)
    sc.Scan()
    n, _ := strconv.Atoi(sc.Text())

    htmlContent := ""
    for i := 0; i < n; i++ {
        sc.Scan()
        line := sc.Text()
        htmlContent += line + "\n"
    }
    
	tagPattern := `<\s*([^/>\s]+)[^>]*>`
    re := regexp.MustCompile(tagPattern)
    matches := re.FindAllStringSubmatch(htmlContent, -1)
    
    tagSet := make(map[string]bool)
    for _, match := range matches {
        tagName := strings.TrimSpace(match[1])
        tagSet[tagName] = true
    }
    
    var tags []string
    for tag := range tagSet {
        tags = append(tags, tag)
    }
    sort.Strings(tags)
    output := strings.Join(tags, ";")
    fmt.Print(output)
}
