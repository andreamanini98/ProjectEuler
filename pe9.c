#include <stdio.h>
#include <stdbool.h>

#define TOTAL 1000

bool isPythagoreanTriple(int a, int b, int c);


int main() {
    for (int a = 0; a <= 500; a++) {
        for (int b = 0; b <= 500; b++) {
            for (int c = 0; c <= 500; c++) {
                if (isPythagoreanTriple(a, b, c) && (a + b + c == TOTAL)) {
                    printf("a: %d\n", a);
                    printf("b: %d\n", b);
                    printf("c: %d\n", c);
                    printf("The product is: %d", a * b * c);
                    break;
                }
            }
        }
    }
}

bool isPythagoreanTriple(int a, int b, int c) {
    if (!(a < b && b < c))
        return false;
    if (a * a + b * b == c * c)
        return true;
    return false;
}