#include <cmath>
#include <fstream>
#include <iostream>
#include <vector>
#include <algorithm>

/* Advent of Code 2021 Day 01  */

int main(int argc, char * argv[]) {
    std::string input = "./inputs/day1.txt";
    std::vector<int> depth;
    std::string line;
    std::fstream file(input);
    int value;
    
    while(file >> value){
        depth.push_back(value);
    }

    // Part 1
    int sum = 0;
    for (int i = 1; i < depth.size(); i++) {
        if (depth[i] > depth[i - 1]) {
            sum++;
        }
    }
    std::cout << "Part 1: " << sum << "\n";

    // Part 2
    sum = 0;
    for (int i = 3; i < depth.size(); i++) {
        if (depth[i] > depth[i - 3]) {
            sum++;
        }
    }
    std::cout << "Part 2: " << sum << "\n";
}



