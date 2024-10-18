#include <exception>
#include <iostream>
#include <sstream>
#include <vector>

class TDeque {
public:
    TDeque() = default;

    TDeque(size_t size) {
        Arr.resize(size);
        Size = 0;
        HeadIndex = 0;
        TailIndex = 0;
    }

    TDeque(const TDeque& deque) {
        Arr = std::move(deque.Arr);
        HeadIndex = deque.HeadIndex;
        TailIndex = deque.TailIndex;
        Size = deque.Size;
    }


    size_t GetCapacity() const noexcept {
        return Arr.capacity();
    }

    size_t GetSize() const noexcept {
        return Arr.size();
    }

    void Print() const noexcept {
        if (!Size) {
            std::cout << "empty";
        } else {
            for (size_t i = HeadIndex; i != TailIndex; ) {
                std::cout << Arr[i] << " ";
                if (i == Arr.size() - 1) {
                    i = 0;
                } else {
                    ++i;
                }
            }
            std::cout << Arr[TailIndex] << " ";
        }
        std::cout << std::endl;
    }

    void PushBack(const std::string& elem) noexcept {
        if (Arr.size() == Size) {
            std::cout << "overflow" << std::endl;
        } else {
            if (TailIndex == HeadIndex && !Size) {
                Arr[TailIndex] = elem;
            } else {
                ++TailIndex;
                if (TailIndex == Arr.size()) {
                    TailIndex = 0;
                }
                Arr[TailIndex] = elem;
            }
            ++Size;
        }
    }

    void PushFront(const std::string& elem) noexcept {
        if (Arr.size() == Size) {
            std::cout << "overflow" << std::endl;
        } else {
            if (TailIndex == HeadIndex && !Size) {
                Arr[HeadIndex] = elem;
            } else {
                --HeadIndex;
                if (HeadIndex > Arr.size()) {
                    HeadIndex = Arr.size() - 1;
                }
                Arr[HeadIndex] = elem;
            }
            ++Size;
        }
    }

    void PopBack() noexcept {
        if (!Size) {
            std::cout << "underflow" << std::endl;
        } else {
            std::cout << Arr[TailIndex] << std::endl;
            --TailIndex;
            if (TailIndex > Arr.size()) {
                TailIndex = Arr.size() - 1;
            }
            --Size;
        }
    }

    void PopFront() noexcept {
        if (!Size) {
            std::cout << "underflow" << std::endl;
        } else {
            std::cout << Arr[HeadIndex] << std::endl;
            ++HeadIndex;
            if (HeadIndex == Arr.size()) {
                HeadIndex = 0;
            }
            --Size;
        }
    }


    TDeque& operator = (TDeque&& deque) {
        HeadIndex = deque.HeadIndex;
        TailIndex = deque.TailIndex;
        Size = deque.Size;
        Arr = std::move(deque.Arr);

        return *this;
    }


private:
    size_t HeadIndex; // индекс первого элемента
    size_t TailIndex; // индекс последнего добавленного элемента
    size_t Size;
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
                    deque = TDeque(ConvertStringToNum(command[1]));
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