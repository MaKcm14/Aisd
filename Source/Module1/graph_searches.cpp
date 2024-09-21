#include <algorithm>
#include <iostream>
#include <queue>
#include <sstream>
#include <unordered_map>
#include <unordered_set>
#include <vector>

std::vector<std::string> SplitToSpaces(const std::string& str) {
    std::vector<std::string> strings;
    std::vector<std::string> result;

    for (size_t last = 0, newLast = 0; last < str.size(); last = newLast) {
        if (!last) {
            last = str.find(" ");
            strings.push_back(str.substr(0, last));
        }
        newLast = str.find(" ", last + 1);
        strings.push_back(str.substr(last + 1, newLast - last - 1));
    }

    result.reserve(strings.size());

    for (const auto& elem : strings) {
        if (!elem.empty()) {
            result.push_back(elem);
        }
    }


    return result;
}

std::unordered_map<std::string, std::vector<std::string>> DeserealizeInput(const std::vector<std::string>& input, const std::string& graphType) {
    std::unordered_map<std::string, std::vector<std::string>> graph;

    for (size_t i = 1; i != input.size(); ++i) {
        auto vertex = SplitToSpaces(input[i]);
        graph[vertex[0]].push_back(vertex[1]);

        if (graphType.find("u") < graphType.size()) {
            graph[vertex[1]].push_back(vertex[0]);
        }
    }

    for (auto& elem : graph) {
        std::sort(elem.second.begin(), elem.second.end());
    }

    return graph;
}

void DFS(std::unordered_map<std::string, std::vector<std::string>>& graph, std::unordered_set<std::string>& visited, std::vector<std::string>& visitedOrder, std::string v) {
	visited.insert(v);
    visitedOrder.push_back(v);
    for (size_t i = 0; i != graph[v].size(); ++i) {
        if (auto itFind = visited.find(graph[v][i]); itFind == visited.end()) {
            DFS(graph, visited, visitedOrder, graph[v][i]);
        }
    } 
}

void BFS(std::unordered_map<std::string, std::vector<std::string>>& graph, std::unordered_set<std::string>& visited, std::vector<std::string>& visitedOrder, std::string v) {
    std::queue<std::string> intermQueue;
	intermQueue.push(v);
	visited.insert(v);
    visitedOrder.push_back(v);

	while (!intermQueue.empty()) {
		auto val = intermQueue.front();
		intermQueue.pop();

		for (size_t i = 0; i != graph[val].size(); ++i) {
			if (auto itFind = visited.find(graph[val][i]); itFind == visited.end()) {
				intermQueue.push(graph[val][i]);
				visited.insert(graph[val][i]);
                visitedOrder.push_back(graph[val][i]);
			}
		}
	}
}

int main() {
    std::vector<std::string> input;

    for (std::string interm; std::getline(std::cin, interm); ) {
        if (!interm.empty()) {
            input.push_back(interm);
        }
    }

    auto outputType = SplitToSpaces(input[0]);
    auto graph = DeserealizeInput(input, outputType[0]);
    std::unordered_set<std::string> visited;
    std::vector<std::string> visitedOrder;

    if (outputType[2].find("d") < outputType[2].size()) {
        DFS(graph, visited, visitedOrder, outputType[1]);
    } else {
        BFS(graph, visited, visitedOrder, outputType[1]);
    }

    for (const auto& vertex : visitedOrder) {
        std::cout << vertex << std::endl;
    }
    
    return 0;
}