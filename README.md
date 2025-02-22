# Unexpected Type Assertion Behavior in Go's Type Switch

This repository demonstrates an uncommon bug related to type assertions within Go's type switch statement when dealing with numeric types like `int` and `float64`. The issue highlights a subtlety in how Go's type system handles these types under certain conditions.

## Description

The `bug.go` file contains a Go program that showcases the unexpected behavior.  It uses a type switch to determine the type of an interface{} variable.  The results are not what one might intuitively expect, especially when comparing `int` and `float64` types.

## How to Reproduce

1. Clone the repository.
2. Navigate to the root directory of the repository.
3. Run the Go program using `go run bug.go`.

## Solution

The `bugSolution.go` provides a more robust way to handle such cases, demonstrating how to avoid the unexpected behavior and ensure accurate type checking.