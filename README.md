# rpn_cli_calculator

## Build Status
[![Build Status](https://travis-ci.com/Alex1sz/rpn_cli_calculator.svg?token=tCgPM15Yp3EA7UzetByu&branch=master)](https://travis-ci.com/Alex1sz/rpn_cli_calculator)

## How to Run
1. Install Golang if you have not already, for instructions see: <https://golang.org/doc/install>
1. Install: `go get github.com/alex1sz/rpn_cli_calculator`
1. After install change directory to the package, build and install: `cd $GOPATH/src/github.com/alex1sz/rpn_cli_calculator && go build && ./rpn_cli_calculator`
1. To Build & Run: `go build && ./rpn_cli_calculator`
1. After building you can also run via: `go run main.go` or `./rpn_cli_calculator`
1. Provide a set of arguments via command line `./rpn_cli_calculator 2 3 +`
1. Or wait:
```
Welcome to: ./rpn_cli_calculator
-> 4 4 +
8.000000
-> 1
1.000000
-> +
9.000000
->
```


## Description
rpn_cli_calculator is a RPN calculator CLI in Golang.
