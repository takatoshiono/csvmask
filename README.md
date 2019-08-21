# csvmask
csvmask is a tool of masking CSV encoded data.

## Usage

1. Prepare CSV encoded data

```
$ cat testdata/test.csv
"ID","Name","Address"
4085ff59-39bd-4cc3-8a55-c5b1c6785922,Adam Smith,Kirkcaldy United Kingdom
```

2. Execute csvmak with masking rules
The rule is formatted as CSV and have same number of fields as the prepared CSV data.

```
$ cat testdata/test.csv | csvmask -rule "Raw,Hash,Raw" -skipheader
ID,Name,Address
4085ff59-39bd-4cc3-8a55-c5b1c6785922,PbZ8hc4alo56RYc9/m+vECyVdjHqZRGMlxUGigh3/uE,Kirkcaldy United Kingdom
```

## Masking Rules

- Raw
  - Output as is.
- Hash
  - Output hashed data.
- Checksum
  - Output checksum of data.

## Install

```
$ go get -u github.com/takatoshiono/csvmask
```

## Known issues

Output CSV format may be different from original due to the csv package of Go.

- A quoted-field will be removed the quote characters.
- A field contains white-space without quote will be quoted.
