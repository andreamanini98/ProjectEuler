#include <stdio.h>

#define GRID_DIM (20 + 1)

unsigned long long getNumOfPaths(int dim) {
    unsigned long long grid[dim][dim];

    // Initialize the grid with zeros
    for (int i = 0; i < dim; i++) {
        for (int j = 0; j < dim; j++) {
            grid[i][j] = 0;
        }
    }

    // Set the base cases
    grid[0][0] = 1;

    // Fill the grid using dynamic programming
    for (int i = 0; i < dim; i++) {
        for (int j = 0; j < dim; j++) {
            if (i > 0) {
                grid[i][j] += grid[i - 1][j]; // Move down
            }
            if (j > 0) {
                grid[i][j] += grid[i][j - 1]; // Move right
            }
        }
    }

    //for (int i = 0; i < GRID_DIM; i++) {
    //    for (int j = 0; j < GRID_DIM; j++) {
    //        printf("%llu ", grid[i][j]);
    //    }
    //    printf("\n");
    //}

    return grid[dim - 1][dim - 1];
}

int main() {
    printf("%llu\n", getNumOfPaths(GRID_DIM));
    return 0;
}
