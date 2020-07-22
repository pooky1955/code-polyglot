# Solve me first
# Table of contents
0. [Go](#Go)
1. [Haskell](#Haskell)
2. [Python](#Python)

## Go
```go
// Submission url : https://www.hackerrank.com/challenges/solve-me-first/submissions/code/168584341
package main
import "fmt"

func solveMeFirst(a uint32,b uint32) uint32{
  return a + b
}

func main() {
    var a, b, res uint32
    fmt.Scanf("%v\n%v", &a,&b)
    res = solveMeFirst(a,b)
    fmt.Println(res)
}
```

## Haskell
```hs
-- Submission url : https://www.hackerrank.com/challenges/solve-me-first/submissions/code/168584255
main = interact $ show . sum . map read . words
```

## Python
```py
# Submission url : https://www.hackerrank.com/challenges/solve-me-first/submissions/code/168661992 
def solveMeFirst(a,b):
	a + b


num1 = int(input())
num2 = int(input())
res = solveMeFirst(num1,num2)
print(res)

```
