import argparse
import csv

parser = argparse.ArgumentParser()
parser.add_argument("-f", "--filename", default="problems.csv", help="name of csv file with structure 'question, answer'")
parser.add_argument("-l", "--limit", type=int, default=5, help="time to answer a question, in seconds")
args = parser.parse_args()
filename = args.filename
limit = args.limit

with open(filename) as csvfile:
    try:
        problems_csv = csv.reader(csvfile)
    except:
        print("Failed to read csv file: %s", filename)
    problems = []
    for q,a in problems_csv:
        problems.append((q,a))

correct = 0
i = 0
for q,a in problems:
    i += 1
    print("Problem #{}: {} = ".format(i, q))
    try:
        answer = str(input())
    except (SyntaxError, NameError):
        print("-> no valid answer recorded <-")
        answer = ""
        pass
    if answer == a:
        correct += 1
print("You answered {} questions of {} correctly".format(correct, len(problems)))
    
