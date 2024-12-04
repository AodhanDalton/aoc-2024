import time

left = []
right = []
with open('../data/data.txt', 'r') as file:
    for line in file.readlines():
        l, r = line.strip().split('   ')
        left.append(int(l))
        right.append(int(r))
left.sort()
right.sort()


def part_a():
    total_distance = 0
    for left_num, right_num in zip(left, right):
        if left_num > right_num:
            total_distance += left_num - right_num
        elif right_num > left_num:
            total_distance += right_num - left_num
    return total_distance


def part_b():
    similarity_score = 0
    for left_num in left:
        count = right.count(left_num)
        similarity_score += left_num * count
    return similarity_score


if __name__ == '__main__':
    start = time.time()
    print('_________')
    print(f'A: {part_a()}')
    print(f'B: {part_b()}')
    print(f'Time taken: {time.time() - start}')
    print('_________')
