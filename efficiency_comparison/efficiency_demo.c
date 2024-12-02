#include <stdio.h>
#include <stdlib.h>
#include <time.h>

long long sumFibonacci(int n) {
    if (n <= 0) {
        return 0;
    }
    long long a = 0, b = 1, sum = 1;
    for (int i = 2; i < n; i++) {
        long long next = a + b;
        sum += next;
        a = b;
        b = next;
    }
    return sum;
}

int main(int argc, char *argv[]) {
    long n = 1000000;
    clock_t start = clock();
    long long sum = sumFibonacci(n);

    clock_t end = clock();
    double elapsed_time = ((double)(end - start)) / CLOCKS_PER_SEC;

    printf("Sum of the first %ld Fibonacci numbers: %lld\n", n, sum);
    printf("C's Execution time: %f seconds\n", elapsed_time);

    return 0;
}
