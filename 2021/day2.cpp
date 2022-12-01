#include <cmath>
#include <fstream>
#include <iostream>
#include <vector>
#include <algorithm>

/* Advent of Code 2021 Day 02  */

int main(int argc, char * argv[]) {
    std::string input = "./inputs/day2.txt";
    std::fstream file(input);

    int horizontal, vertical, aim, depth = 0;
    std::string command;
    int measure;
    while(file >> command >> measure){
        if(command.compare("forward") == 0){
            horizontal += measure;
            depth += measure*aim;
        }else if(command.compare("up") == 0) {
            vertical -= measure;
            aim -= measure;
        } else {
            vertical += measure;
            aim += measure;
        }
    }

    // Part 1
    std::cout << "Part 1: " << horizontal*vertical << "\n";
    // Part 2
    std::cout << "Part 2: " << depth*horizontal << "\n";
}



