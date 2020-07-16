-- Submission url : https://www.hackerrank.com/challenges/compare-the-triplets/submissions/code/168512207
solve :: [Int] -> [Int]
solve xs = let
    as = take 3 xs
    bs = drop 3 xs
    comp_list = [a `compare` b| (a,b) <- zip as bs]
    a_score = length (filter (== GT) comp_list)
    b_score = length (filter (== LT) comp_list)
    in [a_score,b_score]



main = interact $ unwords . map show . solve . map read . words
