# Reverse Polish Notation Calculator

This is a simple RPN calculator written in go that can be run from the command line. It supports 64-bit floats and the four primary arithmetic operators (+,-,*,/).

## Install

Make sure you have [go installed.](https://golang.org/doc/install)

`$ go get -u github.com/delputnam/rpn_calc`

This should fetch the repository and build the executable: `$GOPATH/bin/rpn_calc`.

To run the unit tests, change the current directory to `$GOPATH/rpn_calc` and then:

`$ go test -v ./...`

__Note:__ For convenience I have also added binaries for Linux, Windows and Mac OS X (darwin) to this repository. The Windows and Linux binaries were cross compiled with Go 1.6.2 on Mac OS X 10.11.4. (Only the darwin binary has been thoroughly tested.)

## Usage

Run `rpn_calc` from the command line and input a valid operand or operator at the `>` prompt and press return.

Each valid operand will be pushed onto the stack and when a valid operator is encountered, the last two operands will be popped off and the specified operation will be performed. The result is displayed on the next line.

For example, to add the integers 2 and 5, then divide the result by 2, you would enter the following:

```
> 2
2
> 5
5
> +
7
> 2
2
> /
3.5
>
```

To exit the program gracefully enter `q` or `EOF` (typically, `ctrl-d`) at the prompt.

## Notes

This program will attempt to ignore invalid input and will only exit on an un-recoverable error (or when an `EOF` or `q` is input).

* If invalid input is entered, the application will ignore it and continue. (ex. `> %` is not a valid operand or operator)
  ```
  > %
  Invalid input. Ignoring.
  >
  ```

* If an undefined operation is attempted, the application will ignore it and continue (ex. division by zero). The stack will be unaffected.
  ```
  > 1
  1
  > 0
  0
  > /
  Error performing operation '/': Cannot divide by zero. Ignoring.
  >
  ```

* If there are not enough operands on the stack to support the requested operation, the application will ignore the operator and continue.
  ```
  > 1
  1
  > +
  Not enough operands. Ignoring operator.
  >
  ```

* If the length of an input line, including end of line character(s), is longer than 65,536 (the MaxScanTokenSize in Go is defined as 64*1024), this will cause the input to be read in multiple 'chunks'. This application is only reading the first chunk. This will cause unintended consequences, but since all legal `float64` values can be represented in fewer characters, this application does not attempt to merge lines longer than 65,536.
