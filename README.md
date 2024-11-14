# GoBail

Exit immediately your app when you see an error.

[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/dnnrly/gobail)](https://github.com/dnnrly/gobail/releases/latest)
[![report card](https://goreportcard.com/badge/github.com/dnnrly/gobail)](https://goreportcard.com/report/github.com/dnnrly/gobail)
[![godoc](https://godoc.org/github.com/dnnrly/gobail?status.svg)](http://godoc.org/github.com/dnnrly/gobail)

![GitHub watchers](https://img.shields.io/github/watchers/dnnrly/gobail?style=social)
![GitHub stars](https://img.shields.io/github/stars/dnnrly/gobail?style=social)
[![Twitter URL](https://img.shields.io/twitter/url?style=social&url=https%3A%2F%2Fgithub.com%2Fdnnrly%2Fgobail)](https://twitter.com/intent/tweet?url=https://github.com/dnnrly/gobail)

Sometimes when your app encounters an error, it can't be recovered from and the only option is to exit immediately. It doesn't always make sense to pass the error all the way back to the your error handling, just so you can print a string to the console - you just want to exit now. This library makes it easy to do exactly that using a fluent API. I've also made sure that the library is fully tested and validated so that you don't have to. It may sound like a lot of effort but it's important to make sure that there are no nasty suprises - and have evidence that you made sure.

Using this library allows you to safely exit with a non-zero code or just panic when you need to, without without having to worry about complex test code to raise your code coverage metrics.

### Using `gobail`

```go
package main

import "github.com/dnnrly/gobail"

var err = errors.New("an error")

func returnError() error                        { return err }
func returnValAndError() (int, error)           { return 0, err }
func return2ValsAndError() (int, string, error) { return 0, "str", err }

func main() {
    // Take a look at the package documentation for more details
    gobail.Run(returnError()).OrExit()
    result := gobail.Return(returnValAndError()).OrPanic()
    result1, result2 := gobail.Return(return2ValsAndError()).OrExitMsg("something went wrong: %v")
}
```

### Running Unit Tests

```bash
$ make test
```

### Running Acceptance tests

```bash
$ make acceptance-test
```

## Important `make` targets

* `test` - runs unit tests
* `ci-test` - run tests for CI validation
* `acceptance-test` - run the acceptance tests
* `lint` -  run linting
* `clean` - clean project dependencies


## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/dnnrly/gobail/tags). 

## Authors

* **Pascal Dennerly** - *Original library author* - [dnnrly](https://github.com/dnnrly)

See also the list of [contributors](https://github.com/dnnrly/gobail/contributors) who participated in this project.

## License

This project is licensed under the Apache 2.0 License - see the [LICENSE](LICENSE) file for details
