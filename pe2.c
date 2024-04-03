#include <stdio.h>
#include <stdbool.h>

#define THRESHOLD 4000000

int main() {
    bool stop = false;
    int sum = 0;
    int prev = 0;
    int cur = 1;
    int temp;

    while (!stop) {
        temp = cur;
        cur += prev;
        if (cur > THRESHOLD) {
            stop = true;
        } else {
            prev = temp;
            if (cur % 2 == 0)
                sum += cur;
        }
    }

    printf("The sum is: %d", sum);
}
