import time
OBSTICLE = '#'


def load_file():
    with open('../data/data.txt', 'r') as file:
        return [line.strip() for line in file.readlines()]


def is_valid(x, y):
    return 0 <= x < len(grid) and 0 <= y < len(grid[0])


def simulate_path(grid: list, x: int, y: int, direction: int):
    visited = set()
    position_with_direction = set()
    directions = [(-1, 0), (0, 1), (1, 0), (0, -1)]

    while True:
        visited.add((x, y))
        curr_state = (x, y, direction)

        if curr_state in position_with_direction:
            return visited, True
        position_with_direction.add(curr_state)

        next_x = x + directions[direction][0]
        next_y = y + directions[direction][1]

        if not is_valid(next_x, next_y):
            return visited, False

        if grid[next_x][next_y] == OBSTICLE:
            direction = (direction + 1) % 4
        else:
            x = next_x
            y = next_y

    return visited, False


def get_start(grid: list):
    for i, row in enumerate(grid):
        if '^' in row:
            start_x = i
            start_y = row.index('^')
            return start_x, start_y


def part_a(grid: list):
    visited = set()
    x, y = get_start(grid)

    def move(x, y, direction):
        directions = [(-1, 0), (0, 1), (1, 0), (0, -1)]
        dx, dy = directions[direction]
        while True:
            visited.add((x, y))
            next_x = x + directions[direction][0]
            next_y = y + directions[direction][1]

            if not is_valid(next_x, next_y):
                return len(visited)

            if grid[next_x][next_y] == OBSTICLE:
                direction = (direction + 1) % 4
            else:
                x = next_x
                y = next_y

    return move(x, y, direction=0)


def part_b(grid: list):
    start_x, start_y = None, None
    loop_positions = 0
    total_positions = len(grid) * len(grid[0])
    positions_checked = 0

    start_x, start_y = get_start(grid)

    grid = [list(row) for row in grid]
    for i in range(len(grid)):
        for j in range(len(grid[0])):
            if grid[i][j] != '.' or (i == start_x and j == start_y):
                continue
            positions_checked += 1
            if positions_checked % 100 == 0:  # Print progress every 100 positions
                print(f"Checking position {positions_checked}/{total_positions}")
            # Simulate road block
            grid[i][j] = '#'
            _, forms_loop = simulate_path(grid, start_x, start_y, 0)
            if forms_loop:
                loop_positions += 1
            # Undo simulation
            grid[i][j] = '.'
    return loop_positions


if __name__ == '__main__':
    grid = load_file()

    start = time.time()
    print(part_a(grid))
    print(f'Time Taken: {time.time() - start}')

    start = time.time()
    print(part_b(grid))
    print(f'Time Taken: {time.time() - start}')
