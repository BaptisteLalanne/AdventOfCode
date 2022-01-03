#include <cmath>
#include <fstream>
#include <iostream>
#include <vector>
#include <algorithm>
#include <numeric>
/* Advent of Code 2021 Day 04  */

void update(const int x1,const int y1,const int x2,const int y2, std::unordered_map<std::string,int>& grid);

int main(int argc, char *argv[]) {
    std::string input = "./inputs/day5.txt";
    std::string line;
    std::fstream file(input);


    std::unordered_map<std::string ,int> grid;
    while (std::getline(file, line)) {
        int x1, x2, y1, y2;

        // X1
        std::size_t start = 0;
        std::size_t end = line.find_first_of(",");
        x1 = std::stoi(line.substr(start,end-start));
        // Y1
        start = end + 1; // +1 to get over ","
        end = line.find_first_of("->");
        y1 = std::stoi(line.substr(start,end-start));
        line = line.substr(end+3); // +4 to get over "_->_"
        // X2
        start = 0;
        end = line.find_first_of(",");
        x2 = std::stoi(line.substr(start,end-start));
        // Y2
        start = end + 1; // +1 to get over ","
        y2 = std::stoi(line.substr(start,line.size()));


        update(x1,y1,x2,y2,grid);
    }
    int n = 0;
    for (const auto& [coord, count] : grid) {
        if (count > 1) {
            n++;
        }
    }
    std::cout << n << std::endl;

    return 0;

}

void update(const int x1,const int y1,const int x2,const int y2, std::unordered_map<std::string,int>& grid){
    if(y1 == y2) {
        int diff = (x1 - x2)/abs(x1-x2);
        int start = x2;
        int end = x1;
        while(start != end){
            // ex: grid["43-234"]++;
            std::string coord = std::to_string(start)+"-"+std::to_string(y1);
            grid[coord]++;
            start += diff;
        }
        std::string coord = std::to_string(x1)+"-"+std::to_string(y1);
        grid[coord]++;
    } else if (x1 == x2) {
        int diff = (y1 - y2)/abs(y1-y2);
        int start = y2;
        int end = y1;
        while(start != end){
            std::string coord = std::to_string(x1)+"-"+std::to_string(start);
            grid[coord]++;
            start += diff;
        }
        std::string coord = std::to_string(x1)+"-"+std::to_string(y1);
        grid[coord]++;
    } else {
        int diffX = (x1 - x2)/abs(x1-x2);
        int diffY = (y1 - y2)/abs(y1-y2);
        int startX = x2;
        int startY = y2;
        int endX = x1;
        int endY = y1;
        while(startX != endX){
            std::string coord = std::to_string(startX)+"-"+std::to_string(startY);
            grid[coord]++;
            startX += diffX;
            startY += diffY;
        }
        std::string coord = std::to_string(endX)+"-"+std::to_string(endY);
        grid[coord]++;
    }
}
