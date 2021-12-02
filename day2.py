
if __name__ == '__main__':
    with open("./inputs/day2.txt", 'r') as file:
        data = [line.strip().split(" ") for line in file.readlines()]
        data = [(x, int(y)) for [x, y] in data]
    print(data)
    horizontal = 0
    vertical = 0
    aim = 0
    depth = 0

    for move, measure in data:
        if move == 'forward':
            horizontal += measure
            depth += measure*aim
        elif move == 'up':
            vertical -= measure
            aim -= measure
        elif move == 'down':
            vertical += measure
            aim += measure
    print("Depth (part1)",horizontal*vertical, "Depth (part2): ", depth*horizontal)
