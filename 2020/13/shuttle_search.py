from sympy.ntheory.modular import crt


def find_seq_bus_times_lazy():
    with open("bus_times.txt", "r", encoding="utf-8") as input_file:
        shuttle_data = list(map(lambda line: line.strip(), input_file.readlines()))

    bus_ids = [int(id) if id.isnumeric() else id for id in shuttle_data[1].split(',')]
    ids = [id for id in bus_ids if id != 'x']
    remainders = [id - offset for offset, id in enumerate(bus_ids) if id != 'x']
    print(crt(ids, remainders)[0])


def find_seq_bus_times_brute():
    with open("bus_times.txt", "r", encoding="utf-8") as input_file:
        shuttle_data = list(map(lambda line: line.strip(), input_file.readlines()))

    bus_ids = [int(id) if id.isnumeric() else id for id in shuttle_data[1].split(',')]
    id_to_offset = {id: ndx for ndx, id in enumerate(bus_ids) if id != 'x'}
    done = False
    timestamp = 0

    while not done:
        timestamp += bus_ids[0]
        done = True

        for id, offset in id_to_offset.items():
            if (timestamp + offset) % id:
                done = False
                break

    print(timestamp)




def find_shuttle():
    with open("bus_times.txt", "r", encoding="utf-8") as input_file:
        shuttle_data = list(map(lambda line: line.strip(), input_file.readlines()))

    departure_time = int(shuttle_data[0])
    bus_ids = [int(id) for id in shuttle_data[1].split(',') if id != 'x']
    dep_times_to_id = {}

    for id in bus_ids:
        bus_time = departure_time
        delta = bus_time % id
        bus_time += id - delta
        dep_times_to_id[bus_time] = id

    best_time = min(dep_times_to_id.keys())
    print((best_time - departure_time) * dep_times_to_id[best_time])



find_seq_bus_times_lazy()
#find_seq_bus_times_brute()
