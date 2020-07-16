-- Submission url : https://www.hackerrank.com/challenges/simple-array-sum/submissions/code/168584496
main = interact $ show . sum . map read . tail . words