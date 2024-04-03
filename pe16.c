#include <stdio.h>
#include <stdlib.h>

#define TOT_DIGITS 302 // we will represent the decimal number corresponding to 2^EXPONENT using TOT_DIGITS digits
#define EXPONENT 1000

char *doubleCharNumber(const char *number, int size);

int sumDigits(const char *number, int size);


int main() {
    char *number = malloc(TOT_DIGITS * sizeof(char));
    for (int i = 0; i < TOT_DIGITS; i++)
        number[i] = '0';
    number[TOT_DIGITS - 1] = '1';

    for (int i = 0; i < EXPONENT; i++)
        number = doubleCharNumber(number, TOT_DIGITS);

    printf("%s\n\n", number);
    printf("%d", sumDigits(number, TOT_DIGITS));

    free(number);
}

char *doubleCharNumber(const char *number, int size) {
    char *result = malloc(size * sizeof(char));
    int remainder = 0;
    for (int i = size - 1; i >= 0; i--) {
        int tmp = ((int) number[i] - '0') * 2 + remainder;
        if (tmp > 9) {
            remainder = tmp / 10;
            tmp = tmp % 10;
        } else
            remainder = 0;
        result[i] = (char) tmp + '0';
    }
    return result;
}

int sumDigits(const char *number, int size) {
    int result = 0;
    for (int i = 0; i < size; i++)
        result += ((int) number[i] - '0');
    return result;
}
