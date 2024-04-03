#include <stdio.h>

int is_prime(long num) {
    if (num <= 1) {
        return 0;
    }
    for (long i = 2; i <= num / 2; i++) {
        if (num % i == 0) {
            return 0;
        }
    }
    return 1;
}

int main() {
    long sum = 0;
    for (long i = 2; i < 2000000; i++) {
        if (is_prime(i)) {
            sum += i;
        }
    }
    printf("The sum of all prime numbers below 2 million is %ld.\n", sum);
    return 0;
}
