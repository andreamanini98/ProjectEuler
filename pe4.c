#include <stdio.h>
#include <stdlib.h>
#include <math.h>
#include <stdbool.h>

void getNumOfDigits(int n, int *length);

void splitNumber(int n, int *digitsArray, int length);

bool isPalindrome(int *digitsArray, int length);


int main() {
    int *digitLength = NULL;
    int *digitsArray = NULL;
    digitLength = malloc(sizeof(int));

    int maxPalProd = 0;
    for (int i = 1; i <= 999; i++) {
        for (int j = 1; j <= 999; j++) {
            int prod = i * j;
            printf("%d\n", prod);

            getNumOfDigits(prod, digitLength);

            digitsArray = malloc((*digitLength) * sizeof(int));

            splitNumber(prod, digitsArray, *digitLength);

            if (isPalindrome(digitsArray, *digitLength) && maxPalProd < prod)
                maxPalProd = prod;

            free(digitsArray);
        }
    }

    printf("The maximum palindrome product is: %d", maxPalProd);
}

void getNumOfDigits(int n, int *length) {
    int tester = 9;
    int counter = 1;
    int i = 1;
    while (tester / n == 0) {
        tester += 9 * (int) pow((double) 10, (double) i);
        counter++;
        i++;
    }
    *length = counter;
}

void splitNumber(int n, int *digitsArray, int length) {
    int tmp;
    for (int i = 0; i < length; i++) {
        tmp = n % (int) pow((double) 10, (double) (i + 1));
        digitsArray[i] = (i == 0) ? tmp : (tmp - n % (int) pow((double) 10, (double) i)) / (int) pow((double) 10, (double) i);
    }
}

bool isPalindrome(int *digitsArray, int length) {
    for (int i = 0; i < length / 2; i++)
        if (digitsArray[i] != digitsArray[length - 1 - i])
            return false;
    return true;
}

