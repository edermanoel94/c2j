# C2J

A simple command line for convert csv into an array of json

## Installation (WIP)

```
go install
```

## Usage

```
c2j [option]

Flags:
-d, --delimiter                     Choose a delimiter for csv, example: "," or ";"
```

## Examples


```
cat CSV_FILE.csv | c2j | jq        
sed 's/\,/\;/g' CSV_FILE.csv | c2j --delimiter ";" | jq
```

## Running tests

```
go test
```

## Benchmarks (WIP)

```
go test -bench=.
```

## References:

- [JQ](https://stedolan.github.io/jq/)
