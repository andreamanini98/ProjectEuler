#include <stdio.h>
#include <stdbool.h>


bool hasOver500Divisors(long num);


int main() {
    bool foundSequence = false;
    long sum = 0;
    long index = 1;

    while(!foundSequence) {
        printf("index: %ld\n", index);
        sum += index;
        index++;
        foundSequence = hasOver500Divisors(sum);
    }

    printf("The largest is: %ld ", sum);
}

bool hasOver500Divisors(long num) {
    long counter = 0;
    long i = 0;

    for (; i <= num; i++)
        if (num % i == 0)
            counter++;

    return (counter > 500);
}

