# Generating a large "numbers.txt" file with random integers.
import random

# Parameters for the file
file_name = "./numbers.txt"  # Path to save the file
num_numbers = 10**6  # Total numbers
max_value = 10**6  # Maximum value of integers

# Generate random integers and save to file
with open(file_name, "w") as file:
    for _ in range(num_numbers):
        # Write numbers separated by spaces in chunks of 10 per line
        numbers_line = " ".join(str(random.randint(1, max_value)) for _ in range(10))
        file.write(numbers_line + "\n")

file_name