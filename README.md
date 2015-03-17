go-fullname [![Build Status](https://travis-ci.org/mcuadros/go-fullname.png?branch=master)](https://travis-ci.org/mcuadros/go-fullname) [![GoDoc](http://godoc.org/github.com/mcuadros/go-fullname?status.png)](http://godoc.org/github.com/mcuadros/go-fullname)
==============================

A very small library to score strings as fullname. Given a list of strings the library  returns which is the most accurate Fullname.


Installation
------------

The recommended way to install go-fullname

```
go get github.com/mcuadros/go-fullname
```

Examples
--------

A basic example:

```go
names := []string{"mcuadros", "Máximo Cuadros", "Máximo Cuadros Ortiz"}
best := fullname.GetBestFullName(names)
fmt.Printlm(best)
//Returns "Máximo Cuadros Ortiz"
```

License
-------

MIT, see [LICENSE](LICENSE)
