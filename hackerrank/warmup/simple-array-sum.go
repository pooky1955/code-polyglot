// Submission url : https://www.hackerrank.com/challenges/simple-array-sum/submissions/code/168585396 
package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the simpleArraySum function below.
 */
func simpleArraySum(ar []int32) int32 {
    var sum int32 = 0
    for _, el := range ar {
        sum += el
    }
    return sum
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    reader.ReadString('\n')
    input,_ := reader.ReadString('\n')
    ar_str := strings.Split(input," ")
    ar := make([]int32,len(ar_str))
    for i,num_str := range ar_str {

        parsed ,_ := strconv.ParseInt(num_str,10,32)
        ar[i] = int32(parsed)
    }
    result := simpleArraySum(ar)
    fmt.Printf("%d\n", result)
}
