# randgen

A Golang package to generate random strings and numbers.


# Installation

---

```
go mod install github.com/wisdommatt/randgen
```

# Features

---

* Generate random strings.
* Generate random numbers.


# Requirements

---

go 1.5


# Documentation

---

https://pkg.go.dev/github.com/wisdommatt/randgen


# How to use (examples)

---

* #### Generating 10 random strings.

```go
stringGenerator := randgen.NewStringGenerator()
randomString := stringGenerator.GenerateFromSource(randgen.StringAlphaNumericSource, 10) // source, length
```


* #### Generating 8 random numbers.

```go
numberGenerator := randgen.NewNumberGenerator()
randomNumbers := numberGenerator.GenerateFromSource(randgen.NumberSource, 8) // source, 
```
