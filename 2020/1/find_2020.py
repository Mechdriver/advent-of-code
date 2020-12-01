def get_three_entry_mult():
    goal = 2020

    with open("input_list.txt", "r", encoding="utf-8") as g:
        data = list(map(int, g.readlines()))
        sorted(data)

        for i in range(0, len(data)):
            for j in range(1, len(data)):
                target = goal - data[i] - data[j]
                for k in range(len(data) - 1, -1, -1):
                    if data[k] == target:
                        return data[i] * data[j] * data[k]


def get_entry_mult():
    goal = 2020
    with open("input_list.txt", "r", encoding="utf-8") as g:
        data = list(map(int, g.readlines()))
        sorted(data)

        for entry in data:
            target = goal - entry
            for entry2 in data[::-1]:
                if entry2 == target:
                    return entry * entry2

print(get_three_entry_mult())
