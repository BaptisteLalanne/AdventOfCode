

def solver_problem1(data):
    increase = sum([1 for i in range(1, len(data)) if data[i] > data[i-1]])
    return increase


def solver_problem1_2(data):
    increase = sum([1 for i in range(3, len(data)) if data[i] > data[i-3]])
    return increase


if __name__ == '__main__':
    with open("./inputs/day1.txt", 'r') as file:
        data = [int(line.strip()) for line in file.readlines()]

    print(solver_problem1_2(data))
