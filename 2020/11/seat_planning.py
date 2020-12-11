from pprint import pprint
from copy import deepcopy
from enum import Enum

SEAT_TOLERANCE_P_1 = 4
SEAT_TOLERANCE_P_2 = 5

FLOOR = '.'
EMPTY = 'L'
FULL = '#'

class Direction(Enum):
    UP = 1
    DOWN = 2
    LEFT = 3
    RIGHT = 4
    UP_LEFT = 5
    UP_RIGHT = 6
    DOWN_LEFT = 7
    DOWN_RIGHT = 9


def _check_seats(layout, direction, y, x):
    # Check left
    if direction == Direction.LEFT and x >= 0:
        if layout[y][x] == FULL:
            return 1
        elif layout[y][x] == FLOOR:
            return _check_seats(layout, Direction.LEFT, y, x - 1)

    # Check right
    if direction == Direction.RIGHT and x < len(layout[y]):
        if layout[y][x] == FULL:
            return 1
        elif layout[y][x] == FLOOR:
            return _check_seats(layout, Direction.RIGHT, y, x + 1)

    # Check up
    if direction == Direction.UP and y >= 0:
        if layout[y][x] == FULL:
            return 1
        elif layout[y][x] == FLOOR:
            return _check_seats(layout, Direction.UP, y - 1, x)

    # Check down
    if direction == Direction.DOWN and y < len(layout):
        if layout[y][x] == FULL:
            return 1
        elif layout[y][x] == FLOOR:
            return _check_seats(layout, Direction.DOWN, y + 1, x)

    # Check up left
    if direction == Direction.UP_LEFT and y >= 0 and x >= 0:
        if layout[y][x] == FULL:
            return 1
        elif layout[y][x] == FLOOR:
            return _check_seats(layout, Direction.UP_LEFT, y - 1, x - 1)

    # Check up right
    if direction == Direction.UP_RIGHT and y >= 0 and x < len(layout[y]):
        if layout[y][x] == FULL:
            return 1
        elif layout[y][x] == FLOOR:
            return _check_seats(layout, Direction.UP_RIGHT, y - 1, x + 1)

    # Check down left
    if direction == Direction.DOWN_LEFT and y < len(layout) and x >= 0:
        if layout[y][x] == FULL:
            return 1
        elif layout[y][x] == FLOOR:
            return _check_seats(layout, Direction.DOWN_LEFT, y + 1, x - 1)

    # Check down right
    if direction == Direction.DOWN_RIGHT and y < len(layout)and x < len(layout[y]):
        if layout[y][x] == FULL:
            return 1
        elif layout[y][x] == FLOOR:
            return _check_seats(layout, Direction.DOWN_RIGHT, y + 1, x + 1)
    return 0


def find_occupied_seats(seat_tolerance):
    with open("seat_layout.txt", "r", encoding="utf-8") as input_file:
        next_layout = list(map(lambda line: list(line.strip()), input_file.readlines()))

    state_changed = True

    while state_changed:
    #for i in range(0, 2):
        state_changed = False
        layout = deepcopy(next_layout)

        for y, row in enumerate(layout):
            for x, space in enumerate(row):
                adj_seats = 0

                # Check left
                adj_seats += _check_seats(layout, Direction.LEFT, y, x - 1)
                # Check right
                adj_seats += _check_seats(layout, Direction.RIGHT, y, x + 1)
                # Check up
                adj_seats += _check_seats(layout, Direction.UP, y - 1, x)
                # Check down
                adj_seats += _check_seats(layout, Direction.DOWN, y + 1, x)
                # Check up left
                adj_seats += _check_seats(layout, Direction.UP_LEFT, y - 1, x - 1)
                # Check up right
                adj_seats += _check_seats(layout, Direction.UP_RIGHT, y - 1, x + 1)
                # Check down left
                adj_seats += _check_seats(layout, Direction.DOWN_LEFT, y + 1, x - 1)
                # Check down right
                adj_seats += _check_seats(layout, Direction.DOWN_RIGHT, y + 1, x + 1)

                if space == EMPTY and adj_seats == 0:
                    next_layout[y][x] = FULL
                    state_changed = True
                elif space == FULL and adj_seats >= seat_tolerance:
                    next_layout[y][x] = EMPTY
                    state_changed = True

    occupied_seats = 0
    for row in layout:
        for space in row:
            if space == FULL:
                occupied_seats += 1

    print(occupied_seats)


find_occupied_seats(SEAT_TOLERANCE_P_2)
