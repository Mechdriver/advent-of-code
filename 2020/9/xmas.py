WINDOW_SIZE = 25

def crack_encryption():
    with open("preamble.txt", "r", encoding="utf-8") as input_file:
        nums = list(map(int, input_file.readlines()))

    cracked = False
    contiguous_set_len = 2
    contiguous_set = []
    bad_num = _find_bad_number(nums)
    search_nums = nums[:nums.index(bad_num)]

    while not cracked:
        for ndx in range(0, len(search_nums) - (contiguous_set_len - 1)):
            contiguous_set = search_nums[ndx:ndx + contiguous_set_len]
            if sum(contiguous_set) == bad_num:
                cracked = True
                break

        contiguous_set_len += 1
    return min(contiguous_set) + max(contiguous_set)


def _find_bad_number(nums):
    for ndx in range(0, len(nums) - WINDOW_SIZE):
        done = True
        target_num = nums[ndx + WINDOW_SIZE]
        nums_slice = nums[ndx:ndx + WINDOW_SIZE]
        num_dict = {num: True for num in nums_slice}

        for num in nums_slice:
            if target_num - num in num_dict:
                done = False
                break
        if done:
            return target_num


print(crack_encryption())
