#include <cmath>
#include <fstream>
#include <iostream>
#include <vector>
#include <algorithm>

int convertToBinary(std::vector<int> binVec);
int update(std::vector<bool>& fitsCriteria, const std::vector<std::vector<int>>& values, const int acceptable_val, const int index, std::vector<int>& digitSum);
/* Advent of Code 2021 Day 03  */

int main(int argc, char * argv[]) {
    std::string input = "./inputs/day3.txt";
    std::string line;
    std::fstream file(input);

    int n = 0;
    std::getline(file, line);
    n++;
    const size_t s = line.size();
    std::vector<int> digits(s);
    for (int i = 0; i < line.size(); i++) {
        digits[i] = line[i] - '0';
    }
    while(std::getline(file, line)) {
        n++;
        for (int i = 0; i < s; i++) {
            digits[i] += line[i] - '0';
        }
    }

    for (int& ele : digits) {
        ele = (ele * 2) / n;
    }

    const int gamma = convertToBinary(digits);
    for (int& digit : digits) {
        digit = std::abs(1 - digit);
    }
    const int eps = convertToBinary(digits);

    std::cout << gamma * eps << '\n';

    /**
     * Part 2
     * */
    std::fstream file2(input);
    std::vector<std::vector<int>> values;
    while(std::getline(file2, line)) {
        values.emplace_back();
        for (const auto c : line) {
            values.back().emplace_back(c - '0');
        }
    }

    const int n_values = values.size();
    const int n_digits = values[0].size();

    std::vector<int> digit_sum(n_digits);
    for (const auto& value : values) {
        for (int i = 0; i < n_digits; i++) {
            digit_sum[i] += value[i];
        }
    }

    int o2 = 0;
    int co2 = 0;
    const auto digit_sum_o = digit_sum;
    int n_fit_c = n_values;
    std::vector<bool> fits_criteria(n_values, true);
    for (int index = 0; index < n_digits; index ++) {
        if (digit_sum[index] * 2 >= n_fit_c) {
            n_fit_c -= update(fits_criteria, values, 1, index, digit_sum);
        } else {
            n_fit_c -= update(fits_criteria, values, 0, index, digit_sum);
        }
        if (n_fit_c == 1) {
            const auto o2_index = std::distance(std::begin(fits_criteria), std::find(std::begin(fits_criteria), std::end(fits_criteria), true));
            o2 = convertToBinary(values[o2_index]);
            break;
        }
    }

    n_fit_c = n_values;
    fits_criteria = std::vector<bool>(n_values, true);
    digit_sum = digit_sum_o;
    for (int index = 0; index < n_digits; index ++) {
        if (digit_sum[index] * 2 >= n_fit_c) {
            n_fit_c -= update(fits_criteria, values, 0, index, digit_sum);
        } else {
            n_fit_c -= update(fits_criteria, values, 1, index, digit_sum);
        }
        if (n_fit_c == 1) {
            const auto co2_index = std::distance(std::begin(fits_criteria), std::find(std::begin(fits_criteria), std::end(fits_criteria), true));
            co2 = convertToBinary(values[co2_index]);
            break;
        }
    }

    std::cout << o2 * co2 << '\n';


    return 0;
}

int convertToBinary(std::vector<int> binVec) {
    const int s = binVec.size();
    int binVal = 0;
    for (int i = 0; i < s; i++) {
        binVal += (binVec[s - i - 1]) << i;
    }
    return binVal;
}

int update(std::vector<bool>& fits_criteria, const std::vector<std::vector<int>>& values, const int acceptable_val, const int index, std::vector<int>& digit_sum) {
    int n_not_fit_c = 0;
    for (int i = 0; i < values.size(); i++) {
        if (values[i][index] != acceptable_val && fits_criteria[i]) {
            fits_criteria[i] = false;
            n_not_fit_c++;
            for (int j = 0; j < digit_sum.size(); j++) {
                digit_sum[j] -= values[i][j];
            }
        }
    }
    return n_not_fit_c;
}



