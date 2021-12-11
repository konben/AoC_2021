# AoC Day 2: Dive!

# Part one

course = open('course')

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
        if direction == 'forward':
            distance += int(steps)
        elif direction == 'up':
            depth -= int(steps)
        elif direction == 'down':
            depth += int(steps)
        else:
            # Bad input!
            raise(Exception('This shouldnt have happened!'))

    return (distance, depth)

# Main
dist, depth = moveSub(course)
print(f"The submarine moved {dist} steps forward and {depth} steps down!")
print(f"The product of the dept and distance traveled is {dist*depth}!")
