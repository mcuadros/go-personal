go-personal [![Build Status](https://travis-ci.org/mcuadros/go-personal.png?branch=master)](https://travis-ci.org/mcuadros/go-personal) [![GoDoc](http://godoc.org/github.com/mcuadros/go-personal?status.png)](http://godoc.org/github.com/mcuadros/go-personal)
==============================

A tiny library to score strings as fullname or emails. Given a list of strings
the library returns which is the most accurate Fullname or Email.


Installation
------------

The recommended way to install go-personal

```
go get github.com/mcuadros/go-personal
```

Examples
--------

A basic example:

```go
names := []string{"mcuadros", "Máximo Cuadros", "Máximo Cuadros Ortiz"}
best := personal.GetBestFullName(names)
fmt.Printlm(best)
//Returns "Máximo Cuadros Ortiz"

names := []string{"bar@foo.com", "bar@gmail.com", "bar@localdomain.local"}
best := personal.GetBestEmail(names)
fmt.Printlm(best)
//Returns "bar@gmail.com"
```

License
-------

MIT, see [LICENSE](LICENSE)
