# Monkey Interpreter

Let's write an interpreter in Go, along with the fantastic book
[Writing An Interpreter In Go](https://interpreterbook.com/) by Thorsten Ball.

The interpreter is written in Go, but runs code in Monkey, which looks like this:
```
let fibonacci = fn(x) {
  if (x == 0) {
    0                // Monkey supports implicit returning of values
  } else {
    if (x == 1) {
      return 1;      // and explicit return statements
    } else {
      fibonacci(x - 1) + fibonacci(x - 2); // and Recursion!
    }
  }
};

let result = fibonacci(100);
```

### Run the interpreter

If you have Go installed, you can `git pull` or `go get` this repo, and then:
```
go run cmd/monkey/main.go
```
