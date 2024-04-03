#include <stdio.h>

int main() {
    int sum = 0;

    for (int i = 0; i < 1000; i++)
        (i % 3 == 0 || i % 5 == 0) ? sum += i : sum;

    printf("The sum is: %d", sum);
}
