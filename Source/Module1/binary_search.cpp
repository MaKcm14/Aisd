#include <iostream>
#include <sstream>
#include <string>
#include <vector>


int32_t BinarySearch(auto itBegin, auto itEnd, auto elem, const auto beginRangeIt) {
	int32_t size = itEnd - itBegin;

	if (size <= 1) {
		if (size && elem == *itBegin) {
			while (itBegin != beginRangeIt && elem == *itBegin) {
				--itBegin;
			}

			if (*itBegin != elem) {
				++itBegin;
			}

			return itBegin - beginRangeIt;
		}
		return -1;
	}
	
	if (*(itBegin + size / 2) <= elem) {
		return BinarySearch(itBegin + size / 2, itEnd, elem, beginRangeIt);
	} else {
		return BinarySearch(itBegin, itBegin + size / 2, elem, beginRangeIt);
	}
}

int32_t StringToNum(const std::string& num) {
	int32_t resNum;
	std::istringstream strOut(num);

	strOut >> resNum;

	return resNum;
}


int main() {
	bool flagEofInput = false;
	std::vector<int32_t> vec;
	std::vector<int32_t> findNums;

	for (std::string interm; std::cin >> interm; ) {
		if (interm != "search" && !flagEofInput && !interm.empty()) {
			vec.push_back(StringToNum(interm));
		} else if (interm == "search") {
			flagEofInput = true;
		} else {
			if (!interm.empty()) {
				findNums.push_back(StringToNum(interm));
			} else if (vec.empty()) {
				findNums.push_back(0);
			} else {
				findNums.push_back(vec.back() + 1);
			}
		}
	}

	for (const auto& elem : findNums) {
		std::cout << BinarySearch(vec.begin(), vec.end(), elem, vec.cbegin()) << std::endl;
	}

	return 0;
}