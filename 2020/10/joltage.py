def combine_adapters(adapters):
    pass

def chain_adapters(adapters):
    jolt_start = 0
    jolt_diffs = {1: 0, 2: 0, 3: 1}

    for adapter in adapters:
        jolt_diffs[adapter - jolt_start] += 1
        jolt_start = adapter

    return jolt_diffs[1] * jolt_diffs[3]


with open("adapters.txt", "r", encoding="utf-8") as input_file:
    adapters = list(map(int, input_file.readlines()))

adapters = sorted(adapters)
print(chain_adapters(adapters))
