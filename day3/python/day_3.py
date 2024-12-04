import re
import time
mul_regex = r'mul\(([0-9]+).([0-9]+)\)'


def load_file():
    with open('../data/data.txt', 'r') as file:
        data = file.readlines()
    return ''.join(data)


def part_a(file_data: list):
    muls = []
    total = 0
    muls = re.findall(mul_regex, file_data)
    for left, right in muls:
        total += int(left) * int(right)
    return total


def part_b(file_data: list):
    total = 0
    enabled = True
    pos = 0

    while pos < len(file_data):
        if file_data[pos:].startswith("don't()"):
            enabled = False
            pos += 7
            continue
        if file_data[pos:].startswith("do()"):
            enabled = True
            pos += 4
            continue

        if enabled:
            match = re.match(mul_regex, file_data[pos:])
            if match:
                left, right = match.groups()
                total += int(left) * int(right)
                pos += match.end()
            else:
                pos += 1
        else:
            pos += 1
    return total


if __name__ == '__main__':
    file_data = load_file().replace('\n', '')
    start = time.time()
    print(part_a(file_data))
    print(f'Time Taken: {time.time() - start}')
    start = time.time()
    print(part_b(file_data))
    print(f'Time Taken: {time.time() - start}')
