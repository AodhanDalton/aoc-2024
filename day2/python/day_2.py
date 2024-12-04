def is_safe_sequence(numbers):
    if len(numbers) < 2:
        return False

    increasing = numbers[1] > numbers[0]

    for i in range(len(numbers) - 1):
        diff = numbers[i + 1] - numbers[i]
        if diff == 0:
            return False

        if increasing:
            if diff <= 0 or diff > 3:
                return False
        else:
            if diff >= 0 or diff < -3:
                return False

    return True


def check_with_dampener(numbers):
    if is_safe_sequence(numbers):
        return True

    for i in range(len(numbers)):
        reduced = numbers[:i] + numbers[i+1:]
        if is_safe_sequence(reduced):
            return True

    return False


def part_a():
    safe_count = 0
    with open('../data/data.txt', 'r') as file:
        for line in file:
            numbers = [int(x) for x in line.strip().split()]
            if is_safe_sequence(numbers):
                safe_count += 1
    return safe_count


def part_b():
    safe_count = 0
    with open('../data/data.txt', 'r') as file:
        for line in file:
            numbers = [int(x) for x in line.strip().split()]
            if check_with_dampener(numbers):
                safe_count += 1
    return safe_count


if __name__ == '__main__':
    print(f"Part A: {part_a()}")
    print(f"Part B: {part_b()}")
