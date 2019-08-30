# csvmask [![CircleCI](https://circleci.com/gh/takatoshiono/csvmask.svg?style=svg)](https://circleci.com/gh/takatoshiono/csvmask) [![codecov](https://codecov.io/gh/takatoshiono/csvmask/branch/master/graph/badge.svg)](https://codecov.io/gh/takatoshiono/csvmask)
csvmask is a CSV masking tool.

## Usage

1. Prepare CSV encoded data

```
$ cat testdata/test.csv
"ID","Name","Address"
4085ff59-39bd-4cc3-8a55-c5b1c6785922,Adam Smith,Kirkcaldy United Kingdom
```

2. Create template

```
{{.Field1}},{{hash .Field2}},{{.Field3}}
```

3. Execute csvmask with template

```
$ cat testdata/test.csv | csvmask -template "{{.Field1}},{{hash .Field2}},{{.Field3}}" -skipheader
ID,Name,Address
4085ff59-39bd-4cc3-8a55-c5b1c6785922,PbZ8hc4alo56RYc9/m+vECyVdjHqZRGMlxUGigh3/uE,Kirkcaldy United Kingdom
```

## Template

The template is a text of text/template package of Go.

### Arguments

The fields of CSV record can be referred as arguments. The name is preceeded by a period such as

- .Field1
- .Field2
- .Field3
- ...

### Functions

The following functions are defined.

- hash
- checksum
- right
- left

## Installation

### Just get binaries

Go to the [releases page](https://github.com/takatoshiono/csvmask/releases) and download zip file.

### go get

```
$ go get -u github.com/takatoshiono/csvmask
```

### Build yourself

First clone the repository and

```
$ make install
```
