WINDOW_SIZE = 25

def find_bad_number():
    with open("preamble.txt", "r", encoding="utf-8") as input_file:
        nums = list(map(int, input_file.readlines()))


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


print(find_bad_number())
