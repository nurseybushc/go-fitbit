# Fitbit API Client in Golang

go-fitbit is a client library for accessing the [Fitbit Web API](https://dev.fitbit.com/build/reference/web-api/)

The client is designed to access the API with a personal token.

The library design is inspired by the [go-github](https://github.com/google/go-github) library.

# Status

This is very much a work in progress, for now only entity models are present.

The next point to tackle are:

* finishing the infrastructure in fitbit.go (see (here)[https://github.com/google/go-github/blob/master/github/github.go])
* figure out the authentication
* implement the various calls (possibly starting from heart_rate.go)
* implement tests
* setup CI

# Contributing

Pull requests are more than welcome.

# License

Copyright 2018 Simon Accascina<simon@accascina.me>

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.