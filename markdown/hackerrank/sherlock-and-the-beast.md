# Sherlock and the beast
# Table of contents
0. [Haskell](#Haskell)

## Haskell
```hs
main = interact $ unlines . map (solve . read) . tail . lines
solve :: Int -> String
solve n = answer where
    answers = [x | x <- [n,n-1..1], x `mod` 3 == 0, (n - x) `mod` 5 == 0]
    answer = if null answers then show (-1) else formatSolution (head answers) n

formatSolution :: Int -> Int -> String
formatSolution answer totalLength = replicate num5 '5' ++ replicate num3 '3' where
    num5 = answer
    num3 = totalLength - answer

```
