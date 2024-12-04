import time

# right, down-right, down, down-left, left, up-left, up, up-right
directions = [
    (0, 1), (1, 1), (1, 0), (1, -1),
    (0, -1), (-1, -1), (-1, 0), (-1, 1)
]


def load_file():
    with open('../data/data.txt', 'r') as file:
        return [line.strip() for line in file.readlines()]


def is_valid(x, y, rows: int, cols: int):
    return 0 <= x < rows and 0 <= y < cols


def part_a(grid: list, rows: int, cols: int):
    count = 0

    def check_direction(x: int, y: int, dx: int, dy: int):
        word = ""
        for i in range(4):
            curr_x, curr_y = x + i*dx, y + i*dy
            if not is_valid(curr_x, curr_y, rows, cols):
                return False
            word += grid[curr_x][curr_y]
        return word == "XMAS"

    for x, row in enumerate(grid):
        for y, column in enumerate(row):
            for dx, dy in directions:
                if check_direction(x, y, dx, dy):
                    count += 1
    return count


def part_b(grid: list, rows: int, cols: int):
    count = 0

    def check_direction(x: int, y: int, dx: int, dy: int):
        x = x - dx
        y = y - dy
        word = ""
        for i in range(3):
            curr_x, curr_y = x + i*dx, y + i*dy
            if not is_valid(curr_x, curr_y, rows, cols):
                return False
            word += grid[curr_x][curr_y]
        return word in ("MAS", "SAM")

    def check_x_pattern(x, y):
        # We only want to look from the middle value
        if grid[x][y] != 'A':
            return False

        # Checking on the diags here
        return all([check_direction(x, y, -1, 1), check_direction(x, y, 1, -1),
                   check_direction(x, y, -1, -1), check_direction(x, y, 1, 1)])

    for x, row in enumerate(grid):
        for y, column in enumerate(row):
            if check_x_pattern(x, y):
                count += 1
    return count


if __name__ == '__main__':
    grid = load_file()
    rows = len(grid)
    cols = len(grid[0])
    start = time.time()
    print(part_a(grid, rows, cols))
    print(f'Time Taken: {time.time() - start}')
    start = time.time()
    print(part_b(grid, rows, cols))
    print(f'Time Taken: {time.time() - start}')
