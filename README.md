# results
A Golang error handling library, inspired by Rust (and somewhat Java). No non-std imports, and minimal exports so you can dot import it without a bloated namespace.

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

## Usage

Go get it:
`go get github.com/nobodyawesomer/results`

And then (suggested) dot import it:
```go
import (
	. "github.com/nobodyawesomer/results"
)
```

Then, wrap your function calls, and unwrap the errors:
```go
value := Try(strconv.atoi("bloop"))
	.UnwrapOr(10) // defaults to 10 on any error :)
```

See documentation for more options.

## Roadmap
Basically whatever I find useful. I'll probably look into an ergonomic way to incorporate logging, because I often `if err != nil {log.Fatal(err)}`.

## License
MIT