import time
from itertools import product


def load_file():
    results = []
    with open('../data/data.txt', 'r') as file:
        for line in file:
            result, lst = line.strip().split(':')
            result = int(result)
            lst = [int(x) for x in lst.strip().split()]
            results.append((result, lst))
    return results


def check_combination(numbers: list, operators: list):
    """ Evaluate a single combination of numbers and operators """
    result = numbers[0]
    for i, op in enumerate(operators):
        if op == '+':
            result += numbers[i + 1]
        elif op == '*':
            result *= numbers[i + 1]
        elif op == '||':
            result = int(str(result) + str(numbers[i + 1]))
    return result


def find_target_combination(target: int, numbers: list, operators: list) -> bool:
    """ Check if any combination of operators can reach the target """
    for ops in product(operators, repeat=len(numbers)-1):
        try:
            if check_combination(numbers, ops) == target:
                return True
        except Exception as error:
            print(error)
            continue
    return False


def part_a(equations: list):
    total = 0
    for target, numbers in equations:
        if find_target_combination(target, numbers, ['+', '*']):
            total += target
    return total


def part_b(equations: list):
    total = 0
    for target, numbers in equations:
        if find_target_combination(target, numbers, ['+', '*', '||']):
            total += target
    return total


if __name__ == '__main__':
    equations = load_file()
    start = time.time()
    result_a = part_a(equations)
    time_a = time.time() - start
    print(f'Part A Result: {result_a}')
    print(f'Part A Time: {time_a:.2f} seconds')
    start = time.time()
    result_b = part_b(equations)
    time_b = time.time() - start
    print(f'Part B Result: {result_b}')
    print(f'Part B Time: {time_b:.2f} seconds')
