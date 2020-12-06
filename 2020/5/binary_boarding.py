def find_my_seat():
    sorted_seat_ids = sorted(decrypt_passes())

    for ndx in range(0, len(sorted_seat_ids)):
        if sorted_seat_ids[ndx+1] != sorted_seat_ids[ndx] + 1:
            return sorted_seat_ids[ndx] + 1


def decrypt_passes():
    with open("boarding_passes.txt", "r", encoding="utf-8") as input_file:
        boarding_passes = list(map(lambda line : str(line).strip(), input_file.readlines()))

    seat_ids = []

    for boarding_pass in boarding_passes:
        seat_ids.append(decrypt_boarding_pass(boarding_pass))

    return seat_ids


def decrypt_boarding_pass(boarding_pass):
    row = _get_row(boarding_pass)
    col = _get_col(boarding_pass)
    seat_id = row * 8 + col

    return seat_id

def _get_row(boarding_pass):
    row_min = 0
    row_max = 128

    row_instr = boarding_pass[:-3]
    for instr in row_instr:
        rows = list(range(row_min, row_max))
        if instr == 'F':
            if len(rows) == 2:
                row = rows[0]
            else:
                row_max = rows[len(rows) // 2]
        elif instr == 'B':
            if len(rows) == 2:
                row = rows[1]
            else:
                row_min = rows[len(rows) // 2]
    return row

def _get_col(boarding_pass):
    col_min = 0
    col_max = 8

    col_instr = boarding_pass[7:]
    for instr in col_instr:
        cols = list(range(col_min, col_max))
        if instr == 'L':
            if len(cols) == 2:
                col = cols[0]
            else:
                col_max = cols[len(cols) // 2]
        elif instr == 'R':
            if len(cols) == 2:
                col = cols[1]
            else:
                col_min = cols[len(cols) // 2]
    return col


#decrypt_boarding_pass('BBFFBBFRLL')
#print(decrypt_passes())
print(find_my_seat())
