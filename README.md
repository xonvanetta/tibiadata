# TibiaData API.

A go implementation for an API. Missing a lot of custom struct that i have'nt seen yet nor tested for.  

## Installing

`go get github.com/xonvanetta/tibiadata`

## Running the tests

Run this:  `make test` or  `go test ./... -v -short`.

Reason for the short flag is that the integration test takes a really long time and it test live data. Which means a lot of unmarshal errors.

## Adding tests

Tests should contain a copy of the data that you used to write the added feature/bug. Currently most of them will contain weird json that go has a hard time unmarshal.

## Versioning

We use [SemVer](http://semver.org/) for versioning. Currently this package will be below v1 which will mean breaking changes between versions.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details