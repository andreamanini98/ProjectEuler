#include <stdio.h>

#define THRESHOLD 1000000

int generateCollatz(long n);


int main() {
    long number = 1;
    int longest = 0;

    for (int i = 1; i < THRESHOLD; i++) {
        //printf("%d\n", i);
        int tmp = generateCollatz(i);
        if (longest < tmp) {
            number = i;
            longest = tmp;
        }
    }

    printf("Number: %ld", number);
}

// Returns the length of the Collatz sequence.
int generateCollatz(long n) {
//    int counter = 1;
//    while (n != 1) {
//        if (n % 2 == 0)
//            n = n / 2;
//        else
//            n = 3 * n + 1;
//        counter++;
//    }
//    return counter;
    if (n == 1)
        return 1;
    return 1 + ((n % 2 == 0) ? generateCollatz(n / 2) : generateCollatz(3 * n + 1));
}