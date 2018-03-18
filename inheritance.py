# Day 12: Inheritance - Done in python because the challenge isn't supported in Go

class Person:
    def __init__(self, firstName, lastName, idNumber):
        self.firstName = firstName
        self.lastName = lastName
        self.idNumber = idNumber
    def printPerson(self):
        print ("Name:", self.lastName + ",", self.firstName)
        print ("ID:", self.idNumber)


class Student(Person):
    GRADES = [
        ('O', 90, 100),
        ('E', 80, 89),
        ('A', 70, 79),
        ('P', 55, 69),
        ('D', 40, 54),
        ('T', 0, 39),
    ]

    def __init__(self, firstName, lastName, idNumber, scores):
        super(Student, self).__init__(firstName, lastName, idNumber)
        self.scores = scores

    def calculate(self):
        average = sum(self.scores) / len(self.scores)
        for grade, min_score, max_score in self.GRADES:
            if average >= min_score and average <= max_score:
                return grade


line = input().split()
firstName = line[0]
lastName = line[1]
idNum = line[2]
numScores = int(input()) # not needed for Python
scores = list(map(int, input().split()))
s = Student(firstName, lastName, idNum, scores)
s.printPerson()
print ("Grade:", s.calculate())
