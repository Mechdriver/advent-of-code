from pprint import pprint


def find_bags():
    with open("bags.txt", "r", encoding="utf-8") as input_file:
        bag_inputs = list(map(lambda line : str(line).replace('bags', 'bag').strip()[:-1], input_file.readlines()))

    bag_colors = 0
    bags = [bag.split(' contain ') for bag in bag_inputs]
    bags_to_contents = _build_bag_structure(bags)
    return _find_total_bag_contents(bags_to_contents, 'shiny gold bag')

"""
    for bag in bags_to_contents.keys():
        gold_bags = 0
        if bag != 'shiny gold bag':
            gold_bags += _search_bag_structure(bags_to_contents, bag)
        if gold_bags > 0:
            bag_colors += 1
    return bag_colors
"""


def _find_total_bag_contents(bags_to_contents, bag):
    print(bag)
    if not bags_to_contents[bag]:
        return 0

    content_sum = 0
    for bag_to_content in bags_to_contents[bag].items():
        content_sum += (bag_to_content[1] + (bag_to_content[1] * _find_total_bag_contents(bags_to_contents, bag_to_content[0])))

    return content_sum


def _search_bag_structure(bags_to_contents, bag):
    if bag == 'shiny gold bag':
        return 1
    elif not bag:
        return 0
    else:
        gold_bag_sum = 0
        for content_bag in bags_to_contents[bag].keys():
            if content_bag:
                gold_bag_sum += _search_bag_structure(bags_to_contents, content_bag)
        return gold_bag_sum


def _build_bag_structure(bags):
    bag_to_contents = {}
    for bag in bags:
        bag_key = bag[0]
        if bag[1] == 'no other bag':
            bag_to_contents[bag_key] = {}
        else:
            contents = bag[1].split(', ')
            contents_to_amount = {}

            for content in contents:
                space_ndx = content.index(' ')
                contents_to_amount[content[space_ndx:].strip()] = int(content[:space_ndx])

            bag_to_contents[bag_key] = contents_to_amount

    return bag_to_contents

print(find_bags())
