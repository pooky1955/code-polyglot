# A very big sum
# Table of contents
0. [Haskell](#Haskell)

## Haskell
```hs
-- Submission url : https://www.hackerrank.com/challenges/a-very-big-sum/submissions/code/168540942
main = interact $ show . sum . map read . tail . words

```
