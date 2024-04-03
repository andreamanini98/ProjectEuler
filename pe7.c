#include <stdio.h>
#include <stdbool.h>

bool isPrime(long n);


int main() {
    int counter = 1;
    long currentNumber = 1;

    while (counter < 10001) {
        if (isPrime(currentNumber))
            counter++;
        if (counter < 10001)
            currentNumber++;
    }

    printf("The 10001st prime number is: %ld", currentNumber);
}

bool isPrime(long n) {
    if (n % 2 == 0 || n % 3 == 0)
        return false;

    long i = 2;
    while (i * i <= n) {
        if (n % i == 0)
            return false;
        i++;
    }
    return true;
}
