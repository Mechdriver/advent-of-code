TREE = '#'

def find_tree_collisions(col_inc, row_inc):
    with open("map.txt", "r", encoding="utf-8") as input_file:
        tree_map = list(map(lambda line : str(line).strip(), input_file.readlines()))

    trees_hit = 0
    col = 0

    for row in range(0, len(tree_map), row_inc):
        if tree_map[row][col] == TREE:
            print(f'Row: {row} Col: {col}')
            trees_hit += 1

        col += col_inc
        if col >= len(tree_map[row]):
            col = col - (len(tree_map[row]))

    return trees_hit

print(find_tree_collisions(1, 1) * find_tree_collisions(3, 1) * find_tree_collisions(5, 1) * find_tree_collisions(7, 1) * find_tree_collisions(1, 2))
