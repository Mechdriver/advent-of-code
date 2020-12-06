import re

REQ_FIELDS = 8

def find_valid_passports():
    with open("id_batch.txt", "r", encoding="utf-8") as input_file:
        data = list(map(lambda line : str(line).strip(), input_file.readlines()))

    valid_passports = []
    clean_data = _clean_data(data)
    passports = _build_passports(clean_data)

    for passport in passports:
        if len(passport) == REQ_FIELDS or (len(passport) == REQ_FIELDS - 1 and 'cid' not in passport):
            valid_passports.append(passport)

    return _filter_invalid_data(valid_passports)


def _filter_invalid_data(passports):
    eye_colors = ['amb', 'blu', 'brn', 'gry', 'grn', 'hzl', 'oth']

    valid_passports = []
    for passport in passports:
        # Birth Year
        if int(passport['byr']) not in range(1920, 2002 + 1):
            continue

        # Issued Year
        if int(passport['iyr']) not in range(2010, 2020 + 1):
            continue

        # Expiration Year
        if int(passport['eyr']) not in range (2020, 2030 + 1):
            continue

        # Height
        if 'cm' in passport['hgt']:
            height = int(passport['hgt'][:-2])
            if height not in range(150, 193 + 1):
                continue
        elif 'in' in passport['hgt']:
            height = int(passport['hgt'][:-2])
            if height not in range(59, 76 + 1):
                continue
        else:
            continue

        # Hair Color
        hair_regex = re.compile(r'#[0-9a-f]{6}')
        if not hair_regex.match(passport['hcl']):
            continue

        # Eye Color
        if passport['ecl'] not in eye_colors:
            continue

        # Passport ID
        if len(passport['pid']) != 9 or not passport['pid'].isnumeric():
            continue

        valid_passports.append(passport)
    return valid_passports



def _build_passports(clean_data):
    passports = []
    for raw_passport in clean_data:
        passport = {}
        for info in raw_passport:
            k_v_pair = info.split(':')
            passport[k_v_pair[0]] = k_v_pair[1]
        passports.append(passport)

    return passports


def _clean_data(data):
    cleaned_data = [['']]
    clean_pos = 0

    for line in data:
        if line:
            cleaned_data[clean_pos][0] += ' ' + line
        else:
            clean_pos += 1
            cleaned_data.append([''])
    cleaned_data = [data[0].strip() for data in cleaned_data]

    return [passport_data.split(' ') for passport_data in cleaned_data]

print(len(find_valid_passports()))
