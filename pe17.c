#include <stdio.h>
#include <string.h>

char one[4] = "one\0";
char two[4] = "two\0";
char three[6] = "three\0";
char four[5] = "four\0";
char five[5] = "five\0";
char six[4] = "six\0";
char seven[6] = "seven\0";
char eight[6] = "eight\0";
char nine[5] = "nine\0";
char ten[4] = "ten\0";
char eleven[7] = "eleven\0";
char twelve[7] = "twelve\0";
char thirteen[9] = "thirteen\0";
char fourteen[9] = "fourteen\0";
char fifteen[8] = "fifteen\0";
char sixteen[8] = "sixteen\0";
char seventeen[10] = "seventeen\0";
char eighteen[9] = "eighteen\0";
char nineteen[9] = "nineteen\0";
char twenty[7] = "twenty\0";
char thirty[7] = "thirty\0";
char forty[6] = "forty\0";
char fifty[6] = "fifty\0";
char sixty[6] = "sixty\0";
char seventy[8] = "seventy\0";
char eighty[7] = "eighty\0";
char ninety[7] = "ninety\0";
char hundred[8] = "hundred\0";
char thousand[9] = "thousand\0";
char and[4] = "and\0";

unsigned int getSumFromOneToNine();

unsigned int getSumFromOneToNineteen();

unsigned int getSumFromTwentyToNinetyNine();

unsigned int getSumFromOneHundredToNineHundredAndNinetyNine(char strNumber[]);


int main() {
    unsigned int length = 0;
    length += getSumFromOneToNineteen();
    length += getSumFromTwentyToNinetyNine();
    length += getSumFromOneHundredToNineHundredAndNinetyNine(one);
    length += getSumFromOneHundredToNineHundredAndNinetyNine(two);
    length += getSumFromOneHundredToNineHundredAndNinetyNine(three);
    length += getSumFromOneHundredToNineHundredAndNinetyNine(four);
    length += getSumFromOneHundredToNineHundredAndNinetyNine(five);
    length += getSumFromOneHundredToNineHundredAndNinetyNine(six);
    length += getSumFromOneHundredToNineHundredAndNinetyNine(seven);
    length += getSumFromOneHundredToNineHundredAndNinetyNine(eight);
    length += getSumFromOneHundredToNineHundredAndNinetyNine(nine);
    length += (strlen(one) + strlen(thousand));

    printf("%u", length);
}

unsigned int getSumFromOneToNine() {
    return strlen(one) +
           strlen(two) +
           strlen(three) +
           strlen(four) +
           strlen(five) +
           strlen(six) +
           strlen(seven) +
           strlen(eight) +
           strlen(nine);
}

unsigned int getSumFromOneToNineteen() {
    return getSumFromOneToNine() +
           strlen(ten) +
           strlen(eleven) +
           strlen(twelve) +
           strlen(thirteen) +
           strlen(fourteen) +
           strlen(fifteen) +
           strlen(sixteen) +
           strlen(seventeen) +
           strlen(eighteen) +
           strlen(nineteen);
}

unsigned int getSumFromTwentyToNinetyNine() {
    return strlen(twenty) * 10 +
           strlen(thirty) * 10 +
           strlen(forty) * 10 +
           strlen(fifty) * 10 +
           strlen(sixty) * 10 +
           strlen(seventy) * 10 +
           strlen(eighty) * 10 +
           strlen(ninety) * 10 +
           getSumFromOneToNine() * 8;
}

unsigned int getSumFromOneHundredToNineHundredAndNinetyNine(char strNumber[]) {
    return (strlen(strNumber) + strlen(hundred)) * 100 +
           strlen(and) * 99 +
           getSumFromOneToNineteen() + getSumFromTwentyToNinetyNine();
}