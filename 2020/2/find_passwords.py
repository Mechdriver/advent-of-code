def get_valid_toboggan_passwords():
    with open("passwords.txt", "r", encoding="utf-8") as input_file:
        data = list(map(lambda line : str(line).strip(), input_file.readlines()))

    valid_passwords = []
    policy_passes = [line.split(' ') for line in data]

    for policy_pass in policy_passes:
        ranges = policy_pass[0].split('-')
        ndx1 = int(ranges[0]) - 1
        ndx2 = int(ranges[1]) - 1
        restricted_char = policy_pass[1][0]
        password = policy_pass[2]

        if (password[ndx1] == restricted_char) != (password[ndx2] == restricted_char):
            valid_passwords.append(password)

    return valid_passwords

def get_valid_passwords():
    with open("passwords.txt", "r", encoding="utf-8") as input_file:
        data = list(map(lambda line : str(line).strip(), input_file.readlines()))
    valid_passwords = []
    policy_passes = [line.split(' ') for line in data]

    for policy_pass in policy_passes:
        ranges = policy_pass[0].split('-')
        min = int(ranges[0])
        max = int(ranges[1])
        restricted_char = policy_pass[1][0]
        password = policy_pass[2]

        char_count = password.count(restricted_char)
        if char_count >= min and char_count <= max:
            valid_passwords.append(password)

    return valid_passwords


print(len(get_valid_toboggan_passwords()))
