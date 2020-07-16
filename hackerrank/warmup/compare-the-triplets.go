// Submission url : https://www.hackerrank.com/challenges/compare-the-triplets/submissions/code/168520494 
package main

import (
    "bufio"
    "fmt"
    // "io"
    "os"
    "strconv"
    "strings"
)

// Complete the compareTriplets function below.
func compareTriplets(as []int, bs []int) []int {
    a_score , b_score  := 0, 0
    for i := 0; i < len(as); i++ {
        a,b := as[i], bs[i]
        if a > b {
            a_score += 1
        } else if a < b{
            b_score += 1
        }
    }
    result := make([]int,2)
    result[0] = a_score
    result[1] = b_score
    return result

}

func processLine(s string) []int{
    splitted := strings.Fields(s)
    results := make([]int,len(splitted))
    for i, term := range splitted {
        parsed,_ := strconv.ParseInt(term,10,32)
        results[i] = int(parsed)
    }
    return results
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    as_s,_ := reader.ReadString('\n')
    bs_s,_ := reader.ReadString('\n')
    as := processLine(as_s)
    bs := processLine(bs_s)
    result := compareTriplets(as,bs)
    fmt.Printf("%v %v",result[0],result[1])
    
}
