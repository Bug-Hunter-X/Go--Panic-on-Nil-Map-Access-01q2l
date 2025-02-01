# Go: Panic on Nil Map Access

This repository demonstrates a common error in Go: panicking due to accessing a nil map.  The `bug.go` file shows the problematic code, and `bugSolution.go` offers a solution.

**Problem:**

Attempting to access a map before it's initialized (or if it's `nil`) leads to a runtime panic.  This is unexpected behavior if proper error handling isn't implemented. 

**Solution:**

The solution involves checking for `nil` before accessing the map's elements. This can be achieved using an `if` statement or the built-in `ok` idiom.