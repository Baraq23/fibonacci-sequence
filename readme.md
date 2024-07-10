markdown

# Fibonacci Sequence Program

This Go program calculates the Fibonacci value at a given index. It includes input validation to ensure that the input is a numeric value and within the acceptable range.

## Getting Started

### Prerequisites

- Go installed on your machine. You can download and install Go from [the official Go website](https://golang.org/dl/).

### Installation

1. Clone the repository:

```bash
   git clone https://github.com/Baraq23/fibonacci-sequence.git
   cd fibonacci-sequence
```

### Usage

To run the program, use the following command:

```bash
go run . <int>
```

- Replace <int> with an integer value that represents the index in the Fibonacci sequence you want to calculate. The program will print the Fibonacci value at that index if the input is valid and within the acceptable range (0-45).

_Example:_

```bash
go run . 10
```

_Output:_

```bash
The Fibonacci value at index 10 is 55
```

## Functions
##### ValidateIndex(s string) bool
- This function checks if the input string s is a valid numeric string.

##### FibonacciNum(n int) int
- This function calculates the Fibonacci number at the given index n.
Tests



The project includes a test file, main_test.go, with tests for the ValidateIndex and FibonacciNum functions.
Running Tests

To run the tests, use the following command:

```bash
go test
```

The command will execute the tests and provide feedback on whether they passed or failed.
Contributing

If you wish to contribute to the project, please fork the repository and submit a pull request with your changes.
License

This project is licensed under the MIT License - see the LICENSE file for details.

### Acknowledgments

    Go Documentation
    Fibonacci Sequence