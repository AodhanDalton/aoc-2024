import time
page_order_regex = r'\d+'


def read_sections():
    first_section = []
    second_section = []
    with open('../data/data.txt', 'r') as file:
        # Read first section until empty line
        for line in file:
            if line.strip() == '':
                break
            first_section.append(line.strip())
        second_section = [line.strip() for line in file if line.strip()]
    return first_section, second_section


def is_valid_update(update, rules):
    for i, num1 in enumerate(update):
        for j, num2 in enumerate(update[i+1:], i+1):
            if num2 + '|' + num1 in rules:
                return False
    return True


def compare_pages(page1, page2, rules):
    if f"{page1}|{page2}" in rules:
        return -1
    if f"{page2}|{page1}" in rules:
        return 1
    return 0


def part_a(page_orders: list, updates: list):
    rules = set(page_orders)
    total = 0
    for update in updates:
        if is_valid_update(update, rules):
            middle_idx = len(update) // 2
            total += int(update[middle_idx])
    return total


def part_b(page_orders: list, updates: list):
    rules = set(page_orders)
    total = 0
    for update in updates:
        # Don't care about these
        if is_valid_update(update, rules):
            continue
        sorted_pages = sorted(update, key=lambda x: sum(compare_pages(x, y, rules) for y in update))
        total += int(sorted_pages[len(sorted_pages)//2])
    return total


if __name__ == '__main__':
    page_orders, updates = read_sections()
    updates = [update.split(',') for update in updates]
    start = time.time()
    print(part_a(page_orders, updates))
    print(f'Time Taken: {time.time() - start}')
    start = time.time()
    print(part_b(page_orders, updates))
    print(f'Time Taken: {time.time() - start}')
