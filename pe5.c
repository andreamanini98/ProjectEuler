#include <stdio.h>
#include <stdbool.h>

bool isEvenlyDivisible(int n);


int main() {
    int i = 20;
    while (!isEvenlyDivisible(i))
        i++;
    printf("Smallest is: %d", i);
}

bool isEvenlyDivisible(int n) {
    for (int i = 1; i <= 20; i++)
        if (n % i != 0)
            return false;
    return true;
}