# C2J

[![Go](https://github.com/edermanoel94/c2j/actions/workflows/go.yml/badge.svg)](https://github.com/edermanoel94/c2j/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/edermanoel94/c2j)](https://goreportcard.com/report/github.com/edermanoel94/c2j)
[![codecov](https://codecov.io/gh/edermanoel94/c2j/branch/master/graph/badge.svg)](https://codecov.io/gh/edermanoel94/c2j)

A simple command line for convert **CSV** in **JSON** list of objects based on header.

## Install

With Go 1.17 or higher:

```
go install github.com/edermanoel94/c2j
```

### Usage

#### Convert

To convert, run the `c2j` command to read from **STDIN**, using standard delimiter, which is *comma*.

```
$ cat example_comma.csv | c2j
```

#### Convert with custom delimiter

Use the `--delimiter` or short version`-d` flag to specify a delimiter.

```
$ cat example_semicolon.csv | c2j --delimiter ";"
```

#### Demo

![Demonstration](demo.gif)

## Todo

- [x] Convert without header and generate keys based on index
- [ ] Use a path to file argument for convert
- [ ] Save output to a file
- [ ] Benchmark

## Contributing

Pull requests for new features, bug fixes, and suggestions are welcome!

## License

[MIT](https://github.com/edermanoel94/c2j/blob/master/LICENSE)

## Live

I use [Twitch](https://twitch.tv/thegravidade) to stream, follow me, to see new features on this project.
