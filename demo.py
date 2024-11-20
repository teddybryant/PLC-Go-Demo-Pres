import time
from multiprocessing import Pool

# Function to calculate the square of numbers in a chunk
def calculate_squares(numbers):
    return sum(x * x for x in numbers)

# Parallel sum using multiprocessing
def parallel_sum(numbers):
    chunk_size = len(numbers) // 4
    chunks = [numbers[i * chunk_size : (i + 1) * chunk_size] for i in range(4)]

    with Pool(4) as pool:
        results = pool.map(calculate_squares, chunks)

    return sum(results)

# Read integers from file
def read_integers_from_file(filename):
    numbers = []
    with open(filename, "r") as f:
        for line in f:
            numbers.extend(map(int, line.split()))
    return numbers

if __name__ == "__main__":
    filename = "numbers.txt"  # Create a file with integers separated by spaces/newlines

    # Read numbers from file
    numbers = read_integers_from_file(filename)

    # Measure time for parallel computation
    start = time.time()
    result = parallel_sum(numbers)
    elapsed = time.time() - start

    print(f"Parallel sum of squares: {result}")
    print(f"Time taken: {elapsed:.6f} seconds")
