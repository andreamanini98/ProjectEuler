#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define TOT_NUMBERS 100         // The total numbers we have to sum
#define NUM_SIZE 50             // The total number of digits each number has
#define GLOBAL_REMAINDER_SIZE 4 // Put this big enough to contain the total remainder

int globalRemainder = 0; // This will account for the total remainder after normal sums

void getNumbers(char *numbers[TOT_NUMBERS]);

char *sumTwoIntArrays(const char *a1, const char *a2, int size);


int main() {
    // Initialize the array for containing all numbers
    char *numbers[TOT_NUMBERS];
    for (int i = 0; i < TOT_NUMBERS; i++)
        numbers[i] = malloc(NUM_SIZE + 1 * sizeof(char));

    // Get the numbers
    getNumbers(numbers);

    // Perform a cyclic sum on every couple of numbers to sum all of them
    char *middleComp = numbers[0];
    for (int i = 1; i < TOT_NUMBERS; i++)
        middleComp = sumTwoIntArrays(middleComp, numbers[i], NUM_SIZE);

    // Get the string representation of the global remainder to concatenate it at the end
    char *globalRemainderString = malloc(GLOBAL_REMAINDER_SIZE * sizeof(char));
    sprintf(globalRemainderString, "%d", globalRemainder);

    printf("The sum is: %s\n", middleComp);
    printf("The total remainder is: %d\n", globalRemainder);
    printf("The total number is thus: %s", strcat(globalRemainderString, middleComp));

    for (int i = 0; i < TOT_NUMBERS; i++)
        free(numbers[i]);
    free(globalRemainderString);
}

// Simply a function to load numbers from a .txt file into an array.
// Please write each number on a newline.
void getNumbers(char *numbers[TOT_NUMBERS]) {
    FILE *fp;
    char *line = NULL;
    size_t len = 0;
    int i = 0;

    fp = fopen("/Users/echo/Desktop/misc/ProjectEuler/files/pe13.txt", "r");
    if (fp == NULL)
        exit(EXIT_FAILURE);

    while (getline(&line, &len, fp) != -1) {
        line[NUM_SIZE] = '\0';
        strcpy(numbers[i], line);
        i++;
    }

    fclose(fp);
}

// Function to sum up two arrays of length "size".
char *sumTwoIntArrays(const char *a1, const char *a2, int size) {
    char *result = malloc(size * sizeof(char));
    int tmpRemainder = 0;

    for (int i = size - 1; i >= 0; i--) {
        int tmp = ((int) a1[i] - '0') + ((int) a2[i] - '0') + tmpRemainder;
        // If we have a remainder we have to take it into account
        if (tmp > 9) {
            // If i == 0 we have to update the globalRemainder instead of the local one
            (i > 0) ? (tmpRemainder = tmp / 10) : (globalRemainder += tmp / 10);
            tmp = tmp % 10;
        } else {
            tmpRemainder = 0;
        }
        result[i] = (char) tmp + '0';
    }
    return result;
}
