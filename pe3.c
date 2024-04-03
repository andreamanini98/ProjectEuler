#include <stdio.h>
#include <stdbool.h>

#define NUMBER 600851475143

bool isPrime(long n);


int main() {
    for (long i = 5; i*i < NUMBER; i += 6) {
        if (isPrime(i) && NUMBER % i == 0) {
            printf("The largest prime is: %ld\n", i);
        }
    }
}

bool isPrime(long n) {
    if (n % 2 == 0 || n % 3 == 0)
        return false;

    long i = 2;
    while (i*i <= n) {
        if (n % i == 0)
            return false;
        i++;
    }
    return true;
}