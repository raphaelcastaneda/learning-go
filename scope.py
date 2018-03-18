# Day 14: Scope - Solved in Python because this challenge does not support Go

class Difference:
    def __init__(self, a):
        self.__elements = a

        self.maximumDifference = None

    def computeDifference(self):
        for i, n in enumerate(self.__elements):
            for m in self.__elements[i:]:
                difference = abs(n - m)
                if self.maximumDifference is None or difference > self.maximumDifference:
                    self.maximumDifference = difference


_ = input()
a = [int(e) for e in input().split(' ')]

d = Difference(a)
d.computeDifference()

print(d.maximumDifference)
