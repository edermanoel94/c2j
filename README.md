# C2J

[![Go](https://github.com/edermanoel94/c2j/actions/workflows/go.yml/badge.svg)](https://github.com/edermanoel94/c2j/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/edermanoel94/c2j)](https://goreportcard.com/report/github.com/edermanoel94/c2j)
[![codecov](https://codecov.io/gh/edermanoel94/c2j/branch/master/graph/badge.svg)](https://codecov.io/gh/edermanoel94/c2j)

A simple command line for convert **CSV** in **JSON**

## Install

### Building from Source

With Go 1.17 or higher:

```
go install github.com/edermanoel94/c2j
```

### Usage

⚠️ CSV need to have header.

#### Convert

To convert, run the `c2j` command to read from **stdin**, using standard delimiter, which is *comma*.

```
$ cat example_comma.csv | c2j
```

#### Convert with custom delimiter

Use the `--delimiter` or `-d` flag to specify a delimiter.

```
$ cat example_semicolon.csv | c2j --delimiter ";"
```

#### Demo

![Demonstration](demo.gif)

## Todo

- [ ] Convert without header
- [ ] Use a path argument to convert
- [ ] Save output to a file
- [ ] Benchmark

## Contributing

Pull requests for new features, bug fixes, and suggestions are welcome!

## License

[MIT](https://github.com/edermanoel94/c2j/blob/master/LICENSE)

## Live

I use [Twitch](https://twitch.tv/thegravidade) to stream, follow me, to see new features on this project.
