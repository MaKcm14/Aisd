#include <exception>
#include <iostream>
#include <sstream>
#include <vector>

class TDeque {
public:
    TDeque() = default;

    TDeque(size_t size) {
        Arr.reserve(size);
    }

    TDeque(const TDeque& deque) {
        Arr = std::move(deque.Arr);
        Arr.reserve(deque.GetCapacity());
    }


    size_t GetCapacity() const noexcept {
        return Arr.capacity();
    }

    size_t GetSize() const noexcept {
        return Arr.size();
    }

    void Reserve(size_t size) {
        Arr.reserve(size);
    }

    void Print() const noexcept {
        for (const auto& elem : Arr) {
            std::cout << elem << " ";
        }

        if (Arr.empty()) {
            std::cout << "empty";
        }
        std::cout << std::endl;
    }

    void PushBack(const std::string& elem) noexcept {
        if (Arr.size() == Arr.capacity()) {
            std::cout << "overflow" << std::endl;
        } else {
            Arr.push_back(elem);
        }
    }

    void PushFront(const std::string& elem) noexcept {
        if (Arr.size() == Arr.capacity()) {
            std::cout << "overflow" << std::endl;
        } else {
            Arr.insert(Arr.begin(), elem);
        }
    }

    void PopBack() noexcept {
        if (Arr.empty()) {
            std::cout << "underflow" << std::endl;
        } else {
            std::cout << Arr.back() << std::endl;
            Arr.pop_back();
        }
    }

    void PopFront() noexcept {
        if (Arr.empty()) {
            std::cout << "underflow" << std::endl;
        } else {
            std::cout << Arr.front() << std::endl;
            Arr.erase(Arr.begin());
        }
    }


private:
    std::vector<std::string> Arr;

};


int32_t ConvertStringToNum(const std::string& str) {
    if (str.empty()) {
        return 0;
    }

    int32_t res = 0;
    std::istringstream strIn(str);

    strIn >> res;

    return res;
}

std::vector<std::string> SplitToSpaces(const std::string& str) {
    std::vector<std::string> strings;

    for (size_t last = 0, newLast = 0; last < str.size(); last = newLast) {
        if (!last) {
            last = str.find(" ");
            strings.push_back(str.substr(0, last));
        }
        newLast = str.find(" ", last + 1);
        strings.push_back(str.substr(last + 1, newLast - last));
    }

    
    if (!strings.empty() && strings.back().empty()) {
        strings.pop_back();
    }

    return strings;
}

std::pair<TDeque, std::vector<std::pair<std::string, std::string>>> ParseInput() {
    TDeque deque;
    std::vector<std::pair<std::string, std::string>> actions;
    
    bool setSizeFlag = false;

    for (std::string interm; std::getline(std::cin, interm); ) {
        auto command = SplitToSpaces(interm);

        if (command.empty()) {
            continue;
        }
        
        if (command.front() == "set_size" && !setSizeFlag) {
            if (command.size() == 2) {
                auto reserveNum = ConvertStringToNum(command[1]);
                
                if (reserveNum >= 0) {
                    deque.Reserve(ConvertStringToNum(command[1]));
                    setSizeFlag = true;
                } else {
                    actions.push_back({ "", "" });
                }
            }

        } else if (!setSizeFlag) {
            actions.push_back({ "", "" });

        } else if (command.size() == 2 && (command.front() == "pushb" || command.front() == "pushf")) {
            actions.push_back({ command.front(), command.back() });

        } else if (command.size() == 1 && (command.front() == "popb" || command.front() == "popf")) {
            if (interm.find(command.front() + " ") > interm.size()) {
                actions.push_back({ command.front(), "" });
                
            } else {
                actions.push_back({ "", "" });
            }

        } else if (command.size() == 1 && command.front() == "print") {
            if (interm.find("print ") > interm.size() && setSizeFlag) {
                actions.push_back({ "print", "" });

            } else {
                actions.push_back({ "", "" });
            }

        } else {
            actions.push_back({ "", "" });
        }
    }

    return { deque, actions };
}

int main() {
    
    auto [deque, actions] = ParseInput();

    for (const auto& [action, elem] : actions) {
        if (action == "pushb") {
            deque.PushBack(elem);

        } else if (action == "pushf") {
            deque.PushFront(elem);

        } else if (action == "popb") {
            deque.PopBack();

        } else if (action == "popf") {
            deque.PopFront();

        } else if (action == "print") {
            deque.Print();

        } else {
            std::cout << "error" << std::endl;
        }
    }

    return 0;
}