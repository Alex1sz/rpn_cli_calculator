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
rpn_cli_calculator is a RPN calculator CLI in Golang. Although I felt as a whole developing a repl CLI with Ruby would be easier mostly due to my own familiarity, ease of object oriented programming abstractions, and the ability to plugin gems, I went with Golang instead. In terms of data structure, it was apparent that a stack backed by a linked-list made sense to use. I felt that a Golang implementation would be more concise, efficient, and easier to modify overall. In particular the idea of using a Ruby array as a stack felt dirty, and I didn't like the idea of implementing a linked-list using classes as linked-list behavior is already possible via the standard Ruby array methods. In particular I felt like a Golang LIFO stack backed by a linked-list would be more efficient, cleaner, and subsequently make it easier to reason about.

### Calculator Description
The RPN calculator's evaluation logic is within the 'calculator' subpackage. As is 'calculator/stack.go' a generic LIFO stack implementation. The utility of moving evaluation and the dispatch table into its own 'calculator' subpackage is mainly ease of future changes such as swapping out stack implementation and/or modifications to the evaluator logic 'calculator/evaluator.go'.

### Dispatch Table explanation
The evaluator utilizes a dispatch table. Advantages of using a dispatch table here are fewer conditional statements, and more compact code code structure. Another project specific advantage of utilizing a dispatch table was that it makes it easier to make changes to the user input handling code. Modifying the entry point and input handling code requires no modification or minimal changes to the dispatch table. For a developer relatively unfamiliar with programming repl CLI's, I felt that ease of modifications to the input handling code would be advantageous. Another advantage of a dispatch table in this case is aside from initial dispatchTable setup; using a dispatch table makes it significantly easier to test. Adding additional input/expression test cases is as simple as adding another input str mapped to a set of operators, and providing the expected calculation result for each. A disadvantage of using a dispatch table is that there is an extra level of indirection. Cognitively the way a reverse polish notation operates and evaluates inputs is counter intuitive to the way I calculate and handle arithmetic mentally. As a result, I felt that reducing conditional statements via a dispatch table would reduce the cognitive load surrounding calculations/operations etc as oppose to packing a bunch of conditional logic into a switch statement or having redundant conditional statements littered throughout the code.

### Todos// Things Id Spend Additional Time on
Ideally, the package would have 100% code coverage as oppose to: 85.2%.

1. To achieve better code coverage, I would add tests for main.go calcToShell(), and add tests for remaining gaps.
1. Add output/input color highlighting time permitted.
1. Expand the 'help' info command to be more useful.
1. Add a version command.
1. Potentially move stack to its own subpackage ??
1. Move the general typecasting functions out of calculator.
