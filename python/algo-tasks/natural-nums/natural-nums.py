"""
For a given natural number, print out its possible
split in other natural numbers.
"""

import math

def split(num : int, previousTerm : int):
    if previousTerm != 0:
        splits = [[previousTerm]]
    else:
        splits = [[]]
    splits[0].append(num)

    for i in range(2, int(math.sqrt(num)) + 1):
        term1 = i
        if num % term1 == 0:
            term2 = int(num / term1)

            term2Splits = split(term2, 0)
            for splitList in term2Splits:
                splits.append([ term1 ] + splitList)

    return splits

def printSplits(splits : list, origNum : int):
    for terms in splits:
        print(" * ".join(str(x) for x in terms), end="")
        print(f" = {origNum}")

print("Input a natural number: ", end="")
num = int(input())

if num < 1:
    exit("Invalid natural number")

printSplits(split(num, 1), num)

