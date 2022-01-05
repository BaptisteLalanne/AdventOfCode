#include <cmath>
#include <fstream>
#include <iostream>
#include <vector>
#include <algorithm>
#include <numeric>
/* Advent of Code 2021 Day 07  */

int main(int argc, char *argv[]) {
    std::string input = "./inputs/day7.txt";
    std::string line;
    std::fstream file(input);

    std::getline(file, line);
    std::size_t start = 0;
    std::size_t end = line.find(',', start);

    std::vector<int> crabPosition;

    int sum = 0;
    while (end != std::string::npos) {
        int position = std::stoi(line.substr(start, end - start));
        sum += position;
        crabPosition.emplace_back(position);
        start = end + 1;
        end = line.find(',', start);
    }
    crabPosition.emplace_back(std::stoi(line.substr(start, line.size() - start)));
    sum += std::stoi(line.substr(start, line.size() - start));

    std::sort(crabPosition.begin(),crabPosition.end());


    int median;
    if (crabPosition.size() % 2 == 1) {
        median = crabPosition[(crabPosition.size() + 1) / 2];
    } else {
        int left = crabPosition[crabPosition.size() / 2 - 1];
        int right = crabPosition[crabPosition.size() / 2];
        median = (int)((left + right) / 2);
    }
 
    // Part 1
    int fuel = 0;
    for (const auto& position : crabPosition)
    {
        fuel += (int)abs(position - median);
    }

    // Part 2
    fuel = 0;
    int mean = (int) round(sum / crabPosition.size());
    for (const auto& position: crabPosition) {
        int distance = (int) abs(position - mean); 
        fuel += (distance * (distance+1))/2;
    }
    
    std::cout<<fuel<<std::endl;

    return 0;

}