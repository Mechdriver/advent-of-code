from enum import Enum

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


class Ship():
    def __init__(self):
        self.heading = Compass.E
        self.x = 0
        self.y = 0

    def _calc_rotation(self, degrees):
        if self.heading.value + degrees < 0:
            degrees += 360

        degrees = self.heading.value + degrees
        if degrees >= 360:
            degrees -= 360

        self.heading = degrees_to_compass[degrees]

    def calc_movement(self, nav_action):
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
        elif action == Basic.F:
            self.calc_movement((self.heading, value))
        elif action == Basic.L:
            self._calc_rotation(value * -1)
        elif action == Basic.R:
            self._calc_rotation(value)


def find_manhattan_distance():
    with open("actions.txt", "r", encoding="utf-8") as input_file:
        nav_actions = list(map(lambda line: (action_to_enum[line[0]], int(line[1:])), input_file.readlines()))

    ship = Ship()

    for nav_action in nav_actions:
        ship.calc_movement(nav_action)

    print(abs(ship.x) + abs(ship.y))

find_manhattan_distance()
