"""
For a given set of numbers, print the largest increasing sequence sum
e.g. 1 2 -3 5 -2 7 --> 15 (1 + 2 + 5 + 7)
"""

numbers = list(map(int, input().split()))

longestIncSeq = list()
maxSeq = numbers[0]

for i in range(len(numbers)):
    currentMaxSeq = numbers[i]
    for j in range(i):
        if numbers[i] + longestIncSeq[j] > currentMaxSeq and numbers[i] > numbers[j]:
            currentMaxSeq = numbers[i] + longestIncSeq[j]

    longestIncSeq.append(currentMaxSeq)
    if (currentMaxSeq > maxSeq):
        maxSeq = currentMaxSeq

print(maxSeq)



