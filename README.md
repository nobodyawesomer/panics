# results
A Golang error handling library, inspired by Rust.

Useful for converting those pesky `(value, error)` returns into something useful. No more of this:
```go
value, err := funkyfunction()
if err != nil {
	panic(err)
}
```
And instead, this:
```go
value := Unwrap(funkyfunction())
```

You can also catch specific errors and Unwrap to default values. :)