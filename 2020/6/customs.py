def get_unanimous_sum():
    with open("custom_forms.txt", "r", encoding="utf-8") as input_file:
        custom_forms = list(map(lambda line : str(line).strip(), input_file.readlines()))

    custom_ans = []
    groups = []
    group = []
    unanimous_sum = 0

    for form in custom_forms:
        if not form:
            groups.append(group)
            group = []
        else:
            group.append([ans for ans in form])

    groups.append(group)

    for group in groups:
        ans_dict = {}
        for form in group:
            for ans in form:
                if ans not in ans_dict:
                    ans_dict[ans] = 1
                else:
                    ans_dict[ans] += 1
        for item in ans_dict.items():
            if item[1] == len(group):
                unanimous_sum += 1
    return unanimous_sum


def get_unique_sum():
    with open("custom_forms.txt", "r", encoding="utf-8") as input_file:
        custom_forms = list(map(lambda line : str(line).strip(), input_file.readlines()))

    custom_ans = []
    groups = []
    sum = 0

    for form in custom_forms:
        if not form:
            unique_ans = set([])
            for group in groups:
                for ans in group:
                    unique_ans.add(ans)
            custom_ans.append(unique_ans)
            groups = []
        else:
            group = [ans for ans in form]
            groups.append(group)

    unique_ans = set([])
    for group in groups:
        for ans in group:
            unique_ans.add(ans)
    custom_ans.append(unique_ans)

    for answers_set in custom_ans:
        sum += len(answers_set)

    return sum

print(get_unanimous_sum())
