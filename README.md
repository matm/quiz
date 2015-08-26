# quiz


Q: Given a list of words like https://github.com/NodePrime/quiz/blob/master/word.list find the longest compound-word in the list, which is also a concatenation of other sub-words that exist in the list. The program should allow the user to input different data. The finished solution shouldn't take more than one hour. Any programming language can be used, but Go is preferred.


Fork this repo, add your solution and documentation on how to compile and run your solution, and then issue a Pull Request. 

Obviously, we are looking for a fresh solution, not based on others' code.

A: Here is a Go implementation using a prefix tree data structure, suitable for inserting words from a sorted list and searching prefixes.

# Installation

First, ensure you have a working Go installation. Then install the program with
```
$ go get -u github.com/matm/quiz
```

# Usage

Run the program against a sorted list of words as first argument:
```
$ quiz word.list
```
