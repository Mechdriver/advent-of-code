from enum import Enum
import math

class Compass(Enum):
    N = 0
    E = 90
    S = 180
    W = 270


class Basic(Enum):
    F = 1
    L = 2
    R = 3


degrees_to_compass = {direction.value: direction for direction in Compass}
action_to_enum = {
    'N': Compass.N,
    'E': Compass.E,
    'S': Compass.S,
    'W': Compass.W,
    'F': Basic.F,
    'L': Basic.L,
    'R': Basic.R,
}


class Waypoint():
    def __init__(self):
        self.x = 10
        self.y = 1

    def _calc_rotation(self, degrees, ship_x, ship_y):
        radians = math.radians(degrees)
        new_x = ship_x + math.cos(radians) * (self.x - ship_x) + math.sin(radians) * (self.y - ship_y)
        new_y = ship_y + -math.sin(radians) * (self.x - ship_x) + math.cos(radians) * (self.y - ship_y)
        self.x = new_x
        self.y = new_y

    def calc_movement(self, nav_action, ship_x, ship_y):
        action = nav_action[0]
        value = nav_action[1]

        if action == Compass.N:
            self.y += value
        elif action == Compass.S:
            self.y -= value
        elif action == Compass.E:
            self.x += value
        elif action == Compass.W:
            self.x -= value
        elif action == Basic.L:
            self._calc_rotation(value * -1, 0, 0)
        elif action == Basic.R:
            self._calc_rotation(value, 0, 0)


class Ship():
    def __init__(self):
        self.waypoint = Waypoint()
        self.x = 0
        self.y = 0

    def calc_movement(self, nav_action):
        action = nav_action[0]
        value = nav_action[1]

        if action == Basic.F:
            self.x += self.waypoint.x * value
            self.y += self.waypoint.y * value
        else:
            self.waypoint.calc_movement(nav_action, self.x, self.y)


def find_manhattan_distance():
    with open("actions.txt", "r", encoding="utf-8") as input_file:
        nav_actions = list(map(lambda line: (action_to_enum[line[0]], int(line[1:])), input_file.readlines()))

    ship = Ship()
    for nav_action in nav_actions:
        ship.calc_movement(nav_action)
    return int(abs(ship.x) + abs(ship.y))

print(find_manhattan_distance())
