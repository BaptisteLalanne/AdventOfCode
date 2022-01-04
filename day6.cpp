#include <cmath>
#include <fstream>
#include <iostream>
#include <vector>
#include <algorithm>
#include <numeric>
/* Advent of Code 2021 Day 06  */

int main(int argc, char *argv[]) {
    std::string input = "./inputs/day6.txt";
    std::string line;
    std::fstream file(input);


    std::unordered_map<std::string ,long> grid;
    std::getline(file, line);
    long lanternFish[9];
    std::size_t start = 0;
    std::size_t end = line.find(',', start);
    while (end != std::string::npos) {
        lanternFish[(std::stoi(line.substr(start, end - start)))]++;
        start = end + 1;
        end = line.find(',', start);
    }
    lanternFish[(std::stoi(line.substr(start, line.size() - start)))]++;

    for (long j = 0; j < 256; j++) {
        long newBorn = lanternFish[0];
        for(long i = 0; i < 8 ; i++) {
            lanternFish[i] = lanternFish[i+1];
        }
        lanternFish[8] = newBorn;
        lanternFish[6] += newBorn;
    }

    long n = 0;
    for(long i : lanternFish) {
        n+= i;
    }
    std::cout<<n<<std::endl;

    return 0;

}