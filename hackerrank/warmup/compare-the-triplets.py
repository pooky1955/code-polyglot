# Submission url : https://www.hackerrank.com/challenges/compare-the-triplets/submissions/code/168586003
def solve(alst,blst):
    a, b = 0,0
    for a_i, b_i in zip(alst,blst):
        if a_i  > b_i:
            a += 1
        elif b_i > a_i:
            b += 1
    
    return [a,b]

if __name__ == '__main__':
    alst = map(int,input().split(" "))
    blst = map(int,input().split(" "))

    result = solve(alst,blst)
    result_str = map(str,result)
    print(" ".join(result_str))

