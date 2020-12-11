from pprint import pprint
from copy import deepcopy

FLOOR = '.'
EMPTY = 'L'
FULL = '#'


def find_empty_seats():
    with open("seat_layout.txt", "r", encoding="utf-8") as input_file:
        next_layout = list(map(lambda line: list(line.strip()), input_file.readlines()))

    state_changed = True

    while state_changed:
        state_changed = False
        layout = deepcopy(next_layout)

        for y, row in enumerate(layout):
            for x, space in enumerate(row):
                adj_seats = 0

                # Check left
                if x > 0 and row[x - 1] == FULL:
                    adj_seats += 1
                # Check right
                if x < len(row) - 1 and row[x + 1] == FULL:
                    adj_seats += 1
                # Check up
                if y > 0 and layout[y - 1][x] == FULL:
                    adj_seats += 1
                # Check down
                if y < len(layout) - 1 and layout[y + 1][x] == FULL:
                    adj_seats += 1

                # Check up left
                if y > 0 and x > 0 and layout[y - 1][x - 1] == FULL:
                    adj_seats += 1
                # Check up right
                if y > 0 and x < len(row) - 1 and layout[y - 1][x + 1] == FULL:
                    adj_seats += 1
                # Check down left
                if y < len(layout) - 1 and x > 0 and layout[y + 1][x - 1] == FULL:
                    adj_seats += 1
                # Check down right
                if y < len(layout) - 1 and x < len(row) - 1 and layout[y + 1][x + 1] == FULL:
                    adj_seats += 1

                if space == EMPTY and adj_seats == 0:
                    next_layout[y][x] = FULL
                    state_changed = True
                elif space == FULL and adj_seats >= 4:
                    next_layout[y][x] = EMPTY
                    state_changed = True

    occupied_seats = 0

    for row in layout:
        for space in row:
            if space == FULL:
                occupied_seats += 1

    print(occupied_seats)


find_empty_seats()
