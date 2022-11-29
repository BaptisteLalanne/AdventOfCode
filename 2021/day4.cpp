#include <cmath>
#include <fstream>
#include <iostream>
#include <vector>
#include <algorithm>
#include <numeric>
/* Advent of Code 2021 Day 04  */

std::vector<int> splitBingoInputs(const std::string& s);
std::vector<int> checkBingoWin(const std::string& s);
void updateBingo(std::vector<std::vector<std::vector<int>>>& listeBingo, const int number);
std::tuple<int,int> checkWinBingo(std::vector<std::vector<std::vector<int>>>& listeBingo);

int main(int argc, char *argv[]) {
    std::string input = "./inputs/day4.txt";
    std::string line;
    std::fstream file(input);


    const int bingoSize = 5;

    std::getline(file, line);
    const std::vector<int> randomNumbers = splitBingoInputs(line);
    std::vector < std::vector < std::vector < int>>> bingoList;
    while (std::getline(file, line)) {
        if (line != "/n") {
            // Create new board
            bingoList.emplace_back();
            for (int i = 0; i < bingoSize; i++) {
                // Create new line
                bingoList.back().emplace_back();
                for (int j = 0; j < bingoSize; j++) {
                    int n;
                    file >> n;
                    // Add number to line
                    bingoList.back().back().emplace_back(n);
                }
            }
        }
    }

    int nbWinner = 0;
    int lastWinningScore;
    for (const int &number: randomNumbers) {
        updateBingo(bingoList, number);
        std::tuple<int,int> winner = checkWinBingo(bingoList);
        if (std::get<0>(winner) != -1) {
            nbWinner++;
            lastWinningScore = std::get<1>(winner) * number;
            // Part 1
            if (nbWinner == 1) {
                std::cout << lastWinningScore << std::endl;
            }
        }
    }
    // Part 2
    std::cout << lastWinningScore << std::endl;


    return 0;
}

std::vector<int> splitBingoInputs(const std::string &s) {
    std::size_t start = 0;
    std::size_t end = s.find(',');
    std::vector<int> v;
    while (end != std::string::npos) {
        v.emplace_back(std::stoi(s.substr(start, end - start)));
        start = end + 1;
        end = s.find(',', start);
    }
    v.emplace_back(std::stoi(s.substr(start, s.size() - start)));
    return v;
}

std::tuple<int,int> checkWinBingo(std::vector<std::vector<std::vector<int>>>& listeBingo){
    std::tuple<int,int> out = {-1,-1};

    // For each bingo grid
    int gridPosition = -1;
    for (std::vector<std::vector<int>>& bingo : listeBingo) {
        gridPosition ++;
        for(int i = 0; i < bingo.size(); i++) {
            bool horizontalAlive = true;
            bool verticalAlive = true;
            for(int j = 0; j < bingo[0].size(); j++) {
                if(bingo[i][j] != -1){
                    horizontalAlive = false;
                }
                if(bingo[j][i] != -1) {
                    verticalAlive = false;
                }
                if(verticalAlive || horizontalAlive) {
                    if(j == bingo[0].size() -1){
                        int sum = 0;
                        for(std::vector<int>& line: bingo) {
                            sum += std::accumulate(std::begin(line), std::end(line), 0, [](const auto sum, const auto ele) { return ele != -1 ? sum + ele : sum; });
                        }
                        // Remove winning grid
                        listeBingo.erase(listeBingo.begin()+gridPosition);
                        out = {gridPosition,sum};
                        break;
                    }
                } else {
                    break;
                }
            }
        }
    }

    return out;
}


void updateBingo(std::vector<std::vector<std::vector<int>>>& listeBingo, const int number) {
    for (std::vector<std::vector<int>>& bingo : listeBingo) {
        for(std::vector<int>& line: bingo ) {
            if(const auto it = std::find(std::begin(line), std::end(line), number); it != line.end()){
                *it = -1;
                // Change bingo grid if number already found
                break;
            }
        }
    }
}