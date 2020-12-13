
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



find_shuttle()
