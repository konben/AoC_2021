# AoC Day 2: Dive!

# Part one

def moveSub(course) -> tuple:
    """
    Computes the position of the submarine after following a course.

    Arguments:
        course -- File containing directions of the form <direction> <steps>.
    Returns:
        The position after follonwing the course as a tuple (distance, depth).
    """

    # Initial position
    distance, depth = 0, 0

    # Reading directions
    for line in course:
        direction, steps = line.split()
        steps = int(steps)
        if direction == 'forward':
            distance += steps
        elif direction == 'up':
            depth -= steps
        elif direction == 'down':
            depth += steps
        else:
            # Bad input!
            raise(Exception('This shouldnt have happened!'))

    return (distance, depth)

# Main
course = open('course')
dist, depth = moveSub(course)
print('Part one:')
print(f"The submarine moved {dist} steps forward and {depth} steps down!")
print(f"The product of the dept and distance traveled is {dist*depth}!")

# Part two

def moveSub2(course) -> tuple:
    """
    Computes the position of the submarine after following a course.

    Arguments:
        course -- File containing directions of the form <direction> <steps>.
    Returns:
        The position after follonwing the course as a tuple (distance, depth).
    """

    # Initial position
    distance, depth, aim = 0, 0, 0

    # Reading directions
    for line in course:
        direction, steps = line.split()
        steps = int(steps)
        if direction == 'forward':
            distance += steps
            depth += steps*aim
        elif direction == 'up':
            aim -= steps
        elif direction == 'down':
            aim += steps
        else:
            # Bad input!
            raise(Exception('This shouldnt have happened!'))

    return (distance, depth)

# Main
course = open('course')
dist, depth = moveSub2(course)
print('Part two:')
print(f"The submarine moved {dist} steps forward and {depth} steps down!")
print(f"The product of the dept and distance traveled is {dist*depth}!")
