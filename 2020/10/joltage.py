from pprint import pprint

def combine_adapters(adapters):
    return _find_paths(adapters, {}, -1)


def _find_paths(adapters, cache, ndx):
    if ndx == len(adapters) - 1:
        return 1

    if ndx == -1:
        curr_adapt = 0
    else:
        curr_adapt = adapters[ndx]

    if curr_adapt in cache:
        return cache[curr_adapt]

    paths_num = _find_paths(adapters, cache, ndx + 1)

    if (ndx + 1 < len(adapters) and ndx + 2 < len(adapters) and
        adapters[ndx + 1] - curr_adapt in [1, 2] and
        adapters[ndx + 2] - curr_adapt < 4):

        paths_num += _find_paths(adapters, cache, ndx + 2)

    if (ndx + 2 < len(adapters) and ndx + 3 < len(adapters) and
        adapters[ndx + 2] - curr_adapt in [1, 2] and
        adapters[ndx + 3] - curr_adapt < 4):
        
        paths_num += _find_paths(adapters, cache, ndx + 3)

    cache[curr_adapt] = paths_num

    return paths_num


def chain_adapters(adapters):
    jolt_start = 0
    jolt_diffs = {1: 0, 2: 0, 3: 1}

    for adapter in adapters:
        jolt_diffs[adapter - jolt_start] += 1
        jolt_start = adapter

    return jolt_diffs[1], jolt_diffs[3]


with open("adapters.txt", "r", encoding="utf-8") as input_file:
    adapters = list(map(int, input_file.readlines()))

adapters = sorted(adapters)
print(combine_adapters(adapters))
