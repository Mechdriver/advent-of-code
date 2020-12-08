from pprint import pprint

ACC = 'acc'
JMP = 'jmp'
NOP = 'nop'


class instr_node():
    def __init__(self, instr, val, visited=False):
        self.instr = instr
        self.val = val
        self.visited = visited

    def __repr__(self):
        return f'{self.instr} {self.val}'

def find_loop():
    with open("instructions.txt", "r", encoding="utf-8") as input_file:
        instructions = list(map(lambda line : str(line).strip().split(' '), input_file.readlines()))

    instr_nodes = _build_nodes(instructions)
    swap_mark = acc_sum = 0
    done = False

    while not done:
        acc_sum, done = _find_loop_sum(instr_nodes, swap_mark)
        swap_mark += 1
        for node in instr_nodes:
            node.visited = False


    return acc_sum


def _find_loop_sum(instr_nodes, mark):
    acc_sum = ndx = swap_count = 0
    node = instr_nodes[ndx]

    while not node.visited:
        node.visited = True
        if node.instr == ACC:
            acc_sum += node.val

        if node.instr == NOP and (swap_count == mark) or node.instr == JMP and (swap_count != mark):
                ndx += node.val
        else:
            ndx += 1

        if ndx < len(instr_nodes):
            swap_count += 1
            node = instr_nodes[ndx]
        else:
            return acc_sum, True

    return acc_sum, False


def _build_nodes(instructions):
    instr_nodes = []
    for instr in instructions:
        value = int(instr[1][1:])
        if instr[1][0] == '-':
            value *= -1

        instr_nodes.append(instr_node(instr[0], value))
    return(instr_nodes)

print(find_loop())
